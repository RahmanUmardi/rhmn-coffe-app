package entity

import "time"

type Transaction struct {
	Transaction_id            string    `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"transaction_id"`
	Order_id                  string    `gorm:"type:uuid;not null" json:"order_id"`
	Payment_gateway_reference string    `gorm:"type:varchar(255)" json:"payment_gateway_reference"`
	Payment_status            string    `gorm:"type:varchar(50);not null" json:"payment_status"` // success, pending, failed
	Payment_date              time.Time `gorm:"type:timestamp;default:current_timestamp" json:"payment_date"`
}