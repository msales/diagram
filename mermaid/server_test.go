package mermaid_test

import (
	"net/http"
	"testing"
	"time"

	"github.com/msales/diagram/mermaid"
	"github.com/msales/pkg/v3/httpx"
	"github.com/msales/streams/v2"
	"github.com/stretchr/testify/assert"
)

func TestStartServerWithTopology(t *testing.T) {
	builder := streams.NewStreamBuilder()
	builder.Source("event-source", &sourceMock{"test2"})

	topology, errs := builder.Build()
	assert.Nil(t, errs)
	statWithTopology := mermaid.NewStat(topology)
	serverWithTopology := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	go mermaid.StartServer(serverWithTopology, "127.0.0.1:8080", statWithTopology)
	defer mermaid.StopServer(serverWithTopology)

	time.Sleep(time.Millisecond)

	resp, err := httpx.Get("http://127.0.0.1:8080/diagram/mermaid")

	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestStartServerWithEmptyTopology(t *testing.T) {
	serverWithoutTopology := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}
	emptyTopology, errs := streams.NewStreamBuilder().Build()
	assert.Nil(t, errs)
	statWihtoutTopology := mermaid.NewStat(emptyTopology)

	go mermaid.StartServer(serverWithoutTopology, "127.0.0.1:8081", statWihtoutTopology)
	defer mermaid.StopServer(serverWithoutTopology)

	time.Sleep(time.Millisecond)

	respNoTopology, err := httpx.Get("http://127.0.0.1:8081/diagram/mermaid")

	assert.NoError(t, err)
	assert.Equal(t, 500, respNoTopology.StatusCode)
}
