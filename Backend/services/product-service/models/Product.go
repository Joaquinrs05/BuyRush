package models

type Product struct {
    ID          uint    `gorm:"primaryKey" json:"id"`
    Name        string  `json:"name"`
    Price       float64 `json:"price"`
    Description string  `json:"description"`
    Photo       string  `json:"photo"`
    CreatedAt   string  `json:"created_at"`
}