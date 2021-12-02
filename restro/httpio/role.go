package httpio

type UserRole struct {
	ID        uint64             `gorm:"primaryKey"`
	RoleName  string             `json:"rolename"`
	Resources []ResourcesToRoles `json:"resources" gorm:"foreignKey:RoleID;References:ID;"`
}

type ResourcesToRoles struct {
	RoleID     uint64 `json:"roleid" gorm:"primaryKey;autoIncrement:false"`
	ResourceID uint64 `json:"resourceid"  gorm:"primaryKey;autoIncrement:false"`
}
