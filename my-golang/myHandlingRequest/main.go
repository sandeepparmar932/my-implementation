package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	getCall()
	postCall()
}

func postCall() {
	str := "http://localhost:8080/post"
	responseBody := strings.NewReader(`
	{
		"Name" : "Sandeep",
		"empId" : "18023"
	}
	`)
	res, _ := http.Post(str, "application/json", responseBody)
	defer res.Body.Close()
	content, _ := io.ReadAll(res.Body)
	fmt.Println("Post Content is ", string(content))
}

func getCall() {
	str := "http://localhost:8080/get"

	resp, _ := http.Get(str)
	defer resp.Body.Close()
	r, _ := io.ReadAll(resp.Body)
	fmt.Println("Response is ", string(r))
}
