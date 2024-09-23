package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type User struct {
	ID    int    `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

func main() {
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		user := User{ID: 1, Nome: "Maria", Email: "maria@example.com"}
		json.NewEncoder(w).Encode(user)
	})

	log.Println("Auth Service rodando na porta 8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
