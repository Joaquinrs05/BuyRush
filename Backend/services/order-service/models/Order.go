package models

type Order struct {
    ID        uint   `gorm:"primaryKey" json:"id"`
    UserID    uint   `json:"user_id"`
    Status    string `json:"status"`
    Total     float64 `json:"total"`
    CreatedAt string `json:"created_at"`
}

type OrderProduct struct {
    ID        uint `gorm:"primaryKey" json:"id"`
    OrderID   uint `json:"order_id"`
    ProductID uint `json:"product_id"`
    Quantity  int  `json:"quantity"`
}