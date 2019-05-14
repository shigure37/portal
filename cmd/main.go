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
	"html/template"
	"io"
	"net/http"
	"portal/web/app/handlers"
)

type TemplateRender struct {
	templates *template.Template
}

func (t *TemplateRender) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}
	return t.templates.ExecuteTemplate(w, name, data)
}
func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Static("./static"))
	renderer := &TemplateRender{
		templates: template.Must(template.ParseGlob("./template/*.html")),
	}
	e.Renderer = renderer

	e.GET("/", hello)
	e.GET("/login", handlers.LoginHandler).Name = "xxx"

	e.Logger.Fatal(e.Start(":8080"))

}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "hello")
}
