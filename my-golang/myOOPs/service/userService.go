package service

import (
	"errors"
	"myDesign/model"
)

type UserService struct {
	users  []model.User
	nextID int
}

func NewUserService() *UserService {
	return &UserService{users: []model.User{}, nextID: 1}
}

func (s *UserService) AddUser(user *model.User) {
	user.SetID(s.nextID)
	s.nextID++
	s.users = append(s.users, *user)
}

func (s *UserService) GetUser(id int) (*model.User, error) {
	for _, user := range s.users {
		if user.GetID() == id {
			return &user, nil
		}
	}
	return nil, errors.New("User not found")
}

func (s *UserService) UpdateUser(updatedUser *model.User) error {

	for i, user := range s.users {
		if user.GetID() == updatedUser.GetID() {
			s.users[i] = *updatedUser
			return nil
		}
	}
	return errors.New("user not found")
}

// DeleteUser removes a user by ID
func (s *UserService) DeleteUser(id int) error {
	for i, user := range s.users {
		if user.GetID() == id {
			s.users = append(s.users[:i], s.users[i+1:]...)
			return nil
		}
	}
	return errors.New("User not found")
}

func (s *UserService) GetAllUsers() []model.User {
	return s.users
}
