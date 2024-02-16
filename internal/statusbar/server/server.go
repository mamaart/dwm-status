package server

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

type Server struct {
	upgrader *websocket.Upgrader
	ch       chan<- byte
}

func NewServer(ch chan<- byte) *Server {
	return &Server{
		upgrader: &websocket.Upgrader{},
		ch:       ch,
	}
}

func (s *Server) ListenAndServe() error {
	router := mux.NewRouter()
	router.HandleFunc("/stream", s.ai)

	return http.ListenAndServe(":4545", router)
}

func (s *Server) ai(w http.ResponseWriter, r *http.Request) {
	conn, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for {
		_, data, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		for _, b := range data {
			s.ch <- b
		}
	}
}
