package controller

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type Resource interface {
	GetID() int
	SetID(id int)
}

func (u *User) GetID() int   { return u.ID }
func (u *User) SetID(id int) { u.ID = id }

type GenericHandler[T Resource] struct {
	store  []T
	nextID int
}

func NewGenericHandler[T Resource]() *GenericHandler[T] {
	return &GenericHandler[T]{store: []T{}, nextID: 1}
}

func (h *GenericHandler[T]) AddResource(w http.ResponseWriter, r *http.Request) {
	var resource T
	err := json.NewDecoder(r.Body).Decode(&resource)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid data"})
		return
	}
	defer r.Body.Close()
	resource.SetID(h.nextID)
	h.nextID++
	h.store = append(h.store, resource)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{"message": "Resource created", "Resource": resource})
}

func (h *GenericHandler[T]) GetResource(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	idStr := queryParams.Get("id")
	if idStr == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"message": "ID is required"})
		return
	}
	id, _ := strconv.Atoi(idStr)

	for _, resource := range h.store {
		if resource.GetID() == id {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(resource)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]interface{}{"message": "Resource not found"})
}

func (h *GenericHandler[T]) UpdateResource(w http.ResponseWriter, r *http.Request) {
	var resource T
	err := json.NewDecoder(r.Body).Decode(&resource)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid data"})
		return
	}
	defer r.Body.Close()

	for i, storedResource := range h.store {
		if storedResource.GetID() == resource.GetID() {
			h.store[i] = resource
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]interface{}{"message": "Resource updated", "Resource": resource})
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]interface{}{"message": "Resource not found"})
}

func (h *GenericHandler[T]) DeleteResource(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	idStr := queryParams.Get("id")
	if idStr == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"message": "ID is required"})
		return
	}
	id, _ := strconv.Atoi(idStr)
	for i, resource := range h.store {
		if resource.GetID() == id {
			h.store = append(h.store[:i], h.store[i+1:]...)
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]interface{}{"message": "Resource deleted"})
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]interface{}{"message": "Resource not found"})
}

func (h *GenericHandler[T]) GetAllResources(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(h.store)
}
