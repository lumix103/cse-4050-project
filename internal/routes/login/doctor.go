package login

import (
	"fmt"
	"html/template"
	"net/http"
)

func Doctor(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./web/templates/login/doctor.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	switch r.Method {
	case http.MethodGet:
		if err := tmpl.Execute(w, nil); err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	case http.MethodPost:
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Bad Request: Unable to parse form data", http.StatusBadRequest)
			return
		}
		if err := enforceDoctorLoginValues(r); err != nil {
			http.Error(w, "Bad Requst: missing required data", http.StatusBadRequest)
			return
		}
	default:

	}
}

func enforceDoctorLoginValues(r *http.Request) error {
	if r.PostFormValue("username") == "" {
		return fmt.Errorf("`Username` value was empty")
	}
	if r.PostFormValue("password") == "" {
		return fmt.Errorf("`Password` value was empty")
	}

	return nil
}
