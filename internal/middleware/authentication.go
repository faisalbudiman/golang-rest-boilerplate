package middleware

import (
	"fmt"
	"os"
	"strings"

	"golang-rest-boilerplate/internal/abstraction"
	res "golang-rest-boilerplate/pkg/util/response"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func Authentication(next echo.HandlerFunc) echo.HandlerFunc {
	var (
		jwtKey = os.Getenv("JWT_KEY")
	)

	return func(c echo.Context) error {
		authToken := c.Request().Header.Get("Authorization")
		if authToken == "" {
			return res.ErrorBuilder(&res.ErrorConstant.Unauthorized, nil).Send(c)
		}

		splitToken := strings.Split(authToken, "Bearer ")
		token, err := jwt.Parse(splitToken[1], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method :%v", token.Header["alg"])
			}

			return []byte(jwtKey), nil
		})

		if !token.Valid || err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.Unauthorized, err).Send(c)
		}

		var id int
		destructID := token.Claims.(jwt.MapClaims)["id"]
		if destructID != nil {
			id = int(destructID.(float64))
		} else {
			id = 0
		}

		var name string
		destructName := token.Claims.(jwt.MapClaims)["name"]
		if destructName != nil {
			name = destructName.(string)
		} else {
			name = ""
		}

		var email string
		destructEmail := token.Claims.(jwt.MapClaims)["email"]
		if destructEmail != nil {
			email = destructEmail.(string)
		} else {
			email = ""
		}

		var phone string
		destructPhone := token.Claims.(jwt.MapClaims)["phone"]
		if destructPhone != nil {
			phone = destructEmail.(string)
		} else {
			phone = ""
		}

		var userType string
		destructUserType := token.Claims.(jwt.MapClaims)["user_type"]
		if destructPhone != nil {
			userType = destructUserType.(string)
		} else {
			userType = ""
		}

		cc := c.(*abstraction.Context)
		cc.Auth = &abstraction.AuthContext{
			ID:       id,
			Name:     name,
			Email:    email,
			Phone:    phone,
			UserType: userType,
		}

		return next(cc)
	}
}
