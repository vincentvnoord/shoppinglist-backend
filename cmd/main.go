package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"shopping-list-backend/internal/handlers"
	"shopping-list-backend/pkg/database"

	"github.com/gorilla/mux"
)

var db *sql.DB

func main() {
	err := database.LoadEnv()
	if err != nil {
		fmt.Println("Error loading env")
	}

	db, err = database.Connection()
	if err != nil {
		fmt.Println("Error connecting to database")
	}

	r := mux.NewRouter()
	r.HandleFunc("/list/{public_code}", handlers.ListGet).Methods("GET")
	r.HandleFunc("/list", handlers.ListPost).Methods("POST")
	r.HandleFunc("/list", handlers.ListPut).Methods("PUT")

	r.HandleFunc("/products/{list_id}", handlers.ProductsGet).Methods("GET")
	r.HandleFunc("/product", handlers.ProductPost).Methods("POST")
	r.HandleFunc("/product", handlers.ProductPut).Methods("PUT")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	defer db.Close()

	fmt.Println("Server started on port " + port)

	http.ListenAndServe(":"+port, r)
}
