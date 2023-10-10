package login

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/lumix103/cse-4050-project/internal/auth"
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
			http.Error(w, "Internal FETCH Server Error", http.StatusInternalServerError)
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
			jwtToken, exp, err := auth.GenerateToken(patient.Username, patient.FirstName+" "+patient.LastName, "p")
			if err != nil {
				fmt.Println(err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
			http.SetCookie(w, &http.Cookie{
				Name:     "token",
				Value:    jwtToken,
				Expires:  exp,
				HttpOnly: true,
			})
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
