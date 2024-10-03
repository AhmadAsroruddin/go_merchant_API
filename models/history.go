package models

import "time"

type History struct {
    ID        string    `json:"id"`
    Customer  string    `json:"customer"`
    Action    string    `json:"action"`
    Timestamp time.Time `json:"timestamp"`
}
