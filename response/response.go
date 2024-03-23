package response

import (
	"app/auth"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

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
        html_response += fmt.Sprintf("<li hx-post='/get_details/%s' hx-trigger='click' hx-target='#responsebox' hx-swap='outerHTML'><h1>%s</h1></li>",id, id)
    }
    return html_response
}

func GetDetails(w http.ResponseWriter, r *http.Request, conn *pgx.Conn, code string) string {
    var html_response string

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
            second_choice := response["second_choice"]
            third_choice := response["third_choice"]
            rating := response["rating"]

                if second_choice != "" {
                    component += fmt.Sprintf("<h2>%v</h2>", second_choice)
                    if third_choice != "" {
                        component += fmt.Sprintf("<h3>%v</h3>", third_choice)
                    }
                }
            if rating != "" {
                if _, err := strconv.Atoi(rating); err == nil {
                    component += fmt.Sprintf("<div>rating: %v/10</div>", rating)
                } else {
                    component += fmt.Sprintf("<div>rating: %v</div>", rating)
                }
            }
            internal_div += fmt.Sprintf("<li>%v</li>", component)
        }
        html_response += fmt.Sprintf("<ul id='responsebox'>%v</ul>", internal_div)
    }

    fmt.Printf("%s\n", html_response) 
    return html_response
}
