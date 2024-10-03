package handlers

import (
    "net/http"
    "encoding/json"
    "services"
    "models"
)

func PaymentHandler(w http.ResponseWriter, r *http.Request) {
    var req struct {
        Customer models.Customer `json:"customer"`
        Amount   int             `json:"amount"`
    }
    json.NewDecoder(r.Body).Decode(&req)

    err := services.Payment(req.Customer, req.Amount)
    if err != nil {
        http.Error(w, err.Error(), http.StatusUnauthorized)
        return
    }

    w.WriteHeader(http.StatusOK)
}
