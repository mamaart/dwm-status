package ai

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"

	"github.com/quic-go/quic-go"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
)

func Ask(q string) error {
	pool := x509.NewCertPool()
	pool.AddCert(read("/etc/statusbar/olla"))

	client, err := quic.DialAddr(context.Background(), "localhost:4343", &tls.Config{
		RootCAs:    pool,
		NextProtos: []string{"dwm-status"},
	}, nil)
	if err != nil {
		return fmt.Errorf("failed to dial client: %s", err)
	}

	stream, err := client.OpenStreamSync(context.Background())
	if err != nil {
		return fmt.Errorf("failed to open stream with client: %s", err)
	}
	defer stream.Close()

	llm, err := ollama.New(ollama.WithModel("mistral"))
	if err != nil {
		return fmt.Errorf("failed to make llm: %s", err)
	}

	// b := make([]byte, 1)
	if _, err := llm.Call(
		context.Background(),
		q,
		llms.WithStreamingFunc(func(_ context.Context, chunk []byte) error {
			if _, err := stream.Write(chunk); err != nil {
				return fmt.Errorf("write to stream failed: %s", err)
			}
			return nil
		}),
	); err != nil {
		return fmt.Errorf("failed while streaming: %s", err)
	}
	return nil
}

func read(name string) *x509.Certificate {
	b, _ := pem.Decode(must(os.ReadFile(fmt.Sprintf("%s.crt", name))))
	return must(x509.ParseCertificate(b.Bytes))
}

func must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}
