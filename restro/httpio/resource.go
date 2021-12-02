package httpio

type UserResource struct {
	ID           uint64 `gorm:"primaryKey"`
	ResourceName string `json:"resourcename"`
}
