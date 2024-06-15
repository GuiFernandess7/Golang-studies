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

	db, err := setDBConnection("postgres://postgres:%s@localhost:5432/dunder_mifflin?sslmode=disable")

	if err != nil {
		log.Fatal(err)
		return
	}

	defer db.Close()

	ctx := context.Background()
	handler := database.New(db)

	firstName := sql.NullString{
        String: "Michael",
        Valid:  true,
    }

	employee, err := handler.FindEmployeeByName(ctx, firstName)
    if err != nil {
        log.Fatal(err)
    }

	//employees, err := queries.ListAllEmployeesBy(ctx, 3)

	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println(employee)
}

func setDBConnection(conn_str string) (*sql.DB, error){
	password := os.Getenv("SECRET")
	if password == "" {
		log.Fatal("A variável de ambiente SECRET não está definida")
	}

	encodedPassword := url.QueryEscape(password)
	connStr := fmt.Sprintf(conn_str, encodedPassword)
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		return nil, err
	}

	return db, nil
}

