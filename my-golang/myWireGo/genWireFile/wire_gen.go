// +build wireinject

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"fmt"
)

// Injectors from main.go:

func initialize() *myService {
	mainMyRepo := NewUserRepo()
	mainMyService := NewUserService(mainMyRepo)
	return mainMyService
}

// main.go:

type myRepo struct {
}

func (my *myRepo) getUser(id int) string {
	return fmt.Sprintf("Get User with id - %d ", id)
}

type myService struct {
	Repo *myRepo
}

func (myService2 *myService) getUserByID(id int) string {
	return myService2.Repo.getUser(id)
}

func NewUserRepo() *myRepo {
	return &myRepo{}
}

func NewUserService(repo *myRepo) *myService {
	return &myService{
		Repo: repo,
	}
}

func main() {
	userSer := initialize()
	fmt.Print(userSer.getUserByID(1))
}
