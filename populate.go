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

var schema = `
CREATE TABLE IF NOT EXISTS school (
    id int generated always as identity,
    school_name text,
    primary key (id)
);
CREATE TABLE IF NOT EXISTS teacher (
    id int generated always as identity,
    first_name text,
    last_name text,
    email text,
    school_id int REFERENCES school(id),
    subject text[],
    messages JSON[],
    primary key (id)
);
CREATE TABLE IF NOT EXISTS student (
    id int generated always as identity,
    first_name text,
    last_name text,
    email text,
    school_id int REFERENCES school(id),
    primary key (id)
);
`
func populateDB(conn *pgx.Conn) {
    _, err := conn.Exec(context.Background(), "DROP TABLE IF EXISTS teacher; DROP TABLE IF EXISTS student; DROP TABLE IF EXISTS school;")

	if err != nil {
		panic(err)
	}

	_, err = conn.Exec(context.Background(), schema)


	messages := map[string]string{"kind": "bad"}
	messages_json, _ := json.Marshal(messages)

	_, err = conn.Exec(context.Background(), `INSERT INTO school (school_name) VALUES ('lavender school')`)
	_, err = conn.Exec(context.Background(), `INSERT INTO school (school_name) VALUES ('turkish school')`)

	_, err = conn.Exec(context.Background(),
		`INSERT INTO teacher (email, first_name, last_name, school_id, subject, messages) 
    VALUES ($1, $2, $3, $4, $5, $6)`,
		"missingno@gmail.com", "Missing", "No", 2,
		[]string{"Computer Science"}, [][]byte{messages_json})
	if err != nil {
		panic(err)
	}
	_, err = conn.Exec(context.Background(), `UPDATE teacher SET messages = messages || $1 WHERE teacher.id = 1`, [][]byte{messages_json})

}


