package entity

type Product struct {
	Product_id   string  `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id_product"`
	Product_name string  `json:"product_name"`
	Price        float64 `json:"price"`
}

type UpdateProduct struct {
	Product_name string  `json:"product_name"`
	Price        float64 `json:"price"`
}
