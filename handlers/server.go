package handlers

import "net/http"

type Server struct {}

func (s *Server) Start() error {
	// Mux
	mux := http.NewServeMux()

	// Serve static files
	fs := http.FileServer(http.Dir("web/static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	// Routes
	mux.HandleFunc("/{$}", Home)

	// Init server
	return http.ListenAndServe(":3000", mux)
}

func GetServer() *Server {
	return &Server{}
}