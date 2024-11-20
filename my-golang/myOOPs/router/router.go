package router

import (
	"myDesign/controller"
	"net/http"
)

type Router struct {
	userController *controller.UserController
}

func NewRouter(uc *controller.UserController) *Router {
	return &Router{userController: uc}
}

func (r *Router) InitRoutes() {
	http.HandleFunc("/add", r.userController.AddUser)
	http.HandleFunc("/get", r.userController.GetUser)
	http.HandleFunc("/update", r.userController.UpdateUser)
	http.HandleFunc("/delete", r.userController.DeleteUser)
	http.HandleFunc("/getAll", r.userController.GetAllUsers)

}
