// main.go
//Michael Bowden (10623182) code for Task B, Technical Report Application

package main

import (
	"html/template" // for parsing and rendering html files
	"log"           // for tracking errors and if server is live or not
	"net/http"      // for building the http server and routing pages
)

// function to handle loading and rendering
func renderTemplate(w http.ResponseWriter, tmpl string) {
	t, err := template.ParseFiles("templates/" + tmpl + ".html") // parses HTML file from templates
	if err != nil {                                              // if there is errors
		http.Error(w, "Error loading page", http.StatusInternalServerError) // if parsing fails, return an error
		log.Println("Template parsing error:", err)                         // log the parse error
		return
	}
	err = t.Execute(w, nil) // execute the parsed template, write the output
	if err != nil {         // if there is errors
		http.Error(w, "Error rendering page", http.StatusInternalServerError) // if execution fails, return an error
		log.Println("Template execution error:", err)                         // log the execution error
	}
}

// function to handle routes, loads up the pages connected to the routes using the renderTemplate
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { // route for homepage
		renderTemplate(w, "index") // loads up the HTML from index template
	})
	http.HandleFunc("/gallery", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "gallery") // loads up the HTML from gallery template
	})
	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "about") // loads up the HTML from about template
	})
	http.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "contact") // loads up the HTML from contact template
	})

	log.Println("Starting server on :8080")  // log and print the server start up to terminal (good for testing)
	err := http.ListenAndServe(":8080", nil) // load up the webserver to :8080 port, check if its nil
	if err != nil {
		log.Fatal(err) // if nil, log the error as fatal
	}
}

// end of code, thank you :P
