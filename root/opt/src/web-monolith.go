package main

import (
    "fmt"
    "net/http"
    "os"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Web-monolith version", os.Getenv("VERSION"))
}

func paymentsHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Payments module version", os.Getenv("VERSION"))
}

func inventoryHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Inventory module version", os.Getenv("VERSION"))
}

func shippingHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Shipping module version", os.Getenv("VERSION"))
}

func main() {
	fmt.Println("Starting web-monolith version", os.Getenv("VERSION"))
    http.HandleFunc("/payments/", paymentsHandler)
    http.HandleFunc("/inventory/", inventoryHandler)
    http.HandleFunc("/shipping/", shippingHandler)
    http.HandleFunc("/", rootHandler)
    http.ListenAndServe(":8080", nil)
}
