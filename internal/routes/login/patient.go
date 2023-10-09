package login

import (
	"html/template"
	"net/http"
)

func Patient(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./web/templates/login/patient.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	if r.Method == http.MethodGet {
		if err := tmpl.Execute(w, nil); err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	} else if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Bad Request: Unable to parse form data", http.StatusBadRequest)
		}
	}
}

/*
========================================================================
NOTE: old code I was using for proof of concept that I might reuse later
========================================================================
		if username, ok_name := r.Form["username"]; ok_name {
			if password, ok_pass := r.Form["password"]; ok_pass {
				if username[0] != "user" || password[0] != "password" {
					w.WriteHeader(http.StatusUnauthorized)
					if err := tmpl.Execute(w, "Invalid username or password"); err != nil {
						http.Error(w, "Internal Server Error", http.StatusInternalServerError)
					}
					return
				}
				http.Redirect(w, r, "/", http.StatusSeeOther)
			}
		}
*/
