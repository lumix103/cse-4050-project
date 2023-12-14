package dashboard

import (
	"net/http"
	"text/template"

	"github.com/lumix103/cse-4050-project/internal/auth"
	"github.com/lumix103/cse-4050-project/internal/db/schema"
	"go.mongodb.org/mongo-driver/mongo"
)

func Patient(w http.ResponseWriter, r *http.Request, client *mongo.Client) {
	tmpl, err := template.ParseFiles("./web/templates/dashboard/patient/dashboard.html")

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	tok, _ := r.Cookie("token")
	claims, _ := auth.ParseToken(tok.Value)
	username, ok := claims["username"].(string)

	if !ok {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	patient, err := schema.FetchPatientExistsBy(client, "username", username)

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	appointments, err := schema.FetchFutureAppointments(client, patient, nil)

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	pa := schema.PatientAppointments{Patient: patient, Appointments: appointments}

	if err := tmpl.Execute(w, pa); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func PatientReports(w http.ResponseWriter, r *http.Request, client *mongo.Client) {
	tmpl, err := template.ParseFiles("./web/templates/dashboard/patient/PatientReportsactual.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func PatientSchedule(w http.ResponseWriter, r *http.Request, client *mongo.Client) {
	tmpl, err := template.ParseFiles("./web/templates/dashboard/patient/patientAppointments.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
