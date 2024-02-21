package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
)

func main() {
    router := chi.NewRouter()
    router.Use(middleware.Logger)

    router.Get("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("welcome"))
    })
    
    fs := http.FileServer(http.Dir("static"))
    router.Handle("/static/*", http.StripPrefix("/static/", fs))

    err := http.ListenAndServe(":3333", router)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Listening on port 3333")
}
