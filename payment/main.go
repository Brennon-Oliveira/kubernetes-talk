package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Payment struct {
	ID     int     `json:"id"`
	Order  int     `json:"order_id"`
	Status string  `json:"status"`
	Amount float64 `json:"amount"`
}

func main() {
	http.HandleFunc("/payments", func(w http.ResponseWriter, r *http.Request) {
		payment := Payment{ID: 1, Order: 1, Status: "Aprovado", Amount: 21.98}
		json.NewEncoder(w).Encode(payment)
	})

	log.Println("Payment Service rodando na porta 8084")
	log.Fatal(http.ListenAndServe(":8084", nil))
}
