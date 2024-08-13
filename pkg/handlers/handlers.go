package handlers

import (
	"bookings/pkg/config"
	"bookings/pkg/models"
	"bookings/pkg/render"
	"net/http"
)

// Repo is a global variable to hold the repository instance.
var Repo *Repository

// Repository holds the application configuration.
type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
}

func (m *Repository) AboutPage(w http.ResponseWriter, r *http.Request) {
	// Get the remote IP address from the session
	remoteIP := m.App.Session.Get(r.Context(), "remote_ip")

	// Create a map to pass to the template
	stringMap := map[string]string{
		"test": "hi test here!!",
	}

	// Create a map for additional data
	data := map[string]interface{}{
		"ip": remoteIP,
	}

	// Render the template with the data
	render.RenderTemplate(w, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
		Data:      data,
	})
}
