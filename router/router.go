package router

import (
	"fmt"
	"net/http"
	"todoapp-api/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(uc controller.IUserController) *echo.Echo {
	e := echo.New()

	e.Debug = true

	// すべてのリクエストをログ出力
	e.Use(middleware.Logger())

	// リクエストの詳細をログ出力するミドルウェア
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			fmt.Printf("Request: %s %s\n", c.Request().Method, c.Request().URL.Path)
			return next(c)
		}
	})

	e.GET("/test", func(c echo.Context) error {
		return c.String(http.StatusOK, "Test endpoint working!")
	})

	e.POST("/signup", uc.SignUp)
	e.POST("/login", uc.LogIn)
	e.POST("/logout", uc.LogOut)

	return e
}
