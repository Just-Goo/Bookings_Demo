package handlers

import (
	"fmt"
	"net/http"

	"github.com/Just-Goo/Bookings_Demo/pkg/config"
	"github.com/Just-Goo/Bookings_Demo/pkg/models"
	"github.com/Just-Goo/Bookings_Demo/pkg/render"
)

// Repo is the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (rp *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	rp.App.Session.Put(r.Context(), "remote_ip", remoteIP) // Storing the remote IP in the session

	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
}

func (rp *Repository) About(w http.ResponseWriter, r *http.Request) {
	// Perform some logic
	stringMap := make(map[string]string)
	stringMap["testing"] = "Hello again"

	remoteIp := rp.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIp
	fmt.Println(remoteIp)

	render.RenderTemplate(w, "about.page.html", &models.TemplateData{})
}
