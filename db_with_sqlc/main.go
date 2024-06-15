package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"os"

	database "github.com/GuiFernandess7/db_with_sqlc/db"
	"github.com/GuiFernandess7/db_with_sqlc/use_cases"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main(){
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}

	db, err := setDBConnection(os.Getenv("CONN_STRING"), os.Getenv("SECRET"))

	if err != nil {
		log.Fatal(err)
		return
	}

	defer db.Close()

	ctx := context.Background()
	db_repository := database.New(db)
	sex := "F"

	employee, err := use_cases.GetEmployeesBy(50000, 70000, &sex, db_repository, ctx)
	//use_cases.SearchEmployeeBySex("M", db_repository, ctx)
	//GetEmployeesBySalaryRange(10000, 70000, db_repository, ctx)

	if err != nil {
		fmt.Println(err)
		return
	}

	for i := range employee {
		fmt.Println(employee[i].FirstName.String, "|", employee[i].Salary.Int32)
	}
}

func setDBConnection(conn_str string, password string) (*sql.DB, error){
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

