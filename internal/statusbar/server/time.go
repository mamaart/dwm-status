package server

import (
	"net/http"
	"time"
)

func (s *Server) time(w http.ResponseWriter, r *http.Request) {
	select {
	case s.clockstate <- struct{}{}:
		w.Write([]byte("signal sent\n"))
		return
	case <-time.After(time.Second * 5):
		http.Error(w, "timeout", http.StatusRequestTimeout)
		return
	}
}
