package main

import (
	"app/codegen"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

var schema = `
CREATE TABLE IF NOT EXISTS teachers (
    id int generated always as identity,
    first_name text,
    last_name text,
    email text,
    password text,
    class_codes text[],
    primary key (id)
);
CREATE TABLE IF NOT EXISTS classes (
    id text,
    responses JSON[],
    primary key (id)
);
`
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
    _, err = conn.Exec(context.Background(), "DROP TABLE IF EXISTS teachers; DROP TABLE IF EXISTS classes;")

	if err != nil {
		panic(err)
	}

	_, err = conn.Exec(context.Background(), schema)

    for i := 1; i < 3; i++ {
        class := codegen.RandSeq(6)

        _, err = conn.Exec(context.Background(),
        `INSERT INTO teachers (email, first_name, last_name, class_codes) 
        VALUES ($1, $2, $3, $4)`,
        "missingno@gmail.com", "Iron", "Hands", []string{class})

        if err != nil {
            panic(err)
        }
        _, err = conn.Exec(context.Background(), `INSERT INTO classes (id, responses) VALUES ($1, $2)`, class, message_constructor("the more", "the better"))

    }
     //	_, err = conn.Exec(context.Background(), `UPDATE teacher SET messages = messages || $1 WHERE teacher.id = 1`, message_constructor("bhvr", "fine"))
     // select * from classes c where c.id= any (select unnest(t.class_codes) from teachers t);

}

func message_constructor(key string, value string) [][]byte {
	messages := map[string]string{key: value}
	messages_json, _ := json.Marshal(messages)

    return [][]byte{messages_json}
}

