package restaurant

import (
	"encoding/json"
	"net/http"
	"restro/cmd/config/db"
	"restro/cmd/config/globalvars"
	"restro/cmd/config/helper"

	"github.com/gorilla/mux"

	"restro/httpio"
)

var DB = db.DConn()

func CreateRestaurant(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var restaurant httpio.Restaurant
	json.NewDecoder(r.Body).Decode(&restaurant)
	if restaurant.Latitude == 0 || restaurant.Longitude == 0 {
		geolocation := helper.GetGeocode(restaurant.Address)
		//restaurant.Address = geolocation.Address
		restaurant.Latitude = geolocation.Location.Lat
		restaurant.Longitude = geolocation.Location.Lng
	}
	restaurant.AdminID = globalvars.Glbvs.UserID
	DB.Create(&restaurant)
	json.NewEncoder(w).Encode((restaurant))

}

func GetRestaurants(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var restaurants []httpio.Restaurant
	if globalvars.Glbvs.UserRoleName == "subadmin" {
		DB.Where("admin_id = ?", globalvars.Glbvs.UserID).Find(&restaurants)
	} else {
		DB.Find(&restaurants)
	}

	json.NewEncoder(w).Encode((restaurants))
}

func GetRestaurantByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var restaurant httpio.Restaurant
	if globalvars.Glbvs.UserRoleName == "subadmin" {
		DB.Where("admin_id = ? AND id = ?", globalvars.Glbvs.UserID, mux.Vars(r)["rid"]).First(&restaurant)
	} else {
		DB.Where("id = ?", mux.Vars(r)["rid"]).First(&restaurant)
	}
	//DB.First(&restaurant, mux.Vars(r)["rid"])
	json.NewEncoder(w).Encode((restaurant))
}

func UpdateRestaurant(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var restaurant httpio.Restaurant
	if globalvars.Glbvs.UserRoleName == "subadmin" {
		DB.Where("admin_id = ? AND id = ?", globalvars.Glbvs.UserID, mux.Vars(r)["rid"]).First(&restaurant)
	} else {
		DB.Where("id = ?", mux.Vars(r)["rid"]).First(&restaurant)
	}
	//DB.First(&restaurant, mux.Vars(r)["rid"])
	json.NewDecoder(r.Body).Decode(&restaurant)
	if restaurant.Latitude == 0 || restaurant.Longitude == 0 {
		geolocation := helper.GetGeocode(restaurant.Address)
		//restaurant.Address = geolocation.Address
		restaurant.Latitude = geolocation.Location.Lat
		restaurant.Longitude = geolocation.Location.Lng
	}
	DB.Save(&restaurant)
	json.NewEncoder(w).Encode((restaurant))
}

func DeleteRestaurant(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var restaurant httpio.Restaurant
	if globalvars.Glbvs.UserRoleName == "subadmin" {
		DB.Where("admin_id = ? AND id = ?", globalvars.Glbvs.UserID, mux.Vars(r)["rid"]).Delete(&restaurant)
	} else {
		DB.Where("id = ?", mux.Vars(r)["rid"]).Delete(&restaurant)
	}
	//DB.Delete(&restaurant, mux.Vars(r)["rid"])
	json.NewEncoder(w).Encode("Restaurant is deleted !!")
}
