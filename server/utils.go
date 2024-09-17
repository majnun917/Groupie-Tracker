package server

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

type Text struct {
	ErrorNum int
	ErrorMes string
}

// Custom handler to prevent directory listing
func NoDirListing(h http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/") {
			ErrorHandler(w, r, http.StatusNotFound)
			return
		}
		h.ServeHTTP(w, r)
	})
}

// Constructor function for error text message
func TextPerStatus(status int, em string) Text {
	return Text{
		ErrorNum: status,
		ErrorMes: em,
	}
}

// Handles error messages.
func ErrorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)

	if status == http.StatusNotFound {
		tmpl, err := template.ParseFiles("templates/error.html")
		if err != nil {
			ErrorHandler(w, r, http.StatusInternalServerError)
			return
		}
		em := "HTTP status 404: Page Not Found"
		p := TextPerStatus(status, em)
		tmpl.Execute(w, p)
	}

	if status == http.StatusInternalServerError {
		tmpl, err := template.ParseFiles("templates/error.html")
		if err != nil {
			fmt.Fprint(w, "HTTP status 500: Internal Server Error -missing error.html file")
		}
		em := "HTTP status 500: Internal Server Error"
		p := TextPerStatus(status, em)
		tmpl.Execute(w, p)
	}

	if status == http.StatusBadRequest {
		tmpl, err := template.ParseFiles("templates/error.html")
		if err != nil {
			fmt.Fprint(w, "HTTP status 400: Bad Request!\n\nPlease select artist from the Home Page")
		}
		em := "HTTP status 400: Bad Request!\n\nPlease select artist from the Home Page"
		p := TextPerStatus(status, em)
		tmpl.Execute(w, p)
	}
}
