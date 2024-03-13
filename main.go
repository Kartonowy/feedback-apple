package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
    "app/auth"
)

type Teacher struct {
	id          int32    `db:"id"`
	email       string   `db:"email"`
	first_name  string   `db:"first_name"`
	last_name   string   `db:"last_name"`
	class_codes []string `db:"class_codes"`
}



var sessions = map[string]auth.Session{}

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
		w.Write([]byte("welcome"))
	})

    router.Post("/register/", func(w http.ResponseWriter, r *http.Request) {
        addAccount(w, r, conn) 
    })
    
    router.Post("/login/", func(w http.ResponseWriter, r *http.Request) {
        auth.Login(w, r, conn, sessions)
    })

    router.Get("/welcome/", func(w http.ResponseWriter, r *http.Request) {
        auth.Auth(w, r, sessions)
    })

    router.Post("/logout/", func(w http.ResponseWriter, r *http.Request) {
       auth.ClearCookie(w) 
       w.Header().Add("HX-Redirect", "/static/login.html")
    })

	fs := http.FileServer(http.Dir("static"))
	router.Handle("/static/*", http.StripPrefix("/static/", fs))

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
        r.PostForm.Get("email"), r.PostForm.Get("first_name"), r.PostForm.Get("last_name"), hash  ,[]string{})

        if err != nil {
            panic(err)
        }
}
