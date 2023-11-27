package routes

import (
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	if err := HomePage.Execute(w, isLoggedIn(r)); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func AboutUs(w http.ResponseWriter, r *http.Request) {
	if err := AboutUsPage.Execute(w, isLoggedIn(r)); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func ChoiceLogin(w http.ResponseWriter, r *http.Request) {

	if login := isLoggedIn(r); !login {
		if err := ChoicePage.Execute(w, login); err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	} else {
		http.Redirect(w, r, "/portal", http.StatusSeeOther)
	}

}
