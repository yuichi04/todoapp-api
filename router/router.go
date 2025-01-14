package router

import (
	"fmt"
	"net/http"
	"os"
	"todoapp-api/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(uc controller.IUserController) *echo.Echo {
	e := echo.New()

	// CORSのミドルウェアを設定
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{os.Getenv("CLIENT_URL")},
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
			echo.HeaderAccessControlAllowHeaders,
			echo.HeaderXCSRFToken,
		},
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE"},
		AllowCredentials: true,
	}))

	// CSRFのミドルウェアを設定
	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		CookiePath:     "/",
		CookieDomain:   os.Getenv("API_DOMAIN"),
		CookieHTTPOnly: true,
		CookieSameSite: http.SameSiteNoneMode,
		// CookieSameSite: http.SameSiteDefaultMode,
		// CookieMaxAge:   60,
	}))

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
	e.GET("/csrf", uc.CsrfToken)
	return e
}
