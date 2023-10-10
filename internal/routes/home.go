package routes

import (
	"html/template"
	"net/http"

	"github.com/lumix103/cse-4050-project/internal/auth"
)

func Home(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./web/templates/home.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	cookie, err := r.Cookie("token")

	if err != nil {
		if err == http.ErrNoCookie {
			if err := tmpl.Execute(w, nil); err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
		} else {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	}

	claims, err := auth.ParseToken(cookie.Value)
	if err != nil {
		if err := tmpl.Execute(w, nil); err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		} else {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	} else {
		if err := tmpl.Execute(w, claims["name"]); err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}

}
