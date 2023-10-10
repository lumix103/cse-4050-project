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

	// Retrieve JWT cookie
	cookie, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			// No JWT cookie found, handle the case accordingly
			//fmt.Println("No JWT cookie found")
			if err := tmpl.Execute(w, nil); err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
			return
		}
		// Other cookie retrieval error
		//fmt.Println("Error retrieving JWT cookie:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Parse JWT token
	claims, err := auth.ParseToken(cookie.Value)
	if err != nil {
		// Handle invalid or expired token
		//fmt.Println("Error parsing JWT token:", err)
		if err := tmpl.Execute(w, nil); err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		return
	}

	// JWT token is valid, proceed with handling the authenticated request
	if err := tmpl.Execute(w, claims["name"]); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
