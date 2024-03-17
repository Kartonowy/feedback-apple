package join

import (
	"fmt"
	"net/http"
	"github.com/jackc/pgx/v5"
    "context"
)

func JoinClass(w http.ResponseWriter, r *http.Request, code string, conn *pgx.Conn) {

    res, err := conn.Query(context.Background(), fmt.Sprintf("SELECT class_codes FROM teachers"))
    if err != nil {
        panic(err)
    }


    defer res.Close()

    isCorrect := false

    for res.Next() {
        var codes []string


        err = res.Scan(&codes)

        if err != nil {
            panic(err)
        }

        for _, c := range codes {
            fmt.Printf("code: %s ?=", code)
            fmt.Printf("c: %s\n", c)
            fmt.Printf("isCorrect: %t\n", c == code)
            if code == c {
                isCorrect = true
                break
            } else {
                isCorrect = false
            }
        }
        if isCorrect {
            break
        }
    }

    if !isCorrect {
        http.Error(w, "Invalid class code", http.StatusUnauthorized)
        return
    }
    
    fmt.Printf("Joining class with code: %s\n", code)    
    cookie := &http.Cookie{
        Name: "lobby_code",
        Value: code,
        Path: "/rate/",
    }

    http.SetCookie(w, cookie)
    http.Redirect(w, r, "/static/form.html", http.StatusSeeOther)
}
