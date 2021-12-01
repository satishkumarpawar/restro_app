package dish

import (
	"encoding/json"
	"net/http"

	"restro/cmd/config/db"
	"restro/cmd/config/globalvars"
	"restro/httpio"

	"github.com/gorilla/mux"
)

var DB = db.DConn()

func CreateDish(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var dish httpio.Dish
	json.NewDecoder(r.Body).Decode(&dish)
	dish.AdminID = globalvars.Glbvs.UserID
	DB.Create(&dish)
	json.NewEncoder(w).Encode((dish))

}

func GetDishes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var dishes []httpio.Dish
	if globalvars.Glbvs.UserRoleName == "subadmin" {
		DB.Where("admin_id = ?", globalvars.Glbvs.UserID).Find(&dishes)
	} else {
		DB.Find(&dishes)
	}

	json.NewEncoder(w).Encode((dishes))
}

func GetDishesByRestaurantID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var dishes []httpio.Dish
	if globalvars.Glbvs.UserRoleName == "subadmin" {
		DB.Where("admin_id = ? AND restaurant_id = ?", globalvars.Glbvs.UserID, mux.Vars(r)["rid"]).Find(&dishes)
	} else {
		DB.Where("restaurant_id = ?", mux.Vars(r)["rid"]).Find(&dishes)
	}
	json.NewEncoder(w).Encode((dishes))
}

func GetDishByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var dish httpio.Dish
	if globalvars.Glbvs.UserRoleName == "subadmin" {
		DB.Where("admin_id = ? AND id = ?", globalvars.Glbvs.UserID, mux.Vars(r)["did"]).First(&dish)
	} else {
		DB.Where("id = ?", mux.Vars(r)["did"]).First(&dish)
	}
	//DB.First(&dish, mux.Vars(r)["did"])
	json.NewEncoder(w).Encode((dish))
}

func UpdateDish(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var dish httpio.Dish
	if globalvars.Glbvs.UserRoleName == "subadmin" {
		DB.Where("admin_id = ? AND id = ?", globalvars.Glbvs.UserID, mux.Vars(r)["did"]).First(&dish)
	} else {
		DB.Where("id = ?", mux.Vars(r)["did"]).First(&dish)
	}
	//DB.First(&dish, mux.Vars(r)["did"])
	json.NewDecoder(r.Body).Decode(&dish)
	DB.Save(&dish)
	json.NewEncoder(w).Encode((dish))
}

func DeleteDish(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var dish httpio.Dish
	if globalvars.Glbvs.UserRoleName == "subadmin" {
		DB.Where("admin_id = ? AND id = ?", globalvars.Glbvs.UserID, mux.Vars(r)["did"]).Delete(&dish)
	} else {
		DB.Where("id = ?", mux.Vars(r)["did"]).Delete(&dish)
	}
	//DB.Delete(&dish, mux.Vars(r)["did"])
	json.NewEncoder(w).Encode("Dish is deleted !!")
}
