package api

import (
	"log"
	"net/http"

	"github.com/mamaart/statusbar/internal/ports"
	"github.com/mamaart/statusbar/internal/statusbar/server"
	"github.com/mamaart/statusbar/modules/textmodule"
	"github.com/mamaart/statusbar/modules/timemodule"
)

func Run(
	time *timemodule.TimeModule,
	text *textmodule.TextModule,
) error {
	s := server.NewServer()

	s.HandleFunc(
		"/time",
		func(w http.ResponseWriter, _ *http.Request, _ func() (ports.UniStreamer, error)) {
			if err := time.Toggle(); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		},
	)

	s.HandleFunc(
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

	s.HandleFunc(
		"/text",
		func(w http.ResponseWriter, _ *http.Request, _ func() (ports.UniStreamer, error)) {
			if err := text.Toggle(); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		},
	)

	log.Println("Listen and serving api")
	return s.ListenAndServe()
}
