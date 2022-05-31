package post

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Route(route *echo.Group) {

	route.POST("/posts", func(ctx echo.Context) (err error) {
		post := new(Post)
		if err = ctx.Bind(post); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		newPost := createPost(*post)
		return ctx.JSON(http.StatusOK, newPost)
	})

	route.GET("/posts", func(ctx echo.Context) error {
		posts := findPosts()
		return ctx.JSON(http.StatusOK, posts)
	})

	route.GET("/posts/:id", func(ctx echo.Context) error {
		id := ctx.Param("id")
		post := findPost(id)
		return ctx.JSON(http.StatusOK, post)
	})

	route.PATCH("/posts/:id", func(ctx echo.Context) error {
		id := ctx.Param("id")
		bodyPost := new(Post)
		fmt.Println("bodyPost", bodyPost)
		updatePost := updatePost(id, *bodyPost)
		return ctx.JSON(http.StatusOK, updatePost)
	})

	route.DELETE("/posts/:id", func(ctx echo.Context) error {

		return ctx.String(http.StatusOK, "delete user")
	})
}
