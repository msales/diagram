package mermaid_test

import (
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
	stat := mermaid.NewStat(topology)

	go mermaid.StartServer("127.0.0.1:8080", stat)
	defer mermaid.StopServer()

	time.Sleep(time.Millisecond)

	resp, err := httpx.Get("http://127.0.0.1:8080/diagram/mermaid")
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}
