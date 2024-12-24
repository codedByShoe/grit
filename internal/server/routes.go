package server

import (
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"net/http"
)

func (s *Server) RegisterRoutes(assets embed.FS) http.Handler {
	mux := http.NewServeMux()

	// Register routes
	mux.HandleFunc("GET /", s.HelloWorldHandler)
	mux.HandleFunc("GET /about", s.AboutHandler)
	mux.HandleFunc("GET /login", s.LoginHandler)
	mux.HandleFunc("GET /register", s.RegisterHandler)

	mux.HandleFunc("GET /health", s.healthHandler)

	staticFS, _ := fs.Sub(assets, "static/dist")
	mux.Handle("GET /static/dist/", http.StripPrefix("/static/dist/", http.FileServer(http.FS(staticFS))))

	return s.corsMiddleware(s.InertiaSharedPropsMiddleware(mux))
}

func (s *Server) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	if err := s.inertiaManager.Render(w, r, "Index", nil); err != nil {
		fmt.Fprintf(w, "Error Parsing template: %v", err)
	}
}

func (s *Server) AboutHandler(w http.ResponseWriter, r *http.Request) {
	if err := s.inertiaManager.Render(w, r, "About", nil); err != nil {
		fmt.Fprintf(w, "Error Parsing template: %v", err)
	}

}

func (s *Server) LoginHandler(w http.ResponseWriter, r *http.Request) {
	if err := s.inertiaManager.Render(w, r, "auth/Login", nil); err != nil {
		fmt.Fprintf(w, "Error Parsing template: %v", err)
	}
}
func (s *Server) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if err := s.inertiaManager.Render(w, r, "auth/Register", nil); err != nil {
		fmt.Fprintf(w, "Error Parsing template: %v", err)
	}
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := json.Marshal(s.db.Health())
	if err != nil {
		http.Error(w, "Failed to marshal health check response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(resp); err != nil {
		log.Printf("Failed to write response: %v", err)
	}
}
