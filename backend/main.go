package main

import (
    "encoding/json"
    "log"
    "net/http"
)

type Product struct {
    ID    int     `json:"id"`
    Name  string  `json:"name"`
    Price float64 `json:"price"`
}

var products = []Product{
    {ID: 1, Name: "Футболка", Price: 499},
    {ID: 2, Name: "Кепка", Price: 299},
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(products)
}

func main() {
    http.HandleFunc("/api/products", productsHandler)
    log.Println("✅ Backend запущен на http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
