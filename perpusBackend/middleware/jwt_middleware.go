package middleware

import (
	"project/constants"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func CreateToken(adminId int, name string) (string, error) {

	claims := jwt.MapClaims{}
	claims["adminId"] = adminId
	claims["name"] = name
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(constants.SECRET_JWT))
}

func ExtractTokenAmdminId(e echo.Context) int {
	admin := e.Get("admin").(*jwt.Token)
	if admin.Valid {
		claims := admin.Claims.(jwt.MapClaims)
		adminId := claims["adminId"].(float64)
		return int(adminId)
	}
	return 0
}
