package main

import (
	"fmt"

	"github.com/google/wire"
)

type myRepo struct {
}

func (my *myRepo) getUser(id int) string {
	return fmt.Sprintf("Get User with id - %d ", id)
}

type myService struct {
	Repo *myRepo
}

func (myService *myService) getUserByID(id int) string {
	return myService.Repo.getUser(id)
}

func NewUserRepo() *myRepo {
	return &myRepo{}
}

func NewUserService(repo *myRepo) *myService {
	return &myService{
		Repo: repo,
	}
}

func initialize() *myService {
	wire.Build(NewUserRepo, NewUserService)
	return &myService{}
}
func main() {
	userSer := initialize()
	fmt.Print(userSer.getUserByID(1))
}
