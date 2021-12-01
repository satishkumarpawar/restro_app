package auth

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type AuthDetails struct {
	//AuthUuid string
	UserId uint64
	RoleId uint64
}

func CreateToken(authD AuthDetails) (string, error) {
	//Creating Access Token
	os.Setenv("API_SECRET", "jdnfksdmfksd") //this should be in an env file

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	//claims["auth_uuid"] = authD.AuthUuid
	claims["user_id"] = authD.UserId
	claims["role_id"] = authD.RoleId
	claims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("API_SECRET")))
}

func TokenValid(r *http.Request) error {
	token, err := VerifyToken(r)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return nil
}

func VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//does this token conform to "SigningMethodHMAC" ?
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

//get the token from the request body
func ExtractToken(r *http.Request) string {
	keys := r.URL.Query()
	token := keys.Get("token")
	if token != "" {
		return token
	}
	bearToken := r.Header.Get("Authorization")
	//normally Authorization the_token_xxx
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func ExtractTokenAuth(r *http.Request) (*AuthDetails, error) {
	token, err := VerifyToken(r)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims) //the token claims should conform to MapClaims
	if ok && token.Valid {
		/*authUuid, ok := claims["auth_uuid"].(string) //convert the interface to string
		if !ok {
			return nil, err
		}*/
		userId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)

		if err != nil {
			return nil, err
		}
		roleId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["role_id"]), 10, 64)
		if err != nil {
			return nil, err
		}
		return &AuthDetails{
			//AuthUuid: authUuid,
			UserId: userId,
			RoleId: roleId,
		}, nil
	}
	return nil, err
}
