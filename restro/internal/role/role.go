package role

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"restro/cmd/config/db"
	"restro/httpio"
)

var DB = db.DConn()

func CreateRole(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var role httpio.UserRole
	json.NewDecoder(r.Body).Decode(&role)
	DB.Create(&role)
	json.NewEncoder(w).Encode((role))

}

func GetRoles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var roles []httpio.UserRole
	err := DB.Preload("Resources").Find(&roles).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	/*DB.Find(&roles)
	for k, role := range roles {
		roles[k].Resources = GetResourcesToRoles(role.ID)
	}*/
	json.NewEncoder(w).Encode((roles))
}

func GetRoleByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var role httpio.UserRole
	DB.First(&role, mux.Vars(r)["roid"])
	role.Resources = GetResourcesToRoles(role.ID)
	json.NewEncoder(w).Encode((role))
}

func UpdateRole(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var role httpio.UserRole
	DB.First(&role, mux.Vars(r)["roid"])
	json.NewDecoder(r.Body).Decode(&role)
	DB.Save(&role)
	json.NewEncoder(w).Encode((role))
}

func DeleteRole(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var role httpio.UserRole
	DB.Delete(&role, mux.Vars(r)["roid"])
	json.NewEncoder(w).Encode("Role is deleted !!")
}

func GetRoleByRoleID(roleid uint64) httpio.UserRole {
	var role httpio.UserRole
	DB.Where("id = ?", roleid).First(&role)
	return role
}
func GetResourcesToRoles(roleid uint64) []httpio.ResourcesToRoles {
	var role []httpio.ResourcesToRoles
	DB.Where("role_id = ?", roleid).Find(&role)
	return role
}

func GetResourceToRole(resid uint64, roleid uint64) httpio.ResourcesToRoles {
	var role httpio.ResourcesToRoles
	DB.Where("resource_id = ? AND role_id = ?", resid, roleid).First(&role)
	return role
}
