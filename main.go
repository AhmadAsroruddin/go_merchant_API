package main

import (
    "net/http"
    "handlers"
)

func main() {
    http.HandleFunc("/login", handlers.LoginHandler)
    http.HandleFunc("/payment", handlers.PaymentHandler)
    http.HandleFunc("/logout", handlers.LogoutHandler)

    // Jalankan server di port 8080
    http.ListenAndServe(":8080", nil)
}
