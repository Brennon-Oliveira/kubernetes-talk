package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Bem-vindo ao MiniShop Frontend!")
	})

	log.Println("Frontend rodando na porta 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
