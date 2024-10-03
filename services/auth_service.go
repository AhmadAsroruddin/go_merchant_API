package services

import (
    "repositories"
    "models"
    "errors"
)

func Login(username, password string) (models.Customer, error) {
    customers, err := repositories.GetCustomers()
    if err != nil {
        return models.Customer{}, err
    }

    for i, customer := range customers {
        if customer.Username == username && customer.Password == password {
            customers[i].LoggedIn = true
            repositories.UpdateCustomers(customers)
            return customer, nil
        }
    }
    return models.Customer{}, errors.New("invalid credentials")
}

func Logout(customer models.Customer) error {
    customers, err := repositories.GetCustomers()
    if err != nil {
        return err
    }

    for i, c := range customers {
        if c.ID == customer.ID {
            customers[i].LoggedIn = false
            repositories.UpdateCustomers(customers)
            return nil
        }
    }
    return errors.New("customer not found")
}
