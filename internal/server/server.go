package server

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/mamaart/statusbar/internal/database"
	"github.com/mamaart/statusbar/internal/models"
)

type Server struct {
	db       *database.DB
	server   http.Server
	taskList chan []models.Task
}

func New(db *database.DB) *Server {
	return &Server{
		db: db,
		server: http.Server{
			Addr: ":8080",
			Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.Method == http.MethodDelete {
					id, err := parseId(r.URL.Path)
					if err != nil {
						w.WriteHeader(http.StatusBadRequest)
						w.Write([]byte("invalid path"))
					}
					db.Delete(id)
				}
				if r.Method == http.MethodPost {
					var x models.Task
					if err := json.NewDecoder(r.Body).Decode(&x); err != nil {
						w.WriteHeader(http.StatusBadRequest)
						w.Write([]byte(err.Error()))
						return
					}
					db.Add(x)
				}
			}),
		},
	}
}

func parseId(path string) (int, error) {
	components := strings.Split(path, "/")
	if len(components) == 0 {
		return 0, errors.New("invalid path")
	}
	idStr := components[len(components)-1]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, errors.New("invalid id")
	}

	return id, nil
}

func (s *Server) Run() {
	s.server.ListenAndServe()
}
