package entity

type User struct {
	User_id  string `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id_user"`
	Username string `json:"name"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type UpdateUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}
