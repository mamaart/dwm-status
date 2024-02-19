package server

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/mamaart/statusbar/internal/ports"
)

type Server struct {
	upgrader *websocket.Upgrader
	router   *mux.Router
}

func NewServer() *Server {
	return &Server{
		upgrader: &websocket.Upgrader{},
		router:   mux.NewRouter(),
	}
}

func (s *Server) HandleFunc(
	path string,
	f func(w http.ResponseWriter, r *http.Request, upgrade func() (ports.UniStreamer, error)),
) {
	s.router.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		f(w, r, func() (ports.UniStreamer, error) {
			conn, err := s.upgrader.Upgrade(w, r, nil)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return nil, err
			}
			return NewStreamer(conn), nil
		})
	})
}

func (s *Server) ListenAndServe() error {
	return http.ListenAndServe(":4545", s.router)
}

type Streamer struct {
	conn *websocket.Conn
}

func NewStreamer(conn *websocket.Conn) *Streamer {
	return &Streamer{conn}
}

func (s *Streamer) Reader() <-chan []byte {
	out := make(chan []byte)
	go func() {
		for {
			_, data, err := s.conn.ReadMessage()
			if err != nil {
				log.Println(err)
				return
			}
			out <- data
		}
	}()
	return out
}
