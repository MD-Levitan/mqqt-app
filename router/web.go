package router

import (
	"net/http"

	"github.com/MD-Levitan/mqqt-app/config"
)

func loginWebHandler(w http.ResponseWriter, r *http.Request) {
	config.GetTmpl().ExecuteTemplate(w, "login.html", nil)
}

func viewWebHandler(w http.ResponseWriter, r *http.Request) {
	config.GetTmpl().ExecuteTemplate(w, "view.html", nil)
}

func lightWebHandler(w http.ResponseWriter, r *http.Request) {
	config.GetTmpl().ExecuteTemplate(w, "light.html", nil)
}

func temperatureWebHandler(w http.ResponseWriter, r *http.Request) {
	config.GetTmpl().ExecuteTemplate(w, "temperature.html", nil)
}

func pressureWebHandler(w http.ResponseWriter, r *http.Request) {
	config.GetTmpl().ExecuteTemplate(w, "pressure.html", nil)
}

func humidityWebHandler(w http.ResponseWriter, r *http.Request) {
	config.GetTmpl().ExecuteTemplate(w, "humidity.html", nil)
}

func adminWebHandler(w http.ResponseWriter, r *http.Request) {
	config.GetTmpl().ExecuteTemplate(w, "admin.html", nil)
}

func logoutWebHandler(w http.ResponseWriter, r *http.Request) {
	if err := logout(w, r); err != nil {
		http.Redirect(w, r, "/login", http.StatusUnauthorized)
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func redirectLoginHandler(w http.ResponseWriter) {
	w.Header().Set("Location", "/login")
	w.WriteHeader(http.StatusTemporaryRedirect)
}
