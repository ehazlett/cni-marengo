package server

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetHostPort(t *testing.T) {
	host := "127.0.0.1"
	port := 9000
	addr := fmt.Sprintf("%s:%d", host, port)

	a, p, err := getHostPort(addr)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, a, host, "host should equal %s", host)
	assert.Equal(t, p, port, "port should equal %d", port)
}

func TestGetTunnelPort(t *testing.T) {
	cfg := &ServerConfig{
		Name:       "test",
		ListenAddr: "127.0.0.1:9000",
	}
	s, err := NewServer(cfg)
	if err != nil {
		t.Fatal(err)
	}

	host := "127.0.0.2"
	port := 9000
	addr := fmt.Sprintf("%s:%d", host, port)

	p, err := s.getTunnelPort(addr)
	if err != nil {
		t.Fatal(err)
	}

	expected := 9003
	assert.Equal(t, p, expected, "expected port %d: received %d", expected, p)
}
