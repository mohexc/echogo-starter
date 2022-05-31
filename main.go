package main

import (
	"echogo/modules/post"
	"echogo/modules/user"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()

	app.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	apiGroup := app.Group("/api/v1")
	user.Route(apiGroup)
	post.Route(apiGroup)

	app.Logger.Fatal(app.Start(":5500"))
}
