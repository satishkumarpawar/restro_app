package httpio

import "gorm.io/gorm"

type Restaurant struct {
	gorm.Model
	ID             uint64  `gorm:"primaryKey"`
	AdminID        uint64  `json:"adminid"`
	RestaurantName string  `json:"restaurantname"`
	Address        string  `json:"address"`
	Latitude       float64 `json:"latitude"`
	Longitude      float64 `json:"longitude"`
}
