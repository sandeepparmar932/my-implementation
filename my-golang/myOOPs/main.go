package main

import (
	"fmt"
	"myDesign/controller"
	"myDesign/router"
	"myDesign/service"
	"net/http"
)

func main() {
	userService := service.NewUserService()
	userController := controller.NewUserController(userService)
	appRouter := router.NewRouter(userController)
	appRouter.InitRoutes()

	port := 8080
	fmt.Printf("Server running on Port %d\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
