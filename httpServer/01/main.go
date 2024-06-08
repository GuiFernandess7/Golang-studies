package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)

type Product struct {
	ID string
	Name string
	Price float64
}

func main(){
	product := Product{
		ID: "1",
		Name: "Product 1",
		Price: 100.00,
	}

	e := echo.New()
	e.POST("/products", createProduct)
	e.Logger.Fatal(e.Start(":3000"))

	fmt.Println(product.Name)
	http.HandleFunc("/", homeHandler)
	http.ListenAndServe(":3000", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello World"))
}

func createProduct(c echo.Context) error {
    product := Product{}
    if err := c.Bind(&product); err != nil {
        fmt.Println("Error binding product:", err)
        return err
    }
    fmt.Println("Received product:", product)

    err := SaveProduct(product)
    if err != nil {
        fmt.Println("Error saving product:", err)
        return err
    }
    fmt.Println("Product saved successfully:", product)
    return c.JSON(http.StatusCreated, product)
}

func SaveProduct(product Product) error {
	db, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		return err
	}
	defer db.Close()
	stmt, err := db.Prepare("INSERT INTO products(id, name, price) VALUES(?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(product.ID, product.Name, product.Price)
	if err != nil {
		return err
	}
	return nil
}