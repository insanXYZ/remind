package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	APP_PORT = ":3124"
)

type Server struct {
	cacheRemindMap map[Id]RemindData
}

func NewServer(loadedRemindData map[Id]RemindData) *Server {
	return &Server{
		cacheRemindMap: loadedRemindData,
	}
}

func (s *Server) initRoute(mux *http.ServeMux) {
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		err := json.NewEncoder(w).Encode(s.cacheRemindMap)
		if err != nil {
			fmt.Fprint(w, ErrGetCacheRemindData)
		}
	})
	mux.HandleFunc("DELETE /", func(http.ResponseWriter, *http.Request) {})
	mux.HandleFunc("POST /", func(w http.ResponseWriter, r *http.Request) {})
}

func (s *Server) Run() error {

	mux := http.NewServeMux()
	s.initRoute(mux)
	WriteLog(fmt.Sprintf(SuccRunServer+" , port %s", APP_PORT))
	return http.ListenAndServe(APP_PORT, mux)

}
