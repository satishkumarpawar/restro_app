package httpio

type UserRole struct {
	ID        uint64             `json:"_id"`
	RoleName  string             `json:"rolename"`
	Resources []ResourcesToRoles `json:"resources" gorm:"foreignKey:ID references:RoleID"`
}

type ResourcesToRoles struct {
	RoleID     uint64 `json:"roleid" gorm:"primaryKey;autoIncrement:false"`
	ResourceID uint64 `json:"resourceid"`
}
