package mermaid

import "net/http"

// Reporter represents an a health reporter.
type Stater interface {
	// GetStats gets diagram string
	GetStats() (string, error)
}

// Handler is an http health handler.
type Handler struct {
	stats   []Stater
	showErr bool
}

// NewHandler creates a new Handler instance.
func NewHandler() *Handler {
	return &Handler{}
}

// With adds stats to the handler.
func (h *Handler) With(stats ...Stater) *Handler {
	h.stats = stats
	return h
}

// ServeHTTP serves an http request.
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, stat := range h.stats {
		stats, err := stat.GetStats()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)

			return
		}
		w.Write([]byte(stats))
	}
}
