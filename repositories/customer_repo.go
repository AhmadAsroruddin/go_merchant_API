package repositories

import (
    "encoding/json"
    "io/ioutil"
    "os"
    "models"
)

func GetCustomers() ([]models.Customer, error) {
    var customers []models.Customer
    jsonFile, err := os.Open("data/customers.json")
    if err != nil {
        return customers, err
    }
    defer jsonFile.Close()

    byteValue, _ := ioutil.ReadAll(jsonFile)
    json.Unmarshal(byteValue, &customers)
    return customers, nil
}

func UpdateCustomers(customers []models.Customer) error {
    file, err := json.MarshalIndent(customers, "", " ")
    if err != nil {
        return err
    }
    return ioutil.WriteFile("data/customers.json", file, 0644)
}
