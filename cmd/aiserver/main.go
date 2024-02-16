package main

import (
	"log"

	"github.com/mamaart/statusbar/internal/ai/server"
	"github.com/tmc/langchaingo/llms/ollama"
)

func main() {
	// TODO run server
	llm, err := ollama.New(ollama.WithModel("mistral"))
	if err != nil {
		log.Fatalf("failed to make llm: %s", err)
	}
	s := server.NewServer(llm)
	log.Fatal(s.ListenAndServe())
}
