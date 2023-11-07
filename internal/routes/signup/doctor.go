package signup

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/lumix103/cse-4050-project/internal/db/schema"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func Doctor(w http.ResponseWriter, r *http.Request, client *mongo.Client) {
	tmpl, err := template.ParseFiles("./web/templates/signup/doctor.html")
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
		if err := enforceDoctorSignUpValues(r); err != nil {
			http.Error(w, "Bas Request: missing required data", http.StatusBadRequest)
			return
		}

		hash, err := bcrypt.GenerateFromPassword([]byte(r.PostFormValue("password")), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}
		email := r.PostFormValue("email")
		username := r.PostFormValue("username")

		doctor := schema.NewDoctorUser(email, r.PostFormValue("first-name"), r.PostFormValue("last-name"),
			username, string(hash))

		exists, err := schema.CheckIfDoctorExists(client, email, username)

		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		} else if exists {
			http.Error(w, "User already exists", http.StatusConflict)
			return
		}

		if err := schema.InsertDoctorUser(client, doctor); err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		} else {
			http.Redirect(w, r, "/doctor-login", http.StatusSeeOther)
		}

	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func enforceDoctorSignUpValues(r *http.Request) error {
	if r.PostFormValue("username") == "" {
		return fmt.Errorf("`Username` value was empty")
	}
	if r.PostFormValue("first-name") == "" {
		return fmt.Errorf("`First name` value was empty")
	}
	if r.PostFormValue("last-name") == "" {
		return fmt.Errorf("`Last name` value was empty")
	}
	if r.PostFormValue("email") == "" {
		return fmt.Errorf("`Email` value was empty")
	}
	if r.PostFormValue("password") == "" {
		return fmt.Errorf("`Password` value was empty")
	}

	return nil
}
