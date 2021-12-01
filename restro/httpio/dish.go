package httpio

import "gorm.io/gorm"

type Dish struct {
	gorm.Model
	ID           uint64 `gorm:"primaryKey"`
	RestaurantID uint64 `json:"restaurantid"`
	AdminID      uint64 `json:"adminid"`
	DishName     string `json:"dishname"`
	Price        string `json:"price"`
	Description  string `json:"description"`
}
