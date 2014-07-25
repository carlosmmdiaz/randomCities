package main

import (
	"html/template"
	"net/http"
	"randomcities"
	"time"
)

// Available templates.
var templates = template.Must(template.ParseFiles(
	"assets/cities.html",
))

// data needed to render the template
type Page struct {
	Title  string
	Cities randomcities.Cities
	Time   int
}

// Render templates.
func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Handler to show the random cities.
func viewHandler(w http.ResponseWriter, r *http.Request) {

	timeStarted := time.Now().Nanosecond()

	// generate a random city list
	var cities randomcities.Cities
	cities.New()
	cities.GetRandomCities("assets/US_Cities.txt")

	// get the total execution time
	timeEnded := time.Now().Nanosecond()
	executionTime := timeEnded - timeStarted

	// set the data needed to render the template, then render it
	page := Page{
		Title:  "RandomCities",
		Cities: cities,
		Time:   executionTime / 1e6, // divide to get time in milliseconds
	}
	renderTemplate(w, "cities", &page)
}

func main() {
	http.HandleFunc("/", viewHandler)
	http.ListenAndServe(":8080", nil)
}
