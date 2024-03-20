package join

import (
	"fmt"
	"net/http"
    "encoding/json"
	"github.com/jackc/pgx/v5"
    "context"
)

type Response struct {
    FirstChoice string `json:"first_choice"`
    SecondChoice string `json:"second_choice"`
    ThirdChoice string `json:"third_choice"`
    Rating string `json:"rating"`
}

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
    http.ServeFile(w, r, "static/form.html")
}

func Rate(w http.ResponseWriter, r *http.Request, conn *pgx.Conn) {
    r.ParseForm()
    cookie, err := r.Cookie("lobby_code")

    if err != nil {
        fmt.Println("No cookie found")
        w.WriteHeader(http.StatusUnauthorized)
        return
    }

    var n Response

    decoder := json.NewDecoder(r.Body)

    err = decoder.Decode(&n)

    if err != nil {
        panic(err)
    }

    fmt.Printf("r %v\n", n)


    rat := Response{ 
        FirstChoice: n.FirstChoice,
        SecondChoice: n.SecondChoice,
        ThirdChoice: n.ThirdChoice,
        Rating: n.Rating,
    }

    jsoned, _ := json.Marshal(rat)

    fmt.Println("Rating: ", string(jsoned))


    _, err = conn.Exec(context.Background(),
    "UPDATE classes SET responses = array_append(responses, $1) WHERE id = $2",
    []byte(jsoned), cookie.Value)

    if err != nil {
        panic(err)
    }

	w.Header().Add("HX-Redirect", "/thankyou/")
}
