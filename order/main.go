package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Order struct {
	ID        int     `json:"id"`
	ProductID int     `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Total     float64 `json:"total"`
}

func main() {
	http.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		order := Order{ID: 1, ProductID: 1, Quantity: 2, Total: 21.98}
		json.NewEncoder(w).Encode(order)
	})

	log.Println("Order Service rodando na porta 8083")
	log.Fatal(http.ListenAndServe(":8083", nil))
}
