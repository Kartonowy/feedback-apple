package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
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
	school_id   int32  `db:"id"`
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

	fmt.Printf("Here's the output: %v", teacher)

    AddMessage(conn, 1)

	row = conn.QueryRow(context.Background(), "SELECT * FROM teacher LIMIT 1")
	err = row.Scan(&teacher.id, &teacher.email, &teacher.first_name, &teacher.last_name, &teacher.school_id, &teacher.subject, &teacher.messages)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Here's the output: %v", teacher)

	uifs := http.FileServer(http.Dir("./public"))

	http.Handle("/", uifs)
	cssfs := http.FileServer(http.Dir("./public/style.css"))
	http.Handle("/style.css", cssfs)
	err = http.ListenAndServe(":3333", nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Listening on port 3333")
}

func AddMessage(conn *pgx.Conn, id int32) {
    message := map[string]string{"behaviour": "good"}
    messages_json, _ := json.Marshal(message) 

    _, err := conn.Exec(context.Background(), `UPDATE teacher SET messages = messages || $2 WHERE teacher.id = $1`, id ,[][]byte{messages_json})

    if err != nil {
        panic(err)
    }
}
