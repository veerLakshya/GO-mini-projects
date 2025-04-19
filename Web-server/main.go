package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// to check the incoming requests for debugging
func logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("→ %s %s\n", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// Serve the HTML form
		http.ServeFile(w, r, "./static/form.html")
		return

	case http.MethodPost:
		// Parse the POSTed form data
		if err := r.ParseForm(); err != nil {
			http.Error(w, "ParseForm() error: "+err.Error(), http.StatusBadRequest)
			return
		}
		// Debug log to confirm we got the values
		log.Printf("DEBUG: form values = %#v\n", r.Form)

		name := r.FormValue("name")
		email := r.FormValue("email")

		// Show confirmation
		fmt.Fprintf(w,
			"POST request successful\n\nHello %s!\nEmail: %s\n",
			name, email,
		)
		return

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed.", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprintf(w, "Hello World!")
}

func main() {
	// Show working directory so we can confirm relative paths
	if wd, err := os.Getwd(); err == nil {
		log.Println("Working dir:", wd)
	}

	fileServer := http.FileServer(http.Dir("./static"))

	// Handle root route
	http.Handle("/", fileServer)

	// Handle form route
	http.HandleFunc("/form", formHandler)

	// Handle hello route
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Listening on port 8080\n")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

/*
Server
routes -
	1- 	/ 		=> index.html
	2- 	/hello 	=> hello func
	3- 	/form 	=> form func -> form.html
*/
