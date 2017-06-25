package main

import (
	"net/http"
	"time"
	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/dgrijalva/jwt-go"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"name": "iguchi",
			"nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
		})

		// Sign and get the complete encoded token as a string using the secret
		tokenString, err := token.SignedString( []byte("fadsfasdfadsfadf") )

		fmt.Println(tokenString, err)

		return c.String(http.StatusOK, tokenString)
	})


	// Restricted group
	r := e.Group("/admin")
	r.Use(middleware.JWT([]byte("fadsfasdfadsfadf")))
	r.GET("", func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		name := claims["name"].(string)

		fmt.Println(name)

		return c.String(http.StatusOK, name)
	})


	e.Logger.Fatal(e.Start(":1323"))
}
