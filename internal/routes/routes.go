package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lumix103/cse-4050-project/internal/auth"
	"github.com/lumix103/cse-4050-project/internal/routes/api"
	"github.com/lumix103/cse-4050-project/internal/routes/dashboard"
	"github.com/lumix103/cse-4050-project/internal/routes/login"
	"github.com/lumix103/cse-4050-project/internal/routes/signup"
	"go.mongodb.org/mongo-driver/mongo"
)

func InitalizeRoutes(r *mux.Router, client *mongo.Client) {

	//    Page Routes
	// <================>
	// All users page
	r.HandleFunc("/", Home).Methods("GET")
	r.HandleFunc("/home", Home).Methods("GET")
	r.HandleFunc("/about-us", AboutUs).Methods("GET")
	r.HandleFunc("/choice", ChoiceLogin)
	r.HandleFunc("/logout", Logout).Methods("GET", "POST")
	// All patients pages
	r.HandleFunc("/login", mongodbMiddleware(login.Patient, client)).Methods("GET", "POST")
	r.HandleFunc("/signup", mongodbMiddleware(signup.Patient, client)).Methods("GET", "POST")
	// All doctors pages
	r.HandleFunc("/doctor-login", mongodbMiddleware(login.Doctor, client)).Methods("GET", "POST")
	r.HandleFunc("/doctor-signup", mongodbMiddleware(signup.Doctor, client)).Methods("GET", "POST")
	// All dashboard pages
	r.HandleFunc("/portal", authMongodbMiddleware(portalReroute, client))
	r.HandleFunc("/patient-portal", authMongodbMiddleware(dashboard.Patient, client))
	r.HandleFunc("/patient-reports", authMongodbMiddleware(dashboard.PatientReports, client))
	r.HandleFunc("/patient-schedule", authMongodbMiddleware(dashboard.PatientSchedule, client))

	r.HandleFunc("/doctor-portal", authMongodbMiddleware(dashboard.DoctorDashboard, client))
	r.HandleFunc("/doctor-reports", authMongodbMiddleware(dashboard.DoctorReports, client))
	r.HandleFunc("/doctor-schedule", authMongodbMiddleware(dashboard.DoctorAppointments, client))
	//     API Routes
	// <================>
	apiRoutes := r.PathPrefix("/api").Subrouter()
	apiRoutes.HandleFunc("/doctors/{start:[0-9]+}/{end:[0-9]+}", authMongodbMiddleware(api.FetchDoctors, client)).Methods("GET")
	apiRoutes.HandleFunc("/doctor/{doctor_id}", authMongodbMiddleware(api.FetchDoctor, client)).Methods("GET")
	apiRoutes.HandleFunc("/patient/{patient_id}", authMongodbMiddleware(api.FetchPatient, client)).Methods("GET")
	apiRoutes.HandleFunc("/appointments/{id}", authMongodbMiddleware(api.FetchAppointment, client)).Methods("GET")
	apiRoutes.HandleFunc("/schedule", authMongodbMiddleware(api.Schedule, client)).Methods("POST")
	apiRoutes.HandleFunc("/unschedule/{appointment_id}", authMongodbMiddleware(api.Unschedule, client)).Methods("POST")
}

func isLoggedIn(r *http.Request) bool {
	cookie, err := r.Cookie("token")

	if err != nil {
		return false
	}
	_, err = auth.ParseToken(cookie.Value)

	return err == nil
}

func portalReroute(w http.ResponseWriter, r *http.Request, c *mongo.Client) {
	// auth middleware should of checked for errors already
	cookie, _ := r.Cookie("token")
	claims, _ := auth.ParseToken(cookie.Value)

	if claims["user"] == "p" {
		http.Redirect(w, r, "/patient-portal", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/doctor-portal", http.StatusSeeOther)
	}
}

func mongodbMiddleware(f func(http.ResponseWriter, *http.Request, *mongo.Client), c *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		f(w, r, c)
	}
}

func authMongodbMiddleware(f func(http.ResponseWriter, *http.Request, *mongo.Client), c *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")

		if err != nil {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		_, err = auth.ParseToken(cookie.Value)

		if err != nil {
			http.Redirect(w, r, "/", http.StatusUnauthorized)
			return
		}

		f(w, r, c)
	}
}
