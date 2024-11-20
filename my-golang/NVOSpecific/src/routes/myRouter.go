package main

import (
	"NVOSpecific/src/controller"
	"fmt"
	"net/http"
)

func main() {

	genricHandler := controller.NewGenericHandler[*controller.User]()

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/add", genricHandler.AddResource)
	http.HandleFunc("/update", genricHandler.UpdateResource)
	http.HandleFunc("/get", genricHandler.GetResource)
	http.HandleFunc("/delete", genricHandler.DeleteResource)
	http.HandleFunc("/getAll", genricHandler.GetAllResources)

	port := 8080
	fmt.Printf("Server running on Port %d\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)

}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the homepage!")
}

// func objectHandler(w http.ResponseWriter, r *http.Request) {

// 	defer r.Body.Close()
// 	// Manually parse the path to extract the ID
// 	path := r.URL.Path
// 	segments := strings.Split(path, "/")
// 	id := segments[len(segments)-1] // Get the last segment as the ID

// 	var bodyData map[string]interface{}
// 	body, _ := io.ReadAll(r.Body)
// 	err := json.Unmarshal(body, &bodyData)
// 	fmt.Println("Data is - ", body)
// 	if err != nil {
// 		http.Error(w, "Error parsing JSON body", http.StatusBadRequest)
// 		return
// 	}
// 	bodyData["id"] = id
// 	w.Header().Set("Content-Type", "application/json")
// 	fmt.Println("Return body is - ", bodyData)
// 	json.NewEncoder(w).Encode(bodyData)
// }
