package restaurant

import (
	"encoding/json"
	"net/http"
	"restro/cmd/config/db"
	"restro/cmd/config/globalvars"
	"restro/cmd/config/helper"

	"github.com/gorilla/mux"

	"restro/httpio"
	"restro/internal/user"
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
	//fmt.Println(restaurants)
	var restaurants1 []httpio.Restaurants

	useraddress := user.GetUserAddress(globalvars.Glbvs.UserID)
	//fmt.Println(restaurants)
	for _, restaurant := range restaurants {
		var restaurant1 httpio.Restaurants
		restaurant1.Restaurant = restaurant
		var distance1 []httpio.Distance
		for _, address := range useraddress {
			var distance httpio.Distance
			distance.UserAddress = address.Address
			if address.Latitude != 0 && address.Longitude != 0 && restaurant.Latitude != 0 && restaurant.Longitude != 0 {
				distance.DistanceToRestaurant = helper.Distance(address.Latitude, address.Longitude, restaurant.Latitude, restaurant.Longitude, "K")
			}
			distance1 = append(distance1, distance)
		}
		restaurant1.Distance = distance1
		restaurants1 = append(restaurants1, restaurant1)
	}
	json.NewEncoder(w).Encode((restaurants1))
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
	var restaurant1 httpio.Restaurants
	restaurant1.Restaurant = restaurant
	useraddress := user.GetUserAddress(globalvars.Glbvs.UserID)
	var distance []httpio.Distance
	for k, address := range useraddress {
		distance[k].UserAddress = address.Address
		if address.Latitude != 0 && address.Longitude != 0 && restaurant.Latitude != 0 && restaurant.Longitude != 0 {
			distance[k].DistanceToRestaurant = helper.Distance(address.Latitude, address.Longitude, restaurant.Latitude, restaurant.Longitude, "K")

		}
	}
	restaurant1.Distance = distance

	json.NewEncoder(w).Encode((restaurant1))
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
