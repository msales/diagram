package mermaid

import (
	"context"
	"net/http"
)

// DefaultPattern is the default mermaid diagram http path.
var DefaultPattern = "/diagram/mermaid"

// DefaultHandler is the default diagram http handler.
var DefaultHandler = NewHandler()

// StartServer starts the http diagram server on the given port
// with the given stats.
func StartServer(server *http.Server, addr string, stats Stater) error {
	server.Addr = addr
	server.Handler = newMux(stats)

	return server.ListenAndServe()
}

// StopServer stops the htt[ health server
func StopServer(server *http.Server) error {
	return server.Shutdown(context.Background())
}

func newMux(stats Stater) http.Handler {
	mux := &http.ServeMux{}
	mux.Handle(DefaultPattern, DefaultHandler.With(stats))

	return mux
}
