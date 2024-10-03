package services

import (
    "repositories"
    "models"
    "errors"
)

func Payment(customer models.Customer, amount int) error {
    if !customer.LoggedIn {
        return errors.New("customer is not logged in")
    }

    merchants, err := repositories.GetMerchants()
    if err != nil {
        return err
    }

    // Simulasi transfer ke merchant
    for i, merchant := range merchants {
        if merchant.ID == "1" {
            merchants[i].Balance += amount
            repositories.UpdateMerchants(merchants)
            return nil
        }
    }

    return errors.New("merchant not found")
}
