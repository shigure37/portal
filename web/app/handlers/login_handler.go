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

type Todo struct {
	Title string
	Done  bool
}
type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

func LoginHandler(c echo.Context) error {

	//return c.String(http.StatusOK, "login")
	return c.Render(http.StatusOK, "index.html", map[string]interface{}{
		"name":  "Dolly!",
		"title": "shiguremoe",
	})
}
