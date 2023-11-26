package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lumix103/cse-4050-project/internal/auth"
	"github.com/lumix103/cse-4050-project/internal/routes/dashboard"
	"github.com/lumix103/cse-4050-project/internal/routes/login"
	"github.com/lumix103/cse-4050-project/internal/routes/signup"
	"go.mongodb.org/mongo-driver/mongo"
)

func InitalizeRoutes(r *mux.Router, client *mongo.Client) {
	// All users page
	r.HandleFunc("/", Home).Methods("GET")
	r.HandleFunc("/home", Home).Methods("GET")
	r.HandleFunc("/logout", Logout).Methods("GET", "POST")
	// All patients pages
	r.HandleFunc("/login", mongodbMiddleware(login.Patient, client)).Methods("GET", "POST")
	r.HandleFunc("/signup", mongodbMiddleware(signup.Patient, client)).Methods("GET", "POST")
	// All doctors pages
	r.HandleFunc("/doctor-login", mongodbMiddleware(login.Doctor, client)).Methods("GET", "POST")
	r.HandleFunc("/doctor-signup", mongodbMiddleware(signup.Doctor, client)).Methods("GET", "POST")
	// All dashboard pages
	r.HandleFunc("/patient-portal", authMongodbMiddleware(dashboard.Patient, client))
}

func mongodbMiddleware(f func(http.ResponseWriter, *http.Request, *mongo.Client), c *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		f(w, r, c)
	}
}

func authMongodbMiddleware(f func(http.ResponseWriter, *http.Request, *mongo.Client), c *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")

		if err != nil {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		_, err = auth.ParseToken(cookie.Value)

		if err != nil {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		f(w, r, c)
	}
}
