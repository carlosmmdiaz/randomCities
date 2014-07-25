package main

import (
	"net/http"
    "html/template"
    "RandomCities"
    "time"
)

// Available templates.
var templates = template.Must(template.ParseFiles(
                "view/cities.html",
))

// Struct to save a webpage.
type Page struct{
	Title string
    Cities RandomCities.Cities
    Time int
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
    
    t1 := time.Now().Nanosecond()
    // File name:
    inputFileName := "src/txt/US_Cities.txt"

    // Create the class cities:
    var cities RandomCities.Cities
    cities.New()

    // Gets randomCities from the file:
    cities.GetRandomCities(inputFileName)

    t2 := time.Now().Nanosecond()

    t := t2 - t1

    times := t / 1e6

    //Creates page:
    page := Page{
        Title: "RandomCities",
        Cities: cities,
        Time: times,
    }
    // Show Page:
    renderTemplate(w, "cities", &page)
}

func main() {
    http.HandleFunc("/", viewHandler)
    http.ListenAndServe(":8080", nil)
}