// Auto-generated GORM models from provided schema and relations
package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `gorm:"not null;unique"`
	Password string `gorm:"not null"`
	Role     string `gorm:"not null"`
	Profile  *Profile
	Cart     Cart
	Orders   []Order
}

type Profile struct {
	ID        uint   `gorm:"primaryKey"`
	UserID    uint   `gorm:"unique;not null"`
	User      User   `gorm:"constraint:OnDelete:CASCADE"`
	Name      string `gorm:"not null"`
	Phone     string
	Address   string
	AvatarURL string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// type User struct {
// 	ID        uint   `gorm:"primaryKey"`
// 	Name      string `gorm:"not null"`
// 	Email     string `gorm:"unique;not null"`
// 	Password  string `gorm:"not null"`
// 	Role      string `gorm:"default:user;not null"`
// 	Phone     string
// 	CreatedAt time.Time
// 	UpdatedAt time.Time
// 	DeletedAt gorm.DeletedAt `gorm:"index"`
// 	Cart      Cart
// 	Orders    []Order
// }

type Cart struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint `gorm:"uniqueIndex;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	CartItems []CartItem
}

type Category struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"unique;not null"`
	ParentID  *uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Children  []Category     `gorm:"foreignKey:ParentID"`
}

type Product struct {
	ID               uint   `gorm:"primaryKey"`
	ProductCode      int    `gorm:"unique;not null"`
	ProductType      string `gorm:"not null"`
	Name             string `gorm:"not null"`
	Description      string
	Price            float64 `gorm:"type:numeric(10,2);not null"`
	Stock            int     `gorm:"default:0"`
	ImageURL         string
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt `gorm:"index"`
	BookID           *uint
	Book             *Book
	LearningSupplyID *uint
	LearningSupply   *LearningSupply
	OfficeSupplyID   *uint
	OfficeSupply     *OfficeSupply
	Categories       []Category  `gorm:"many2many:book_categories;"`
	BookImages       []BookImage `gorm:"foreignKey:ProductID"` // ✅ บอกว่า BookImage ใช้ ProductID เป็น FK()
}

type Book struct {
	gorm.Model
	ProductID uint `gorm:"uniqueIndex"` // One-to-one
	Author    string
	ISBN      string
	Discount  float64
}

type LearningSupply struct {
	gorm.Model
	ProductID uint `gorm:"uniqueIndex"` // One-to-one
	Brand     string
	Material  string
}

type OfficeSupply struct {
	gorm.Model
	ProductID uint `gorm:"uniqueIndex"` // One-to-one
	Color     string
	Size      string
}

type Order struct {
	ID        uint    `gorm:"primaryKey"`
	UserID    uint    `gorm:"not null"`
	Total     float64 `gorm:"type:numeric(10,2);not null"`
	Status    string  `gorm:"default:pending;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Items     []OrderItem
}

type OrderItem struct {
	ID              uint `gorm:"primaryKey"`
	OrderID         uint `gorm:"not null"`
	ProductID       *uint
	GroupID         *uint
	Quantity        int     `gorm:"default:1;not null"`
	PriceAtPurchase float64 `gorm:"type:numeric(10,2);not null"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`
}

type CartItem struct {
	ID        uint `gorm:"primaryKey"`
	CartID    uint `gorm:"not null"`
	ProductID *uint
	GroupID   *uint
	Quantity  int `gorm:"default:1;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type BookCategory struct {
	ProductID  uint `gorm:"primaryKey"`
	CategoryID uint `gorm:"primaryKey"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

type BookImage struct {
	gorm.Model
	ProductID uint
	ImageURL  string
}

type Group struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"not null"`
	EduLevel    string `gorm:"not null"`
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Products    []GroupProduct
}

type GroupProduct struct {
	ID        uint `gorm:"primaryKey"`
	GroupID   uint `gorm:"not null"`
	ProductID uint `gorm:"not null"`
	Quantity  int  `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
