package dashboard

import (
	"net/http"
	"text/template"

	"github.com/lumix103/cse-4050-project/internal/auth"
	"github.com/lumix103/cse-4050-project/internal/db/schema"
	"go.mongodb.org/mongo-driver/mongo"
)

func DoctorDashboard(w http.ResponseWriter, r *http.Request, client *mongo.Client) {
	tmpl, err := template.ParseFiles("./web/templates/dashboard/doctor/dashboard.html")

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
	doctor, err := schema.FetchDoctorBy(client, "username", username)

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, doctor); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
