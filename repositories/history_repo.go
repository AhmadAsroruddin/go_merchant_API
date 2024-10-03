package repositories

import (
    "encoding/json"
    "io/ioutil"
    "os"
    "models"
)

func GetHistory() ([]models.History, error) {
    var history []models.History
    jsonFile, err := os.Open("data/history.json")
    if err != nil {
        return history, err
    }
    defer jsonFile.Close()

    byteValue, _ := ioutil.ReadAll(jsonFile)
    json.Unmarshal(byteValue, &history)
    return history, nil
}

func UpdateHistory(history []models.History) error {
    file, err := json.MarshalIndent(history, "", "  ")
    if err != nil {
        return err
    }
    return ioutil.WriteFile("data/history.json", file, 0644)
}
