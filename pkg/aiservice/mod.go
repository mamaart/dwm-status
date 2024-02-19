package aiservice

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
)

type AiService struct {
	llm *ollama.LLM
}

func New(llm *ollama.LLM) *AiService {
	return &AiService{llm: llm}
}

func (s *AiService) ListenAndServe() error {
	router := mux.NewRouter()
	router.HandleFunc("/ask", s.ask)

	return http.ListenAndServe(":4343", router)
}

func (s *AiService) ask(w http.ResponseWriter, r *http.Request) {
	q := r.FormValue("question")
	if q == "" {
		http.Error(w, "missing question", http.StatusNotAcceptable)
		return
	}

	conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:4545/stream", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if _, err := s.llm.Call(
		context.Background(),
		q,
		llms.WithStreamingFunc(func(_ context.Context, chunk []byte) error {
			if err := conn.WriteMessage(websocket.BinaryMessage, chunk); err != nil {
				return fmt.Errorf("write to stream failed: %s", err)
			}
			return nil
		}),
	); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
