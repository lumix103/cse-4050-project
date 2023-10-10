package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lumix103/cse-4050-project/internal/routes/login"
	"github.com/lumix103/cse-4050-project/internal/routes/signup"
	"go.mongodb.org/mongo-driver/mongo"
)

func InitalizeRoutes(r *mux.Router, client *mongo.Client) {
	r.HandleFunc("/", Home).Methods("GET")
	r.HandleFunc("/login", mongodbMiddleware(login.Patient, client)).Methods("GET", "POST")
	r.HandleFunc("/signup", mongodbMiddleware(signup.Patient, client)).Methods("GET", "POST")
}

func mongodbMiddleware(f func(http.ResponseWriter, *http.Request, *mongo.Client), c *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		f(w, r, c)
	}
}
