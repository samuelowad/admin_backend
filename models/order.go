package models

type Order struct {
	Id         uint        `json:"id"`
	FirstName  string      `json:"-"`
	LastName   string      `json:"-"`
	Name       string      `gorm:"-" json:"name"`
	Email      string      `json:"email"`
	UpdatedAt  string      `json:"updated_at"`
	CreatedAt  string      `json:"created_at"`
	OrderItems []OrderItem `json:"order_items" gorm:"foreignKey:OrderId"`
}

type OrderItem struct {
	Id           uint    `json:"id"`
	OrderId      uint    `json:"order_id"`
	ProductTitle string  `json:"product_title"`
	Price        float32 `json:"price"`
	Quantity     uint    `json:"quantity"`
}
