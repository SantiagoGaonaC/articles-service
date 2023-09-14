package models

type User struct {
	ID        uint       `gorm:"primarykey"`
	Username  string     `gorm:"unique;not null"`
	Password  string     `gorm:"not null"`
	Favorites []Favorite `gorm:"foreignKey:UserID"`
}

type Product struct {
	ID         uint   `gorm:"primarykey"`
	Vendor     string `gorm:"not null"`
	SellerName string `gorm:"not null"`
	Rating     float32
	ImageURL   string
	Favorites  []Favorite `gorm:"foreignKey:ProductID"`
}

// Products
type Favorite struct {
	ID        uint `gorm:"primarykey"`
	UserID    uint
	ProductID uint
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Token    string `json:"token"`
}
