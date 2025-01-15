package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	APP_PORT = ":3124"
)

type Server struct {
	mux            *http.ServeMux
	lastId         int
	cacheRemindMap map[Id]*RemindData
	mu             sync.Mutex
}

func NewServer(loadedRemindData map[Id]*RemindData, lastId int) *Server {
	s := &Server{
		mux:            http.NewServeMux(),
		lastId:         lastId,
		cacheRemindMap: loadedRemindData,
	}

	return s

}

func (s *Server) saveData() {
	WriteFile(APP_DATA_FILENAME, s.cacheRemindMap, false)
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

	req, err := s.validateSetRequest(r)
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}

	remindData := &RemindData{
		Id:   s.incLastId(),
		Name: req.Name,
		Time: req.Time,
		Date: req.Date,
		Type: RemindType(req.Type),
	}

	if strings.Contains(req.Name, ":") {
		splits := strings.Split(req.Name, ":")
		if len(splits) > 1 {
			remindData.Title, remindData.Name = splits[0], splits[1]
		}
	}

	s.Set(remindData)
}

func (s *Server) Set(data *RemindData) {
	s.cacheRemindMap[data.Id] = data
}

func (s *Server) delete(id string) error {
	i, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	delete(s.cacheRemindMap, i)
	return nil
}

// VALIDATOR
func (s *Server) validateSetRequest(req *http.Request) (*CreateRequest, error) {
	r := new(CreateRequest)
	now := time.Now()

	err := json.NewDecoder(req.Body).Decode(r)
	if err != nil {
		return nil, err
	}

	r.Name, r.Type, r.Date, r.Time = TrimSpace(r.Name), TrimSpace(r.Type), TrimSpace(r.Date), TrimSpace(r.Time)+":00"

	if r.Name == "" {
		return nil, ErrValidateNameRequired
	}

	if r.Type == "" {
		r.Type = string(ALARM)
	} else if r.Type != string(ALARM) && r.Type != string(TASK) {
		return nil, ErrWrongType
	}

	if r.Date != "" {
		if r.Date != "every-day" {
			if _, err := time.Parse(time.DateOnly, r.Date); err != nil {
				return nil, ErrWrongDate
			}
		}
	} else {
		r.Date = now.Format(time.DateOnly)
	}

	return r, nil

}
