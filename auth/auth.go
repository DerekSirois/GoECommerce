package auth

import (
	"GoECommerce/utils"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"time"
)

const secretKey = "supersecretkey"

type UserClaim struct {
	jwt.RegisteredClaims
	ID       int
	UserName string
	IsAdmin  bool
}

func CreateJWTToken(id int, name string, isAdmin bool) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaim{
		RegisteredClaims: jwt.RegisteredClaims{ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 10))},
		ID:               id,
		UserName:         name,
		IsAdmin:          isAdmin,
	})

	// Create the actual JWT token
	signedString, err := token.SignedString([]byte(secretKey))

	if err != nil {
		return "", fmt.Errorf("error creating signed string: %v", err)
	}

	return signedString, nil
}

func VerifyJWT(endpointHandler func(writer http.ResponseWriter, request *http.Request), adminOnly bool) http.HandlerFunc {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if request.Header["Token"] != nil {
			var jwtToken = request.Header["Token"][0]
			var userClaim UserClaim
			token, err := jwt.ParseWithClaims(jwtToken, &userClaim, func(token *jwt.Token) (interface{}, error) {
				return []byte(secretKey), nil
			})
			if err != nil {
				utils.Respond(writer, request, &utils.Response{Msg: err.Error()}, http.StatusBadRequest)
				return
			}
			if !token.Valid {
				utils.Respond(writer, request, &utils.Response{Msg: "Invalid token"}, http.StatusBadRequest)
				return
			}
			if adminOnly && !userClaim.IsAdmin {
				utils.Respond(writer, request, &utils.Response{Msg: "Admin only"}, http.StatusUnauthorized)
				return
			}
			endpointHandler(writer, request)
		} else {
			utils.Respond(writer, request, &utils.Response{Msg: "Missing token"}, http.StatusBadRequest)
			return
		}
	})
}
