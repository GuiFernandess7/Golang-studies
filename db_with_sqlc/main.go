package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"os"

	database "github.com/GuiFernandess7/db_with_sqlc/sqlc"
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

	encodedPassword := url.QueryEscape(password)
	connStr := fmt.Sprintf("postgres://postgres:%s@localhost:5432/dunder_mifflin?sslmode=disable", encodedPassword)
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	ctx := context.Background()
	queries := database.New(db)

	employees, err := queries.ListEmployees(ctx)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(employees)

}