package main

import (
    "fmt"
    "net/http"
    "os"
)

func paymentsHandler(w http.ResponseWriter, r *http.Request) {
    version, _ := os.Getenv("VERSION")
    fmt.Fprintln(w, "Payments module version ", version)
    r.Write(w)
}

func inventoryHandler(w http.ResponseWriter, r *http.Request) {
    version, _ := os.Getenv("VERSION")
    fmt.Fprintln(w, "Inventory module version ", version)
    r.Write(w)
}

func shippingHandler(w http.ResponseWriter, r *http.Request) {
    version, _ := os.Getenv("VERSION")
    fmt.Fprintln(w, "Shipping module version ", version)
    r.Write(w)
}

func main() {
    http.HandleFunc("/payments/", paymentsHandler)
    http.HandleFunc("/inventory/", inventoryHandler)
    http.HandleFunc("/shipping/", shippingHandler)
    http.ListenAndServe(":8080", nil)
}
