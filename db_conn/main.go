package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)


func main(){
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}

	password := os.Getenv("SECRET")
	if password == "" {
		log.Fatal("A variável de ambiente SECRET não está definida")
	}

	connStr := fmt.Sprintf("postgres://postgres:%s@localhost:5432/dunder_mifflin", password)

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Ping realizado com sucesso!")
	}

	employees, err := Listemployees(db)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(employees)

	defer db.Close()
}

