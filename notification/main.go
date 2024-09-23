package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/notify", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Notificação enviada com sucesso!")
	})

	log.Println("Notification Service rodando na porta 8085")
	log.Fatal(http.ListenAndServe(":8085", nil))
}
