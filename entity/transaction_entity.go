package entity

import "time"

type Transaction struct {
	Transaction_id string    `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"transaction_id"`
	Order_id       string    `gorm:"type:uuid;not null" json:"order_id"`
	Payment_status string    `gorm:"type:varchar(50);not null" json:"payment_status"`
	Payment_date   time.Time `gorm:"type:timestamp;default:current_timestamp" json:"payment_date"`
}
