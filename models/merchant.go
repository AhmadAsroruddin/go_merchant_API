package models

type Merchant struct {
    ID      string `json:"id"`
    Name    string `json:"name"`
    Balance int    `json:"balance"`
}
