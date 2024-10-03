package handlers

import (
    "net/http"
    "encoding/json"
    "services"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
    var req struct {
        Username string `json:"username"`
        Password string `json:"password"`
    }
    json.NewDecoder(r.Body).Decode(&req)

    customer, err := services.Login(req.Username, req.Password)
    if err != nil {
        http.Error(w, err.Error(), http.StatusUnauthorized)
        return
    }

    json.NewEncoder(w).Encode(customer)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
    var customer models.Customer
    json.NewDecoder(r.Body).Decode(&customer)

    err := services.Logout(customer)
    if err != nil {
        http.Error(w, err.Error(), http.StatusUnauthorized)
        return
    }

    w.WriteHeader(http.StatusOK)
}
