package main

import (
	"net/http"
)

const (
	APP_PORT = ":3124"
)

type Server struct {
	cacheRemindMap map[int]RemindData
}

func NewServer(loadedRemindData map[int]RemindData) *Server {
	return &Server{
		cacheRemindMap: loadedRemindData,
	}
}

func (s *Server) initRoute(mux *http.ServeMux) {
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {})
	mux.HandleFunc("DELETE /", func(http.ResponseWriter, *http.Request) {})
	mux.HandleFunc("POST /", func(w http.ResponseWriter, r *http.Request) {})
}

func (s *Server) Run() error {

	mux := http.NewServeMux()
	s.initRoute(mux)

	return http.ListenAndServe(APP_PORT, mux)

}
