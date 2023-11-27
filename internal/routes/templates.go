package routes

import "html/template"

var HomePage *template.Template

var AboutUsPage *template.Template

var PatientLogin *template.Template

var DoctorLogin *template.Template

var DoctorSignup *template.Template

var PatientSignup *template.Template

func InitalizeTemplates() {
	HomePage = template.Must(template.ParseFiles("./web/templates/home.html"))
	AboutUsPage = template.Must(template.ParseFiles("./web/templates/aboutUs.html"))
	PatientLogin = template.Must(template.ParseFiles("./web/templates/login/patient.html"))
	DoctorLogin = template.Must(template.ParseFiles("./web/templates/login/doctor.html"))
	DoctorSignup = template.Must(template.ParseFiles("./web/templates/signup/doctor.html"))
	PatientSignup = template.Must(template.ParseFiles("./web/templates/login/patient.html"))
}
