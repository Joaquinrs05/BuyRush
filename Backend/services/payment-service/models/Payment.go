package models

type Payment struct {
    ID         uint   `gorm:"primaryKey" json:"id"`
    OrderID    uint   `json:"order_id"`
    Method     string `json:"method"`
    Amount     float64 `json:"amount"`
    PaidAt     string `json:"paid_at"`
}