package resource

import (
	"encoding/json"
	"net/http"

	"restro/cmd/config/db"
	"restro/httpio"

	"github.com/gorilla/mux"
)

var DB = db.DConn()

func CreateResource(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var resource httpio.UserResource
	json.NewDecoder(r.Body).Decode(&resource)
	DB.Create(&resource)
	json.NewEncoder(w).Encode((resource))

}

func GetResources(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var resources []httpio.UserResource
	DB.Find(&resources)
	json.NewEncoder(w).Encode((resources))
}

func GetResourceByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var resource httpio.UserResource
	DB.First(&resource, mux.Vars(r)["rcid"])
	json.NewEncoder(w).Encode((resource))
}

func UpdateResource(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var resource httpio.UserResource
	DB.First(&resource, mux.Vars(r)["rcid"])
	json.NewDecoder(r.Body).Decode(&resource)
	DB.Save(&resource)
	json.NewEncoder(w).Encode((resource))
}

func DeleteResource(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var resource httpio.UserResource
	DB.Delete(&resource, mux.Vars(r)["rcid"])
	json.NewEncoder(w).Encode("resource is deleted !!")
}

func GetResourceByIDs(resourceid []uint64) []httpio.UserResource {
	var resources []httpio.UserResource
	DB.Find(&resources, resourceid)
	return resources
}
func GetResourceByName(resourcename string) httpio.UserResource {
	var resource httpio.UserResource
	DB.Where("resource_name = ?", resourcename).First(&resource)
	//DB.Find(&resource, resourcename)
	return resource
}
