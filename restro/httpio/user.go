package httpio

import "gorm.io/gorm"

type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	RoleID   uint64 `json:"roleid"`
}

type User struct {
	gorm.Model
	ID        uint64          `gorm:"primaryKey"`
	AdminID   uint64          `json:"adminid"`
	Name      string          `json:"name"`
	Email     string          `json:"email"`
	Password  string          `json:"password"`
	Address   []UserAddress   `json:"address"`
	Privilege []UserPrivilege `json:"privilege"`
}

type UserAddress struct {
	ID        uint64  `gorm:"primaryKey"`
	UserID    uint64  `json:"userid"`
	Address   string  `json:"address"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type UserPrivilege struct {
	UserID uint64 `json:"userid" gorm:"primaryKey;autoIncrement:false"`
	RoleID uint64 `json:"roleid" gorm:"primaryKey;autoIncrement:false"`
}
