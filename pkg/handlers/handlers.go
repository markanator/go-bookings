package handlers

import (
	"github.com/markanator/go-bookings/pkg/models"
	"net/http"

	"github.com/markanator/go-bookings/pkg/config"
	"github.com/markanator/go-bookings/pkg/render"
)

type Repository struct {
	App *config.AppConfig
}

// Repo is the repository used by the handlers
var Repo *Repository

// creates a new Repository
func NewRepository(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home page handler
func (m *Repository) HomeHandler(w http.ResponseWriter, r *http.Request) {
	remoteIp := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIp)
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About page handler
func (m *Repository) AboutHandler(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."

	remoteIp := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIp

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// ShowReservation renders the make a reservation page
func (m *Repository) ContactHandler(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "contact.page.tmpl", &models.TemplateData{})
}

// ShowReservation renders the make a reservation page
func (m *Repository) ShowSearchAvailability(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "search-availability.page.tmpl", &models.TemplateData{})
}

// ShowStandardRoom renders the make a reservation page
func (m *Repository) ShowStandardRoom(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "standard-room.page.tmpl", &models.TemplateData{})
}

// ShowStandardRoom renders the make a reservation page
func (m *Repository) ShowMasterRoom(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "master-room.page.tmpl", &models.TemplateData{})
}

// SearchReservations action handler
func (m *Repository) ShowMakeReservation(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "search-availability.page.tmpl", &models.TemplateData{})
}
