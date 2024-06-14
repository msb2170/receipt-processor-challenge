package models

// Item represents an item in a purchase receipt
type Item struct {
    ShortDescription string  `json:"shortDescription"`
    Price            float64 `json:"price"`
}
