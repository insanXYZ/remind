package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	APP_PORT = ":3124"
)

type Server struct {
	lastId         int
	cacheRemindMap map[Id]*RemindData
}

func NewServer(loadedRemindData map[Id]*RemindData, lastId int) *Server {
	return &Server{
		lastId:         lastId,
		cacheRemindMap: loadedRemindData,
	}
}

func (s *Server) incLastId() int {
	s.lastId += 1
	return s.lastId
}

func (s *Server) initRoute(mux *http.ServeMux) {
	mux.HandleFunc("GET /", s.ListController)
	mux.HandleFunc("DELETE /{id}", s.DeleteController)
	mux.HandleFunc("PATCH /", s.CheckController)
	mux.HandleFunc("POST /", s.SetController)
}

func (s *Server) Run() error {

	mux := http.NewServeMux()
	s.initRoute(mux)
	WriteLog(fmt.Sprintf(SuccRunServer+" , port %s", APP_PORT))
	return http.ListenAndServe(APP_PORT, mux)

}

// CONTROLLER

func (s *Server) ListController(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(s.cacheRemindMap)
	if err != nil {
		fmt.Fprint(w, err.Error())
	}
}

func (s *Server) CheckController(w http.ResponseWriter, r *http.Request) {

}

func (s *Server) DeleteController(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		fmt.Fprint(w, "id kudu aya")
		return
	}

	err := s.delete(id)
	if err != nil {
		fmt.Fprint(w, "error delete")
		return
	}

	fmt.Fprint(w, "success delete")

}

func (s *Server) SetController(w http.ResponseWriter, r *http.Request) {
	req := new(CreateRequest)
	now := time.Now()

	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}

	req.Name, req.Type, req.Time, req.Date = TrimSpace(req.Name), TrimSpace(req.Type), TrimSpace(req.Time), TrimSpace(req.Date)

	if req.Name == "" {
		fmt.Fprint(w, "name kudu aya")
		return
	}

	if strings.Contains(req.Name, ":") {
		split := strings.Split(req.Name, ":")
		if len(split) > 1 {
			req.Title = split[0]
			req.Name = split[1]
		} else {
			req.Name = split[0]
		}
	}

	if req.Type == "" {
		req.Type = string(ALARM)
	}

	if req.Type == string(ALARM) || req.Type == string(TASK) {
		if req.Date == "" {
			req.Date = now.Format(time.DateOnly)
		}

		remindData := &RemindData{
			Id:    s.incLastId(),
			Title: req.Title,
			Name:  req.Name,
			Time:  req.Time,
			Date:  req.Date,
			Type:  RemindType(req.Type),
		}

		err = s.Set(remindData)
		if err != nil {
			fmt.Fprint(w, err.Error())
		}
		fmt.Fprint(w, SuccSetRemind)

	} else {
		fmt.Fprint(w, ErrSetRemind)
		fmt.Fprint(w, ErrWrongType)
		return
	}

}

func (s *Server) Set(data *RemindData) error {
	s.cacheRemindMap[data.Id] = data

	return WriteFile(APP_DATA_FILENAME, s.cacheRemindMap, false)
}

func (s *Server) delete(id string) error {
	i, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	delete(s.cacheRemindMap, i)

	return WriteFile(APP_DATA_FILENAME, s.cacheRemindMap, false)
}
