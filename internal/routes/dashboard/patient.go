package dashboard

import (
	"net/http"
	"text/template"

	"go.mongodb.org/mongo-driver/mongo"
)

func Patient(w http.ResponseWriter, r *http.Request, client *mongo.Client) {
	tmpl, err := template.ParseFiles("./web/templates/dashboard/patient/menu.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
