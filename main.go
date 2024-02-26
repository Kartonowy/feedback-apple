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
)

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

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	router.HandleFunc("/schools/", func(w http.ResponseWriter, r *http.Request) {
        encoder := json.NewEncoder(w)
        encoder.SetEscapeHTML(false)
        encoder.Encode(GetSchools(conn))
	})

	router.HandleFunc("/addstudent/", func(w http.ResponseWriter, r *http.Request) { addStudent(w, r, conn) })

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

func GetSchools(conn *pgx.Conn) string {
	rows, err := conn.Query(context.Background(), "SELECT id, school_name FROM school")
	output := make(map[string]string)

    var htmled_output string

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
        htmled_output += fmt.Sprintf("<option value='%v'>%v</option>", school.school_name, school.school_name)
	}
    fmt.Println(htmled_output)
	return htmled_output
}

func addStudent(w http.ResponseWriter, r *http.Request, conn *pgx.Conn) {
	r.ParseForm()
	var query int32
	defer r.Body.Close()

	row := conn.QueryRow(context.Background(), fmt.Sprintf("SELECT id FROM school WHERE school_name = '%s'", r.PostForm.Get("school")))

	row.Scan(&query)
	_, err := conn.Exec(context.Background(),
		fmt.Sprintf("INSERT INTO student (email, first_name, last_name, school_id) VALUES ('%s', '%s', '%s', '%d')",
			r.PostForm.Get("email"),
			r.PostForm.Get("first_name"),
			r.PostForm.Get("last_name"),
			query))

	if err != nil {
		log.Printf("ERROR WHILE INSERTING STUDENTS INTO DB \n Query was: INSERT INTO student (email, first_name, last_name, school_id) VALUES ('%s', '%s', '%s', '%d') \n",
			r.PostForm.Get("email"),
			r.PostForm.Get("first_name"),
			r.PostForm.Get("last_name"),
			query)
		log.Fatalf("%v", err)
	}
	fmt.Printf("query: %v\n", query)
}

func addTeacher(w http.ResponseWriter, r *http.Request, conn *pgx.Conn) {

}
