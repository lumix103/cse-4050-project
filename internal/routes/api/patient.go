package api

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lumix103/cse-4050-project/internal/db/schema"
	"go.mongodb.org/mongo-driver/mongo"
)

func FetchPatient(w http.ResponseWriter, r *http.Request, client *mongo.Client) {
	vars := mux.Vars(r)

	if patientId, ok := vars["patient_id"]; ok {
		patient, err := schema.FetchPatientExistsBy(client, "_id", patientId)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		tmpl, err := template.ParseFiles("./web/templates/api/patientInfo.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if err := tmpl.Execute(w, patient); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}
