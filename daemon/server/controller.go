package server

import (
	"net/http"
	"remind-daemon/model"
	"strings"
)

func (s *Server) listController(w http.ResponseWriter, r *http.Request) {
	s.giveResponse(w, 200, s.cacheRemindMap, model.SuccGetAllRemind)
}

func (s *Server) checkController(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		s.giveResponse(w, 400, nil, model.ErrIdRequired.Error())
		return
	}

	rflag := r.URL.Query().Has("r")

	if err := s.check(id, rflag); err != nil {
		s.giveResponse(w, 400, nil, err.Error())
		return
	}

	s.giveResponse(w, 200, nil, model.SuccCheckRemind)

}

func (s *Server) deleteController(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		s.giveResponse(w, 400, nil, model.ErrIdRequired.Error())
		return
	}

	err := s.delete(id)
	if err != nil {
		s.giveResponse(w, 400, nil, err.Error())
		return
	}

	s.giveResponse(w, 200, nil, model.SuccDeleteRemind)
}

func (s *Server) setController(w http.ResponseWriter, r *http.Request) {

	req, err := s.validateSetRequest(r)
	if err != nil {
		s.giveResponse(w, 400, nil, err.Error())
		return
	}

	remindData := &model.RemindData{
		Id:   s.incLastId(),
		Name: req.Name,
		Time: req.Time,
		Date: req.Date,
	}

	if strings.Contains(req.Name, ":") {
		splits := strings.Split(req.Name, ":")
		if len(splits) > 1 {
			remindData.Title, remindData.Name = splits[0], splits[1]
		}
	}

	if err := s.set(remindData); err != nil {
		s.giveResponse(w, 400, nil, err.Error())
		return
	}

	s.giveResponse(w, 200, nil, model.SuccSetRemind)
}
