package auth

import (
	"app/codegen"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)


type Credentials struct {
    Email string `json:"email"`
    Password string `json:"password"`
}

type Session struct {
    Username string
    expiry time.Time
}


func (s Session) isExpired() bool {
    return s.expiry.Before(time.Now())
}

var sessions = map[string]Session{}

func ClearCookie(w http.ResponseWriter) {
    cookie := &http.Cookie{
        Name: "session_token",
        Value: "",
        Expires: time.Now().Add(-1 * time.Minute),
        Path: "/",
    }

    http.SetCookie(w, cookie)
}

func CodeGen(w http.ResponseWriter, r *http.Request, conn *pgx.Conn) string {
        fmt.Println("Generating code")
        cookie, err := r.Cookie("session_token")

        fmt.Println(cookie.Value)

        if err != nil {
            panic(err)
        }
        
        fmt.Printf("%s\n", sessions[cookie.Value].Username)

        // Check if code already exists
        res, err := conn.Query(context.Background(), fmt.Sprintf("SELECT class_codes FROM teachers"))
        if err != nil {
            panic(err)
        }

        defer res.Close()
        code := codegen.RandSeq(6)
        isUnique := false
        for !isUnique {

            code = codegen.RandSeq(6)

            for res.Next() {
                var codes []string

                err := res.Scan(&codes)

                if err != nil {
                    panic(err)
                }
                for _, v := range codes {
                    if v == code {
                        isUnique = false
                        break
                    } else {
                        isUnique = true
                    }
                }
            }
        }

        _, err = conn.Exec(context.Background(),
        fmt.Sprintf("UPDATE teachers SET class_codes = array_append(class_codes, '%s') WHERE email = '%s';", 
        code, sessions[cookie.Value].Username))
        fmt.Println("Code added to teacher " + sessions[cookie.Value].Username)
        if err != nil {
            panic(err)
        }

        _, err = conn.Exec(context.Background(),
        fmt.Sprintf("INSERT INTO classes (id, responses) VALUES ('%s', '{}');", code))
        fmt.Println("Code added to classes")
        if err != nil {
            panic(err)
        }

        fmt.Println("Code generated")

        return code
}

func Login(w http.ResponseWriter, r *http.Request, conn *pgx.Conn) {

    err := r.ParseForm()

    if err != nil {
        fmt.Printf("%v", err)
        w.WriteHeader(http.StatusBadRequest)
        return
    }

    credentials := Credentials{
        Email: r.PostForm.Get("email"),
        Password: r.PostForm.Get("password"),
    }
    var password string
    fmt.Println("Logging in!")

    res := conn.QueryRow(context.Background(), fmt.Sprintf("SELECT password FROM teachers WHERE email = '%s';", r.PostForm.Get("email")));
    res.Scan(&password)


    if err := bcrypt.CompareHashAndPassword([]byte(password), []byte(r.PostForm.Get("password"))); err != nil {
        w.WriteHeader(http.StatusUnauthorized)
        return
    }
    
    sessionToken := uuid.NewString()
    expiresAt := time.Now().Add(120 * time.Minute)

    sessions[sessionToken] = Session {
        Username: credentials.Email,
        expiry: expiresAt,
    }

    http.SetCookie(w, &http.Cookie{
        Name: "session_token",
        Value: sessionToken,
        Expires: expiresAt,
        Path: "/",
    })

    fmt.Println("Passwords match!")
    http.Redirect(w, r, "/welcome/", http.StatusSeeOther)
}

func Auth(w http.ResponseWriter, r *http.Request) {
    c, err := r.Cookie("session_token")

    if err != nil {
        if err == http.ErrNoCookie {
            w.WriteHeader(http.StatusUnauthorized)
            return
        }
        w.WriteHeader(http.StatusBadRequest)
        return
    }

    sessionToken := c.Value

    userSession, exists := sessions[sessionToken]

    if !exists {
        w.WriteHeader(http.StatusUnauthorized)
        return
    }

    if userSession.isExpired() {
        delete(sessions, sessionToken)
        w.WriteHeader(http.StatusUnauthorized)
        return
    }
    fmt.Println("Redirecting to web.html")
    http.ServeFile(w, r, "static/web.html")
}

func GetSession(uuid string) (string, error) {
    if _, exists := sessions[uuid]; !exists {
        return "", fmt.Errorf("No session found")
    }
    return sessions[uuid].Username, nil
}
