package main

import (
	"fmt"
	"todoapp-api/controller"
	"todoapp-api/db"
	"todoapp-api/repository"
	"todoapp-api/router"
	"todoapp-api/usecase"
	"todoapp-api/validator"
)

func main() {
	db := db.NewDB()
	userValidator := validator.NewUserValidator()
	taskValidator := validator.NewTaskValidator()
	userRepository := repository.NewUserRepository(db)
	taskRepository := repository.NewTaskRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository, userValidator)
	taskUsecase := usecase.NewTaskUsecase(taskRepository, taskValidator)
	userController := controller.NewUserController(userUsecase)
	taskController := controller.NewTaskController(taskUsecase)
	e := router.NewRouter(userController, taskController)

	// サーバー設定を明示的に指定
	e.Server.Addr = ":8080"

	fmt.Println("Starting server on :8080")
	if err := e.Start("0.0.0.0:8080"); err != nil {
		e.Logger.Fatal(err)
	}
}
