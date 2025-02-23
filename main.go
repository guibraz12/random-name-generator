package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

func main() {
	// Registra os handlers antes de iniciar o servidor
	http.HandleFunc("/random-name", randomNameHandler)
	http.Handle("/", http.FileServer(http.Dir("static")))
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("static/css"))))

	fmt.Println("Servidor rodando em http://localhost:1233")
	err := http.ListenAndServe(":1233", nil)
	if err != nil {
		fmt.Printf("Erro ao iniciar o servidor: %v\n", err)
	}
}

func randomNameHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*") // Adiciona o cabeçalho CORS

	firstNames := []string{"Alice", "Bob", "Charlie", "Diana"}
	lastNames := []string{"Smith", "Johnson", "Williams", "Brown"}

	firstName := firstNames[rand.Intn(len(firstNames))]
	lastName := lastNames[rand.Intn(len(lastNames))]

	fmt.Fprintf(w, "%s %s", firstName, lastName)
}
