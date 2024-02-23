package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

type OpinionForm struct {
	teacherselect int32
	review        string
}

type Teacher struct {
	id         int32               `db:"id"`
	email      string              `db:"email"`
	first_name string              `db:"first_name"`
	last_name  string              `db:"last_name"`
	school_id  int32               `db:"schoold_id"`
	subject    []string            `db:"subject"`
	messages   []map[string]string `db:"messages"`
}

type Student struct {
	id         int32  `db:"id"`
	email      string `db:"email"`
	first_name string `db:"first_name"`
	last_name  string `db:"last_name"`
	school_id  int    `db:"school_id"`
}
type School struct {
	id          int32  `db:"id"`
	school_name string `db:"school_name"`
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

	var teacher Teacher
	row := conn.QueryRow(context.Background(), "SELECT * FROM teacher LIMIT 1")
	err = row.Scan(&teacher.id, &teacher.email, &teacher.first_name, &teacher.last_name, &teacher.school_id, &teacher.subject, &teacher.messages)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Here's the output: %v\n", teacher)

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	router.HandleFunc("/schools/", func(w http.ResponseWriter, r *http.Request) {
		w.Write(GetSchools(conn))
	})

    router.HandleFunc("/addstudent/", func(w http.ResponseWriter, r *http.Request) {
        r.ParseForm()
        query := r.PostForm.Get("name")
        defer r.Body.Close()
        fmt.Printf("%v\n", query)
    })
    // TODO: ADD ADDING TO THe DB ASAP


	fs := http.FileServer(http.Dir("static"))
	router.Handle("/static/*", http.StripPrefix("/static/", fs))


	err = http.ListenAndServe(":3333", router)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Listening on port 3333")
}

func AddMessage(conn *pgx.Conn, id int32) {
	message := map[string]string{"behaviour": "good"}
	messages_json, _ := json.Marshal(message)

	_, err := conn.Exec(context.Background(), `UPDATE teacher SET messages = messages || $2 WHERE teacher.id = $1`, id, [][]byte{messages_json})

	if err != nil {
		panic(err)
	}
}

func GetSchools(conn *pgx.Conn) []byte {
	rows, err := conn.Query(context.Background(), "SELECT id, school_name FROM school")
	output := make(map[string]string)

	defer rows.Close()
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var school School
		err = rows.Scan(&school.id, &school.school_name)
		if err != nil {
			panic(err)
		}
		output[fmt.Sprintf("%v", school.id)] = school.school_name
	}
	realoutput, _ := json.Marshal(output)
	fmt.Printf("%v\n", realoutput)
	return realoutput
}
