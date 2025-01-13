package main

import (
	"fmt"
	"todoapp-api/controller"
	"todoapp-api/db"
	"todoapp-api/repository"
	"todoapp-api/router"
	"todoapp-api/usecase"
)

func main() {
	db := db.NewDB()
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userController := controller.NewUserController(userUsecase)
	e := router.NewRouter(userController)

	// サーバー設定を明示的に指定
	e.Server.Addr = ":8080"

	fmt.Println("Starting server on :8080")
	if err := e.Start("0.0.0.0:8080"); err != nil {
		e.Logger.Fatal(err)
	}
}
