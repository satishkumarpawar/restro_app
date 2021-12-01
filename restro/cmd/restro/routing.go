package main

import (
	"fmt"
	"log"
	"net/http"
	"restro/internal/dish"
	"restro/internal/resource"
	"restro/internal/restaurant"
	"restro/internal/role"
	"restro/internal/user"
	"strings"

	"restro/cmd/config/globalvars"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

func Middleware(h http.HandlerFunc, act string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var jwtKey = []byte("jdnfksdmfksd")
		// We can obtain the session token from the requests cookies, which come with every request
		c, err := r.Cookie("token")
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
		tknStr := c.Value
		if tknStr == "" {
			// If the cookie is not set, return an unauthorized status
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// Initialize a new instance of `Claims`
		claims := jwt.MapClaims{}

		// Parse the JWT string and store the result in `claims`.
		// Note that we are passing the key in this method as well. This method will return an error
		// if the token is invalid (if it has expired according to the expiry time we set on sign in),
		// or if the signature does not match
		tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if !tkn.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// (END) The code up-till this point is the same as the first part of the `Welcome` route

		// We ensure that a new token is not issued until enough time has elapsed
		// In this case, a new token will only be issued if the old token is within
		// 30 seconds of expiry. Otherwise, return a bad request status
		/*if time.Unix(claims["exp"].(int64), 0).Sub(time.Now()) > 30*time.Second {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Now, create a new token for the current use, with a renewed expiration time
		expirationTime := time.Now().Add(5 * time.Minute)
		claims["exp"] = expirationTime.Unix()
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString(jwtKey)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// Set the new token as the users `token` cookie
		http.SetCookie(w, &http.Cookie{
			Name:    "token",
			Value:   tokenString,
			Expires: expirationTime,
		})*/
		for key, val := range claims {
			fmt.Printf("Key: %v, value: %v\n", key, val)
		}

		user_id := uint64(claims["user_id"].(float64))
		role_id := uint64(claims["role_id"].(float64))
		//check user permission for resources
		//for admin
		res := resource.GetResourceByName("all")
		urole := role.GetResourceToRole(res.ID, role_id)
		//for other users
		res2 := resource.GetResourceByName(act)
		urole2 := role.GetResourceToRole(res2.ID, role_id)

		log.Println("all role_id", urole.RoleID)
		log.Println("act role_id", urole2.RoleID)

		if urole.RoleID != role_id && urole2.RoleID != role_id && act != "LogoutUser" {
			http.Error(w, "This action is not permitted to you as user role!", http.StatusUnauthorized)
			return
		}

		//set user type
		urole3 := role.GetRoleByRoleID(role_id)
		log.Println("user_role", urole3)
		role_name := "user"
		if strings.Contains(strings.ToLower(urole3.RoleName), "sub") && strings.Contains(strings.ToLower(urole3.RoleName), "admin") {
			role_name = "subadmin"
		} else if strings.Contains(strings.ToLower(urole3.RoleName), "admin") {
			role_name = "admin"
		}

		globalvars.Glbvs.UserID = user_id
		globalvars.Glbvs.RoleID = role_id
		globalvars.Glbvs.UserRoleName = role_name

		//log.Println("role_name", role_name)

		/*reso := role.GetResourcesToRoles(uint64(role_id))
		var resids []uint64
		for _, res := range reso {
			resids = append(resids, res.ResourceID)
		}
		resources := resource.GetResourceByIDs(resids)
		*/
		//log.Println(resids)
		//log.Println(resources)
		log.Println("middleware", r.URL)
		h.ServeHTTP(w, r)
	}
}

func HandlerRouting() {
	r := mux.NewRouter()

	r.HandleFunc("/login", user.LoginUser).Methods("POST")
	r.HandleFunc("/logout", Middleware(user.LogoutUser, "LogoutUser")).Methods("GET")
	r.HandleFunc("/users", Middleware(user.GetUsers, "GetUsers")).Methods("GET")
	r.HandleFunc("/user/{uid}", Middleware(user.GetUserByID, "GetUserByID")).Methods("GET")
	r.HandleFunc("/user", Middleware(user.CreateUser, "CreateUser")).Methods("POST")
	r.HandleFunc("/user/{uid}", Middleware(user.UpdateUser, "UpdateUser")).Methods("PUT")
	r.HandleFunc("/user/{uid}", Middleware(user.DeleteUser, "DeleteUser")).Methods("DELETE")

	r.HandleFunc("/restaurants", Middleware(restaurant.GetRestaurants, "GetRestaurants")).Methods("GET")
	r.HandleFunc("/restaurant/{rid}", Middleware(restaurant.GetRestaurantByID, "GetRestaurantByID")).Methods("GET")
	r.HandleFunc("/restaurant", Middleware(restaurant.CreateRestaurant, "CreateRestaurant")).Methods("POST")
	r.HandleFunc("/restaurant/{rid}", Middleware(restaurant.UpdateRestaurant, "UpdateRestaurant")).Methods("PUT")
	r.HandleFunc("/restaurant/{rid}", Middleware(restaurant.DeleteRestaurant, "DeleteRestaurant")).Methods("DELETE")

	r.HandleFunc("/dishes", Middleware(dish.GetDishes, "GetDishes")).Methods("GET")
	r.HandleFunc("/dishes/{rid}", Middleware(dish.GetDishesByRestaurantID, "GetDishesByRestaurantID")).Methods("GET")
	r.HandleFunc("/dish/{did}", Middleware(dish.GetDishByID, "GetDishByID")).Methods("GET")
	r.HandleFunc("/dish", Middleware(dish.CreateDish, "CreateDish")).Methods("POST")
	r.HandleFunc("/dish/{did}", Middleware(dish.UpdateDish, "UpdateDish")).Methods("PUT")
	r.HandleFunc("/dish/{did}", Middleware(dish.DeleteDish, "DeleteDish")).Methods("DELETE")

	r.HandleFunc("/resources", Middleware(resource.GetResources, "GetResources")).Methods("GET")
	r.HandleFunc("/resource/{rcid}", Middleware(resource.GetResourceByID, "GetResourceByID")).Methods("GET")
	r.HandleFunc("/resource", Middleware(resource.CreateResource, "CreateResource")).Methods("POST")
	r.HandleFunc("/resource/{rcid}", Middleware(resource.UpdateResource, "UpdateResource")).Methods("PUT")
	r.HandleFunc("/resource/{rcid}", Middleware(resource.DeleteResource, "DeleteResource")).Methods("DELETE")

	r.HandleFunc("/roles", Middleware(role.GetRoles, "GetRoles")).Methods("GET")
	r.HandleFunc("/role/{roid}", Middleware(role.GetRoleByID, "GetRoleByID")).Methods("GET")
	r.HandleFunc("/role", Middleware(role.CreateRole, "CreateRole")).Methods("POST")
	r.HandleFunc("/role/{roid}", Middleware(role.UpdateRole, "UpdateRole")).Methods("PUT")
	r.HandleFunc("/role/{roid}", Middleware(role.DeleteRole, "DeleteRole")).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8081", r))

}
