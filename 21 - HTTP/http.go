package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Olá mundo!"))
}

func usuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Carregando pagina de usuários!"))
}

func main() {
	// HTTP é um procolo de comunicação - base da comunicação WEB

	// Cliente (faz requisição) - Servidor (processa a requisição e envia resposta)

	// Request - Response

	// Rota
	// URI - Identificador do Recurso
	// Método - GET, POST, PUT, DELETE

	http.HandleFunc("/home", home)

	http.HandleFunc("/usuarios", usuarios)

	log.Fatal(http.ListenAndServe(":5000", nil))
}
