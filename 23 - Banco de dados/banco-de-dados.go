package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	stringConexao := "golang:golang@tcp(localhost:3366)/devbook?charset=utf8&parseTime=True&loc=Local"

	db, erro := sql.Open("mysql", stringConexao)
	checkError(erro)
	defer db.Close()

	erro = db.Ping()
	checkError(erro)

	fmt.Println("Conexão está aberta")

	linhas, erro := db.Query("select * from usuarios")
	checkError(erro)
	defer linhas.Close()
}
