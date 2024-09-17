package server

import (
	"fmt"
	"html/template"
	"net/http"

	api "goupie-tracker/api"
)

// Homepage handler that executes the template.html file.
func HomePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		ErrorHandler(w, r, http.StatusNotFound)
		return
	}

	data := api.ArtistData()

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, data)
	fmt.Println("Endpoint: Main page")
}

// Manages the artist page display when an artist image is clicked by matching
// the "ArtistName" value against names in the Data.Artist.Name field.
func ArtistPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/artistInfo" {
		ErrorHandler(w, r, http.StatusNotFound)
		return
	}

	value := r.FormValue("ArtistName")

	if value == "" {
		ErrorHandler(w, r, http.StatusBadRequest)
		return
	}

	a := api.CollectData()
	var b api.Data
	found := false

	for i, ele := range a {
		if value == ele.A.Name {
			b = a[i]
			found = true
			break
		}
	}

	if !found {
		ErrorHandler(w, r, http.StatusBadRequest)
		return
	}

	tmpl, err := template.ParseFiles("templates/artist.html")
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, b)
	fmt.Println("Endpoint: " + value + "'s page")
}
