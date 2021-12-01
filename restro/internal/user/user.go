package user

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"restro/cmd/config/auth"
	"restro/cmd/config/db"
	"restro/cmd/config/globalvars"
	"restro/cmd/config/helper"
	"restro/httpio"
)

var DB = db.DConn()

func LoginUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user httpio.LoginUser
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid json provided", http.StatusUnprocessableEntity)
		return
	}

	var u httpio.User
	DB.Where("email= ? AND password = ?", user.Email, user.Password).Find(&u)
	u.Privilege = GetUserPrivilege(u.ID)
	//compare the user from the request, with the one we defined:
	if u.Email != user.Email || u.Password != user.Password {
		http.Error(w, "Please provide valid login details", http.StatusUnauthorized)
		return
	}

	RoleVerified := false
	for _, prvi := range u.Privilege {
		if prvi.RoleID == user.RoleID {
			RoleVerified = true
			break
		}
	}
	//fmt.Println(user.RoleID)
	//fmt.Println(u)
	//fmt.Println(RoleVerified)
	if !RoleVerified {
		http.Error(w, "This user role is not permitted", http.StatusUnauthorized)
		return
	}

	var authD auth.AuthDetails
	authD.UserId = u.ID
	authD.RoleId = user.RoleID
	//authD.AuthUuid = "session"

	token, err := auth.CreateToken(authD)

	if err != nil {
		http.Error(w, "Please try to login later", http.StatusForbidden)
		return
	}

	expirationTime := time.Now().Add(15 * time.Minute)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   token,
		Expires: expirationTime,
	})

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(token); err != nil {
		//panic(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

}

func LogoutUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	_, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			// If the cookie is not set, return an unauthorized status
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		// For any other type of error, return a bad request status
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get the JWT string from the cookie
	/*token := c.Value

	expirationTime := time.Now().Add(-20 * time.Minute)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   token,
		Expires: expirationTime,
	})
	*/
	c := &http.Cookie{
		Name:     "token",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		Expires:  time.Unix(0, 0),
		HttpOnly: true,
	}

	http.SetCookie(w, c)

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode("Successfully logged out"); err != nil {
		//panic(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

}

func CreateUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var user httpio.User
	json.NewDecoder(r.Body).Decode(&user)

	//fmt.Println(user.Address)
	//fmt.Println(user.Privilege)
	//DB.Create(&user)
	for k, address := range user.Address {
		if address.Latitude == 0 || address.Longitude == 0 {
			geolocation := helper.GetGeocode(address.Address)
			//user.Address[k].Address = geolocation.Address
			user.Address[k].Latitude = geolocation.Location.Lat
			user.Address[k].Longitude = geolocation.Location.Lng
		}

	}
	user.AdminID = globalvars.Glbvs.UserID
	DB.Create(&user)
	//addr := user.Address
	/*DB.Table("user_addresses").Create(&addr)
	//DB.Create(&addr)
	prvi := user.Privilege
	DB.Table("user_privileges").Create(&prvi)
	//DB.Create(&prvi)*/

	json.NewEncoder(w).Encode(user)

}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var users []httpio.User
	if globalvars.Glbvs.UserRoleName == "subadmin" {
		DB.Where("admin_id = ?", globalvars.Glbvs.UserID).Find(&users)
	} else {
		DB.Find(&users)
	}

	for k, user := range users {
		users[k].Address = GetUserAddress(user.ID)
		users[k].Privilege = GetUserPrivilege(user.ID)

	}
	json.NewEncoder(w).Encode((users))
}

func GetUserByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user httpio.User
	if globalvars.Glbvs.UserRoleName == "subadmin" {
		err := DB.Where("admin_id = ? AND id=?", globalvars.Glbvs.UserID, mux.Vars(r)["uid"]).Preload("Address").Preload("Privilege").Find(&user).Error
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	} else {
		err := DB.Where("id = ?", mux.Vars(r)["uid"]).Find(&user).Error
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	/*DB.First(&user, mux.Vars(r)["uid"])
	user.Address = GetUserAddress(user.ID)
	user.Privilege = GetUserPrivilege(user.ID)
	*/
	json.NewEncoder(w).Encode((user))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user httpio.User
	if globalvars.Glbvs.UserRoleName == "subadmin" {
		DB.Where("admin_id = ? AND id = ?", globalvars.Glbvs.UserID, mux.Vars(r)["uid"]).First(&user)
	} else {
		DB.Where("id = ?", mux.Vars(r)["uid"]).First(&user)
	}
	//DB.First(&user, mux.Vars(r)["uid"])
	json.NewDecoder(r.Body).Decode(&user)
	DB.Save(&user)
	json.NewEncoder(w).Encode((user))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user httpio.User
	if globalvars.Glbvs.UserRoleName == "subadmin" {
		DB.Where("admin_id = ? AND id = ?", globalvars.Glbvs.UserID, mux.Vars(r)["uid"]).Delete(&user)
	} else {
		DB.Where("id = ?", mux.Vars(r)["uid"]).Delete(&user)
	}
	//DB.Delete(&user, mux.Vars(r)["uid"])
	json.NewEncoder(w).Encode("User is deleted !!")
}

func GetUserAddress(userid uint64) []httpio.UserAddress {
	var addr []httpio.UserAddress
	DB.Where("user_id = ?", userid).Find(&addr)
	return addr
}
func GetUserPrivilege(userid uint64) []httpio.UserPrivilege {
	var priv []httpio.UserPrivilege
	DB.Where("user_id = ?", userid).Find(&priv)
	//DB.Find(&priv, userid)
	return priv
}
