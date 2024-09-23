package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Product struct {
	ID    int     `json:"id"`
	Nome  string  `json:"nome"`
	Preco float64 `json:"preco"`
}

func main() {
	http.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		products := []Product{
			{ID: 1, Nome: "Produto A", Preco: 10.99},
			{ID: 2, Nome: "Produto B", Preco: 20.50},
		}
		json.NewEncoder(w).Encode(products)
	})

	log.Println("Product Service rodando na porta 8082")
	log.Fatal(http.ListenAndServe(":8082", nil))
}
