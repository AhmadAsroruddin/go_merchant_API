package repositories

import (
    "encoding/json"
    "io/ioutil"
    "os"
    "models"
)

func GetMerchants() ([]models.Merchant, error) {
    var merchants []models.Merchant
    jsonFile, err := os.Open("data/merchants.json")
    if err != nil {
        return merchants, err
    }
    defer jsonFile.Close()

    byteValue, _ := ioutil.ReadAll(jsonFile)
    json.Unmarshal(byteValue, &merchants)
    return merchants, nil
}

func UpdateMerchants(merchants []models.Merchant) error {
    file, err := json.MarshalIndent(merchants, "", "  ")
    if err != nil {
        return err
    }
    return ioutil.WriteFile("data/merchants.json", file, 0644)
}
