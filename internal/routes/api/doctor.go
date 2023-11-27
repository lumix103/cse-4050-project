package api

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/lumix103/cse-4050-project/internal/db/schema"
	"go.mongodb.org/mongo-driver/mongo"
)

func FetchDoctors(w http.ResponseWriter, r *http.Request, client *mongo.Client) {
	vars := mux.Vars(r)

	start, err := strconv.Atoi(vars["start"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	end, err := strconv.Atoi(vars["end"])

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	doctors, err := schema.FetchDoctors(client, start, end)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	for _, doc := range doctors {
		fmt.Println(doc.Email)
	}

}

func FetchDoctor(w http.ResponseWriter, r *http.Request, client *mongo.Client) {
	vars := mux.Vars(r)

	if id, ok := vars["doctor_id"]; ok {
		doc, err := schema.FetchDoctorBy(client, "_id", id)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		tmpl, err := template.ParseFiles("./web/templates/api/doctorInfo.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if err := tmpl.Execute(w, doc); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

}
