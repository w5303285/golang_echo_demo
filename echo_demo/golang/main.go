package main

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"myapp/controller/user"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World1!")
	})

	e.GET("/users/:id", user.GetUser)

	e.Logger.Fatal(e.Start(":8088"))
}

