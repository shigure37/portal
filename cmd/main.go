/*================================================================
*   FileName    : main.go
*   Author      : mavinyeh
*   Date        : 2019-05-12
*   Description :
================================================================*/
package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
	"portal/web/app/handlers"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", hello)
	e.GET("/login", handlers.LoginHandler)

	e.Logger.Fatal(e.Start(":8080"))

}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "hello")
}
