package auth

import (
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
    username string
    expiry time.Time
}

func (s Session) isExpired() bool {
    return s.expiry.Before(time.Now())
}

func ClearCookie(w http.ResponseWriter) {
    cookie := &http.Cookie{
        Name: "session_token",
        Value: "",
        Expires: time.Now().Add(-1 * time.Minute),
        Path: "/",
    }

    http.SetCookie(w, cookie)
}

func Login(w http.ResponseWriter, r *http.Request, conn *pgx.Conn, sessions map[string]Session) {

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
        username: credentials.Email,
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

func Auth(w http.ResponseWriter, r *http.Request, sessions map[string]Session) {
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
    http.Redirect(w, r, "/static/web.html", http.StatusSeeOther)
}
