package testhelp

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

// Kill server is responsible for sending a context.cancel to the actual application because unless the app self exits
// "go test" will not generate a coverage report.

type KillServer struct {
	server http.Server
	cancel context.CancelFunc
}

func NewKillServer(addr string, cancel context.CancelFunc) *KillServer {
	return &KillServer{
		server: http.Server{
			Addr: addr,
		},
		cancel: cancel,
	}
}

func (s *KillServer) Start() {
	s.server.Handler = s

	err := s.server.ListenAndServe()
	if err != nil {
		fmt.Println("KillServer Error:", err)
	}
}

func (s *KillServer) Shutdown(ctx context.Context) {
	s.server.Shutdown(ctx)
}

func (s *KillServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	log.Print("Killing Server ")
	// cancel the context
	s.cancel()
}
