package entity

import "time"

type Order struct {
	Order_id       string        `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"order_id"`
	User_id        string        `gorm:"type:uuid;not null" json:"user_id"`
	Total_price    float64       `gorm:"type:decimal(10,2);not null" json:"total_price"`
	Payment_method string        `gorm:"type:varchar(50);not null" json:"payment_method"`
	Status         string        `gorm:"type:varchar(50);default:pending" json:"status"`
	Created_at     time.Time     `gorm:"autoCreateTime" json:"created_at"`
	Updated_at     time.Time     `gorm:"autoUpdateTime" json:"updated_at"`
	Order_items    []OrderItem   `gorm:"foreignKey:OrderID" json:"order_items"`
	Transaction    []Transaction `gorm:"foreignKey:OrderID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"-"`
}

type OrderItem struct {
	Order_item_id string  `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"order_item_id"`
	Order_id      string  `gorm:"type:uuid;not null" json:"order_id"`
	Product_id    string  `gorm:"type:uuid;not null" json:"product_id"`
	Quantity      int     `gorm:"not null" json:"quantity"`
	Subtotal      float64 `gorm:"type:decimal(10,2);not null" json:"subtotal"`
}
