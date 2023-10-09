package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/lumix103/cse-4050-project/internal/routes"
	"github.com/lumix103/cse-4050-project/internal/routes/login"
	"github.com/lumix103/cse-4050-project/internal/routes/signup"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("No environmental variable 'MONGODB_URI' is set")
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))

	if err != nil {
		panic(err)
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	r := mux.NewRouter()
	r.HandleFunc("/", routes.Home).Methods("GET")
	r.HandleFunc("/login", login.Patient).Methods("GET", "POST")
	r.HandleFunc("/signup", signup.Patient).Methods("GET", "POST")
	http.ListenAndServe(":8000", r)
}
