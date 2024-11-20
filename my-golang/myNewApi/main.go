package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Course struct {
	CourseId    string  `json :"courseid"`
	CourseName  string  `json :"coursename"`
	CoursePrice int     `json:"price"`
	IsFree      bool    `json : "isFree"`
	Author      *Author `json :"author"`
}
type Author struct {
	FullName string `json :"fullname"`
	Website  string `json:"website"`
}

var courses []Course

func (c *Course) isEmpty() bool {
	if c.CourseName == "" || (c.CoursePrice == 0 && !c.IsFree) {
		return true
	}
	return false
}

func serverHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>First API</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//get id
	params := mux.Vars(r)
	for _, course := range courses {
		if params["id"] == course.CourseId {
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	message := fmt.Sprintf("No course found with id - %s ", params["id"])
	err := errors.New(message)
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
}

func createCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//get id
	if r.Body == nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Please send data"})
		return
	}

	var course Course
	json.NewDecoder(r.Body).Decode(&course)

	if course.isEmpty() {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Please send mandatory data"})
		return
	}

	//generating new courseId
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)

}

func updateCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//get id
	parmas := mux.Vars(r)

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send data")
	}

	var course Course
	json.NewDecoder(r.Body).Decode(&course)

	if course.isEmpty() || course.CourseId == "" {
		json.NewEncoder(w).Encode("Please send mandatory data")
	}

	for i, cou := range courses {
		if cou.CourseId == parmas["id"] {
			courses = append(courses[:i], courses[i+1:]...)
			course.CourseId = parmas["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("Data with id not found")
	return

}

func main() {
	r := mux.NewRouter()

	//seeding
	courses = append(courses, Course{CourseId: "1", CourseName: "Java", CoursePrice: 1500, Author: &Author{
		FullName: "Sandeep", Website: "google",
	}})
	courses = append(courses, Course{CourseId: "2", CourseName: "Go", CoursePrice: 1200, Author: &Author{
		FullName: "Kumar", Website: "google",
	}})

	r.HandleFunc("/", serverHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getCourse).Methods("GET")
	r.HandleFunc("/course", createCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateCourse).Methods("PUT")

	//listing to port
	log.Fatal(http.ListenAndServe(":4000", r))

}
