package main

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/kmontag42/idle-of-building/routes"
)

type Templates struct {
	templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func renderer() *Templates {
	return &Templates{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Renderer = renderer()

	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index", nil)
	})

        e.GET("/ws", routes.WebSockets)

	e.POST("/upload", routes.UploadFile)
        e.POST("/upload-export-string", routes.UploadExportString)
        e.POST("/run-map", routes.RunMapForCharacter)
	e.Logger.Fatal(e.Start(":42069"))
}
