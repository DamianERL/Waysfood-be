package models

import "time"

type Cart struct {
	ID        int       `json:"id" gorm:"primary_key:auto_increment" `
	SubTotal  int       `json:"sub_total" gorm:"type:int"`
	Shipping  int       `json:"shipping" gorm:"type:int"`
	Total     int       `json:"total" gorm:"type:int"`
	ProductID []int     `json:"product_id" gorm:"type:int"`
	Product   []Product `json:"product" gorm:"many2many:cart_products"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
