/*================================================================
*   FileName    : login_handler.go
*   Author      : mavinyeh
*   Date        : 2019-05-12
*   Description :
================================================================*/
package handlers

import (
	"github.com/labstack/echo"
	"net/http"
)

func LoginHandler(c echo.Context) error {
	return c.String(http.StatusOK, "login")
}
