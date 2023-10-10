package login

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/lumix103/cse-4050-project/internal/db/schema"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func Patient(w http.ResponseWriter, r *http.Request, client *mongo.Client) {
	tmpl, err := template.ParseFiles("./web/templates/login/patient.html")
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
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Bad Request: Unable to parse form data", http.StatusBadRequest)
			return
		}
		if err := enforcePatientLoginValues(r); err != nil {
			http.Error(w, "Bad Requst: missing required data", http.StatusBadRequest)
			return
		}
		patient, err := schema.FetchPatientExistsBy(client, "username", r.PostFormValue("username"))

		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		if patient == nil {
			w.WriteHeader(http.StatusUnauthorized)
			if err := tmpl.Execute(w, "Invalid username or password"); err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
			return
		}

		if err := bcrypt.CompareHashAndPassword([]byte(patient.Password), []byte(r.PostFormValue("password"))); err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			if err := tmpl.Execute(w, "Invalid username or password"); err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
			return
		} else {
			//TODO jwt auth
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	}
}

func enforcePatientLoginValues(r *http.Request) error {
	if r.PostFormValue("username") == "" {
		return fmt.Errorf("`Username` value was empty")
	}
	if r.PostFormValue("password") == "" {
		return fmt.Errorf("`Password` value was empty")
	}

	return nil
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
