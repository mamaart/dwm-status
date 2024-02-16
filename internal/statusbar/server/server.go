package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

type Server struct {
	upgrader   *websocket.Upgrader
	ch         chan<- byte
	clockstate chan<- struct{}
}

func NewServer(ch chan<- byte, clockstate chan<- struct{}) *Server {
	return &Server{
		upgrader:   &websocket.Upgrader{},
		ch:         ch,
		clockstate: clockstate,
	}
}

func (s *Server) ListenAndServe() error {
	router := mux.NewRouter()
	router.HandleFunc("/stream", s.stream)
	router.HandleFunc("/time", s.time)

	return http.ListenAndServe(":4545", router)
}
