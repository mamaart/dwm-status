package server

import (
	"log"
	"net/http"
)

func (s *Server) stream(w http.ResponseWriter, r *http.Request) {
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
