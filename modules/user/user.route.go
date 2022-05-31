package user

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Route(route *echo.Group) {

	route.POST("/users", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "create user")
	})

	route.GET("/users", func(ctx echo.Context) error {
		users := findUsers()
		fmt.Println(users)
		return ctx.JSON(http.StatusOK, users)
	})

	route.GET("/users/:id", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "get user id")
	})

	route.PATCH("/users/:id", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "update user")
	})

	route.DELETE("/users/:id", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "delete user")
	})
}
