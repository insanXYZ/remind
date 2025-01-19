package server

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os/exec"
	"remind-daemon/model"
	"remind-daemon/util"
	"strconv"
	"time"
)

const (
	AppPort = ":5555"
)

type Server struct {
	mux            *http.ServeMux
	lastId         int
	cacheRemindMap model.RemindDatas
}

func NewServer() *Server {
	s := &Server{
		mux:            http.NewServeMux(),
		cacheRemindMap: make(model.RemindDatas),
	}

	go s.tickRemind()

	return s
}

func (s *Server) incLastId() int {
	s.lastId += 1
	return s.lastId
}

func (s *Server) loadData() error {
	changes := false
	missed := 0
	now := time.Now()

	file, err := util.ReadFile(model.APP_DATA_FILENAME)
	if err != nil {
		return err
	}

	err = json.NewDecoder(bytes.NewReader(file)).Decode(&s.cacheRemindMap)
	if err != nil {
		if errors.Is(err, io.EOF) {
			return nil
		}
		return err
	}

	for _, v := range s.cacheRemindMap {
		if v.Id > s.lastId {
			s.lastId = v.Id
		}

		if v.CheckedAt != "" && (v.Date == model.EVERY_DAY_DATE || v.Date == time.Now().Format(time.DateOnly)) {
			t, err := time.Parse(time.DateOnly, v.CheckedAt)
			if err != nil {
				return err
			}

			if util.GetStartOfDay().Unix() > t.Unix() {
				changes = true
				s.cacheRemindMap[v.Id].CheckedAt = ""
			}
		}

		if v.CheckedAt == "" && (v.Date == model.EVERY_DAY_DATE || v.Date == time.Now().Format(time.DateOnly)) {

			if v.Time == "" {
				missed++
			} else {

				p, err := time.Parse(time.TimeOnly, v.Time)

				if err != nil {
					return err
				}

				d := time.Date(now.Year(), now.Month(), now.Day(), p.Hour(), p.Minute(), p.Second(), 0, now.Location())
				if now.Unix() > d.Unix() {
					missed++
				}
			}
		}
	}

	go func() {
		if missed != 0 {
			s.notify("", fmt.Sprintf("you have %v missed remind", missed))
		}
	}()

	if changes {

		m, err := util.StructToJsonString(s.cacheRemindMap)
		if err != nil {
			return err
		}

		err = util.WriteFile(model.APP_DATA_FILENAME, m, false)
		if err != nil {
			return err
		}
	}

	return err

}

func (s *Server) initRoute() {
	s.mux.HandleFunc("GET /", s.listController)
	s.mux.HandleFunc("DELETE /{id}", s.deleteController)
	s.mux.HandleFunc("PATCH /{id}", s.checkController)
	s.mux.HandleFunc("POST /", s.setController)
}

func (s *Server) Run() error {
	s.initRoute()

	err := s.loadData()
	if err != nil {
		return err
	}

	util.WriteLog(fmt.Sprintf(model.SuccRunServer+" , port %s", AppPort))
	return http.ListenAndServe(AppPort, s.mux)
}

func (s *Server) set(data *model.RemindData) error {
	s.cacheRemindMap[data.Id] = data

	return s.saveData()
}

func (s *Server) delete(id string) error {
	i, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	_, ok := s.cacheRemindMap[i]

	if !ok {
		return fmt.Errorf("id with %v doesnt exist", i)
	}

	delete(s.cacheRemindMap, i)
	return s.saveData()

}

func (s *Server) check(id string, rflag bool) error {
	i, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	v, ok := s.cacheRemindMap[i]

	if !ok {
		return fmt.Errorf("id with %v doesnt exist", i)
	}

	if rflag {
		v.CheckedAt = ""
	} else {
		v.CheckedAt = time.Now().Format(time.DateOnly)
	}

	return s.saveData()

}

func (s *Server) saveData() error {
	m, err := util.StructToJsonString(s.cacheRemindMap)
	if err != nil {
		return err
	}

	err = util.WriteFile(model.APP_DATA_FILENAME, m, false)
	if err != nil {
		util.WriteLog(err.Error())
		return err
	}
	return nil
}

func (s *Server) giveResponse(w http.ResponseWriter, statusCode int, data any, message string) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	m, err := util.StructToJsonString(model.Response{
		Data:    data,
		Message: message,
	})

	if err != nil {
		return err
	}

	_, err = w.Write([]byte(m))
	return err
}

func (s *Server) tickRemind() {
	tick := time.NewTicker(1 * time.Second)

	for t := range tick.C {

		go func() {

			date := t.Format(time.DateOnly)
			clock := t.Format(time.TimeOnly)

			for _, v := range s.cacheRemindMap {
				if v.Date == model.EVERY_DAY_DATE || v.Date == date {
					if v.Time == clock {
						err := s.notify(v.Title, v.Name)
						if err != nil {
							fmt.Println(err.Error())
						}
					}
				}
			}
		}()

	}
}

func (s *Server) notify(title, name string) error {
	if title == "" {
		title = model.APP_NAME
	}

	dunst := fmt.Sprintf("dunstify \"%s\" \"%s\"", title, name)

	cmd := exec.Command("bash", "-c", dunst)
	return cmd.Run()

}

// VALIDATOR
func (s *Server) validateSetRequest(req *http.Request) (*model.CreateRequest, error) {
	r := new(model.CreateRequest)
	now := time.Now()

	err := json.NewDecoder(req.Body).Decode(r)
	if err != nil {
		return nil, err
	}

	r.Name, r.Date = util.TrimSpace(r.Name), util.TrimSpace(r.Date)

	if r.Name == "" {
		return nil, model.ErrValidateNameRequired
	}

	if r.Time != "" {

		r.Time = util.TrimSpace(r.Time) + ":00"

		if _, err := time.Parse(time.TimeOnly, r.Time); err != nil {
			return r, model.ErrWrongTime
		}

	}

	if r.Date != "" {
		if _, err := time.Parse(time.DateOnly, r.Date); err != nil {
			if r.Date != model.EVERY_DAY_DATE {
				return r, model.ErrWrongDate
			}
		}
	} else {
		r.Date = now.Format(time.DateOnly)
	}

	return r, nil

}
