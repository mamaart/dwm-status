package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/mamaart/statusbar/internal/ports"
	"github.com/mamaart/statusbar/modules/textmodule"
	"github.com/mamaart/statusbar/modules/timemodule"
)

type Api struct {
	upgrader *websocket.Upgrader
	router   *mux.Router
}

func New() *Api {
	return &Api{
		upgrader: &websocket.Upgrader{},
		router:   mux.NewRouter(),
	}
}

func (a *Api) HandleFunc(
	path string,
	f func(w http.ResponseWriter, r *http.Request, upgrade func() (ports.UniStreamer, error)),
) {
	a.router.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		f(w, r, func() (ports.UniStreamer, error) {
			conn, err := a.upgrader.Upgrade(w, r, nil)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return nil, err
			}
			return NewStreamer(conn), nil
		})
	})
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

func (a *Api) Run(
	time *timemodule.TimeModule,
	text *textmodule.TextModule,
) error {
	a.HandleFunc(
		"/time",
		func(w http.ResponseWriter, _ *http.Request, _ func() (ports.UniStreamer, error)) {
			if err := time.Toggle(); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		},
	)

	a.HandleFunc(
		"/stream",
		func(_ http.ResponseWriter, _ *http.Request, upgrade func() (ports.UniStreamer, error)) {
			r, err := upgrade()
			if err != nil {
				log.Println(err)
				return
			}

			for e := range r.Reader() {
				if _, err := text.Write(e); err != nil {
					log.Println(err)
					return
				}
			}
		},
	)

	a.HandleFunc(
		"/text",
		func(w http.ResponseWriter, _ *http.Request, _ func() (ports.UniStreamer, error)) {
			if err := text.Toggle(); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		},
	)

	log.Println("Listen and serving api")
	return http.ListenAndServe(":4545", a.router)
}
