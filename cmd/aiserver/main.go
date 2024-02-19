package main

import (
	"log"

	"github.com/mamaart/statusbar/pkg/aiservice"
	"github.com/tmc/langchaingo/llms/ollama"
)

func main() {
	llm, err := ollama.New(ollama.WithModel("mistral"))
	if err != nil {
		log.Fatalf("failed to make llm: %s", err)
	}
	log.Fatal(aiservice.New(llm).ListenAndServe())
}
