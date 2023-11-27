package routes

import (
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	if err := HomePage.Execute(w, nil); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func AboutUs(w http.ResponseWriter, r *http.Request) {
	if err := AboutUsPage.Execute(w, nil); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
