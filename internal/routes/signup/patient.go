package signup

import (
	"fmt"
	"net/http"
	"text/template"
)

func Patient(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./web/templates/signup/patient.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	switch r.Method {
	case http.MethodGet:
		if err := tmpl.Execute(w, nil); err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	case http.MethodPost:
		fmt.Println("Something?")
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Bad Request: Unable to parse form data", http.StatusBadRequest)
		}
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

/*
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
*/
