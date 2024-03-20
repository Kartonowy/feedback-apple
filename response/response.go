package response

import (
	"app/auth"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"path"

	"github.com/jackc/pgx/v5"
)

func GetOpinions(w http.ResponseWriter, r *http.Request, conn *pgx.Conn) string {
    uuid, err := r.Cookie("session_token")

    if err != nil {
        panic(err)
    }

    email, err := auth.GetSession(uuid.Value)

    if err != nil {
        panic(err)
    }

    var html_response string

    res, err := conn.Query(context.Background(),
    fmt.Sprintf("SELECT id, responses FROM classes WHERE id = ANY ((SELECT class_codes from teachers WHERE email = '%s')::text[])",
        email))

    if err != nil {
        panic(err)
    }

    for res.Next() {
        var id string
        var responses []string
        err = res.Scan(&id, &responses)
        if err != nil {
            panic(err)
        }
        html_response += fmt.Sprintf("<div hx-get='/redirect/%s' hx-trigger='click'><h1>Lobby: %s</h1></div>",id, id)
    }
    return html_response
}

func GetDetails(w http.ResponseWriter, r *http.Request, conn *pgx.Conn) string {
    var html_response string

    code := path.Base(r.URL.Path)
    fmt.Printf("%s\n", code)
    res, err := conn.Query(context.Background(),
    fmt.Sprintf("SELECT responses FROM classes WHERE id = '%s'", code))

    if err != nil {
        panic(err)
    }

    for res.Next() {
        var responses []string
        err = res.Scan(&responses)
        if err != nil {
            panic(err)
        }
        var internal_div string

        for _, v := range responses {
            var component string
            fmt.Printf("%v\n", v)
            var response map[string]string
            json.Unmarshal([]byte(v), &response)
            first_choice := response["first_choice"]
            second_choice := response["second_choice"]
            third_choice := response["third_choice"]
            rating := response["rating"]

            if first_choice != "" {
                component += fmt.Sprintf("<h1>first choice: %v</h1>", first_choice)
                if second_choice != "" {
                    component += fmt.Sprintf("<h2>second choice: %v</h2>", second_choice)
                    if third_choice != "" {
                        component += fmt.Sprintf("<h3>third choice: %v</h3>", third_choice)
                    }
                }
            }
            if rating != "" {
                component += fmt.Sprintf("<div>rating: %v</div>", rating)
            }
            internal_div += fmt.Sprintf("<div>%v</div>", component)
        }
        html_response += fmt.Sprintf("<div>%v</div>", internal_div)
    }

    fmt.Printf("%s\n", html_response) 
    return html_response
}
