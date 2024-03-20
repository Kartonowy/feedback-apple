package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
    "strings"

    "app/response"
	"app/auth"
    "app/join"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

type Teacher struct {
	id          int32    `db:"id"`
	email       string   `db:"email"`
	first_name  string   `db:"first_name"`
	last_name   string   `db:"last_name"`
	class_codes []string `db:"class_codes"`
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("Error loading env file: %s", err)
	}

	conn, err := pgx.Connect(context.Background(), "postgres://octavia:twojastara@localhost:5432/postgres")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	defer conn.Close(context.Background())

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "static/join.html")
	})
    router.Get("/account/", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "static/login.html")
    })

	router.Post("/register/", func(w http.ResponseWriter, r *http.Request) {
		addAccount(w, r, conn)
	})

	router.Post("/login/", func(w http.ResponseWriter, r *http.Request) {
		auth.Login(w, r, conn)
	})

	router.Get("/welcome/", func(w http.ResponseWriter, r *http.Request) {
		auth.Auth(w, r)
	})

	router.Post("/logout/", func(w http.ResponseWriter, r *http.Request) {
		auth.ClearCookie(w)
		w.Header().Add("HX-Redirect", "/account/")
	})

	router.Post("/gencode/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(auth.CodeGen(w, r, conn)))

	})

    router.Post("/redirect_lobby/", func(w http.ResponseWriter, r *http.Request) {
        r.ParseForm()
        http.Redirect(w, r, fmt.Sprintf("/join/%s", strings.Trim(r.PostForm.Get("code"), " ")), http.StatusSeeOther)
    })


    router.Get("/join/{code}", func(w http.ResponseWriter, r *http.Request) {
        code := chi.URLParam(r, "code")
        join.JoinClass(w, r, code, conn)
    })

    router.Post("/rate/", func(w http.ResponseWriter, r *http.Request) {
        join.Rate(w, r, conn)
    })

    router.Post("/get_opinions/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte(response.GetOpinions(w, r, conn)))
    })

    router.Post("/get_details/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte(response.GetDetails(w, r, conn)))
    })

    router.Get("/redirect/{code}", func(w http.ResponseWriter, r *http.Request) {
        code := chi.URLParam(r, "code")
        w.Header().Add("HX-Redirect", fmt.Sprintf("/details/%s", code))
    })

    router.Get("/details/{code}", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "static/details.html")
    })

    router.Get("/styles/{file}", func(w http.ResponseWriter, r *http.Request) {
        file := chi.URLParam(r, "file")
        http.ServeFile(w, r, fmt.Sprintf("static/styles/%s", file))
    })

    router.Get("/scripts/{file}", func(w http.ResponseWriter, r *http.Request) {
        file := chi.URLParam(r, "file")
        http.ServeFile(w, r, fmt.Sprintf("static/scripts/%s", file))
    })

    router.Get("/thankyou/", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "static/thankyou.html")
    })

	err = http.ListenAndServe(":3333", router)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Listening on port 3333")
}


func AddMessage(conn *pgx.Conn) {
	message := map[string]string{"behaviour": "good"}
	messages_json, _ := json.Marshal(message)

	_, err := conn.Exec(context.Background(), `UPDATE teacher SET messages = messages || $2 WHERE teacher.id = $1`, [][]byte{messages_json})

	if err != nil {
		panic(err)
	}
}

func addAccount(w http.ResponseWriter, r *http.Request, conn *pgx.Conn) {
	r.ParseForm()

	defer r.Body.Close()

	hash, err := bcrypt.GenerateFromPassword([]byte(r.PostForm.Get("password")), 14)

	if err != nil {
		panic(err)
	}

	_, err = conn.Exec(context.Background(),
		`INSERT INTO teachers (email, first_name, last_name, password ,class_codes) 
        VALUES ($1, $2, $3, $4, $5)`,
		r.PostForm.Get("email"), r.PostForm.Get("first_name"), r.PostForm.Get("last_name"), hash, []string{})

	if err != nil {
		panic(err)
	}
}
