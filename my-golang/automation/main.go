package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/*
	func main() {
		pw, err := playwright.Run()

		if err != nil {
			log.Fatalf("Could not start Playwright: %v", err)
		}
		browser, err := pw.Chromium.Launch(
			playwright.BrowserTypeLaunchOptions{
				Headless: playwright.Bool(false),
			},
		)

		if err != nil {
			log.Fatalf("Could not launch the browser: %v", err)
		}

		page, err := browser.NewPage()
		if err != nil {
			log.Fatalf("Could not open a new page: %v", err)
		}

		if _, err = page.Goto("https://www.youtube.com/"); err != nil {
			log.Fatalf("Could not visit the desired page: %v", err)
		}

		html, err := page.Content()
		if err != nil {
			log.Fatalf("Could retrieve the HTML of the page: %v", err)
		}
		//fmt.Println(html)

		searchInput := page.Locator("input#search")
		err = searchInput.WaitFor()
		if err != nil {
			log.Fatalf("Could not find search input: %v", err)
		}

		if err := searchInput.Focus(); err != nil {
			log.Fatalf("Could not focus on search input: %v", err)
		}

		// Clear any existing text in the input field
		if err := searchInput.Fill(""); err != nil {
			log.Fatalf("Could not clear the search input: %v", err)
		}

		// Type the search query
		if err := searchInput.Fill("playwright automation"); err != nil {
			log.Fatalf("Could not type in search input: %v", err)
		}

		// Press Enter to trigger the search
		if err := searchInput.Press("Enter"); err != nil {
			log.Fatalf("Could not press Enter key: %v", err)
		}

		// Wait for the results to load
		time.Sleep(5 * time.Second)

		// Get the page content after the search
		html, err = page.Content()
		if err != nil {
			log.Fatalf("Could not retrieve the HTML of the page: %v", err)
		}
		fmt.Println(html)

}
*/
type User struct {
	ID   int
	Name string
}

func getUser(w http.ResponseWriter, r *http.Request) {
	user := User{
		ID:   1,
		Name: "Sandeep",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)

}

func main() {
	http.HandleFunc("/user", getUser)
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
