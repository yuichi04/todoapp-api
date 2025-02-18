package controller

import (
	"fmt"
	"net/http"
	"os"
	"time"
	"todoapp-api/model"
	"todoapp-api/usecase"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type IUserController interface {
	SignUp(c echo.Context) error
	LogIn(c echo.Context) error
	LogOut(c echo.Context) error
	CsrfToken(c echo.Context) error
	GetMe(c echo.Context) error
}

type userController struct {
	uu usecase.IUserUsecase
}

func NewUserController(uu usecase.IUserUsecase) IUserController {
	return &userController{uu}
}

func (uc *userController) SignUp(c echo.Context) error {
	fmt.Println("SignUp handler called")

	// リクエストボディをUser構造体に直接バインド
	var user model.User
	if err := c.Bind(&user); err != nil {
		fmt.Printf("Error binding request: %v\n", err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	fmt.Printf("Request body: %+v\n", user)

	// アカウントを作成し、クライアントに返すユーザー情報を取得
	userRes, err := uc.uu.SignUp(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, userRes)
}

func (uc *userController) LogIn(c echo.Context) error {
	user := model.User{}

	// コンテキストを構造体にバインド
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// ログイン処理を実行し、JWTを取得
	tokenString, err := uc.uu.LogIn(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// Cookie構造体を用意し、各種値を設定
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = tokenString
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.Path = "/"
	cookie.Domain = os.Getenv("API_DOMAIN")
	cookie.Secure = true
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteNoneMode

	// Cookie構造体をCookieにセット
	c.SetCookie(cookie)

	return c.NoContent(http.StatusOK)
}

func (uc *userController) LogOut(c echo.Context) error {
	// ログアウト処理（Cookieを削除）
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = ""
	cookie.Expires = time.Now()
	cookie.Path = "/"
	cookie.Domain = os.Getenv("API_DOMAIN")
	cookie.Secure = true
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteNoneMode
	c.SetCookie(cookie)
	return c.NoContent(http.StatusOK)
}

func (uc *userController) CsrfToken(c echo.Context) error {
	token := c.Get("csrf").(string)
	return c.JSON(http.StatusOK, echo.Map{
		"csrf_token": token,
	})
}

func (uc *userController) GetMe(c echo.Context) error {
	// JWTからユーザーIDを取得
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	// ユーザー情報を取得
	userRes, err := uc.uu.GetUser(uint(userId.(float64)))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, userRes)
}
