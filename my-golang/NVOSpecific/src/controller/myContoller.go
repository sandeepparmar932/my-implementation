package controller

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type User struct {
	ID    int    `json :"id"`
	Name  string `json : "name"`
	Email string `json : "email"`
}

var users []User
var createId int

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Please send data"})
		return
	}
	defer r.Body.Close()
	user := &User{}
	json.NewDecoder(r.Body).Decode(&user)
	if user.Email == "" || user.Name == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"error": "Please provide mandatory data"})
		return
	}
	createId = createId + 1
	user.ID = createId
	users = append(users, *user)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{"message": "User created", "User": *user})
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Please send data"})
		return
	}
	defer r.Body.Close()
	user := &User{}
	json.NewDecoder(r.Body).Decode(&user)
	if user.ID == 0 || user.Name == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"error": "Please provide mandatory data"})
	}

	for index, value := range users {
		if value.ID == user.ID {
			users = append(users[:index], users[index+1:]...)
			users = append(users, *user)
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]interface{}{"message": "User updated", "User": *user})
			return
		}
	}
	w.WriteHeader(http.StatusNoContent)
	json.NewEncoder(w).Encode(map[string]interface{}{"message": "User with Id not found", "User": *user})
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	id := queryParams.Get("id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"message": "Please provide mandatory Info"})
	}
	intID, _ := strconv.Atoi(id)

	for _, user := range users {
		if user.ID == intID {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]interface{}{"User": user})
			return
		}
	}
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(map[string]interface{}{"message": "Data not found"})
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	id := queryParams.Get("id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"message": "Please provide mandatory Info"})
	}
	intID, _ := strconv.Atoi(id)

	for _, user := range users {
		if user.ID == intID {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]interface{}{"message": "User Deleted"})
			return
		}
	}
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(map[string]interface{}{"message": "Data not found"})
}
