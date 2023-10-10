package signup

import (
	"fmt"
	"net/http"
	"text/template"
	"time"

	"github.com/lumix103/cse-4050-project/internal/db/schema"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func Patient(w http.ResponseWriter, r *http.Request, client *mongo.Client) {
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
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Bad Request: Unable to parse form data", http.StatusBadRequest)
			return
		}
		if err := enforcePatientSignUpValues(r); err != nil {
			http.Error(w, "Bad Requst: missing required data", http.StatusBadRequest)
			return
		}
		hash, err := bcrypt.GenerateFromPassword([]byte(r.PostFormValue("password")), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}
		date, err := time.Parse("2006-01-02", r.PostFormValue("dob"))
		if err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}
		patient := schema.NewPatientUser(r.PostFormValue("email"), r.PostFormValue("first-name"),
			r.PostFormValue("last-name"), date, r.PostFormValue("gender"),
			r.PostFormValue("username"), string(hash))
		exists, err := schema.CheckIfPatientExists(client, r.PostFormValue("username"), r.PostFormValue("email"))

		if err != nil {
			fmt.Println(err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		if exists {
			http.Error(w, "User already exists", http.StatusConflict)
			return
		}

		if err := schema.InsertPatientUser(client, patient); err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		} else {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		}

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func enforcePatientSignUpValues(r *http.Request) error {
	if r.PostFormValue("username") == "" {
		return fmt.Errorf("`Username` value was empty")
	}
	if r.PostFormValue("first-name") == "" {
		return fmt.Errorf("`First name` value was empty")
	}
	if r.PostFormValue("last-name") == "" {
		return fmt.Errorf("`Last name` value was empty")
	}
	if r.PostFormValue("dob") == "" {
		return fmt.Errorf("`Date of Birth` value was empty")
	}
	if r.PostFormValue("email") == "" {
		return fmt.Errorf("`Email` value was empty")
	}
	if r.PostFormValue("password") == "" {
		return fmt.Errorf("`Password` value was empty")
	}

	return nil
}
