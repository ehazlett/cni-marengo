package server

import "testing"

func TestGetLocalPeers(t *testing.T) {
	cfg := &ServerConfig{
		Name:       "test",
		ListenAddr: "10.0.1.5:9000",
	}

	peers := []*PeerInfo{
		{
			Name:    "node-20",
			Address: "10.0.1.20:8080",
		},
		{
			Name:    "node-10",
			Address: "10.0.1.10:8080",
		},
		{
			Name:    "node-01",
			Address: "10.0.1.1:8080",
		},
		{
			Name:    "node-02",
			Address: "10.0.1.2:5050",
		},
	}
	s, err := NewServer(cfg)
	if err != nil {
		t.Fatal(err)
	}

	localPeers := s.getLocalPeers(peers)
	expected := map[string]struct{}{
		"10.0.1.2:5050":  struct{}{},
		"10.0.1.10:8080": struct{}{},
	}
	for _, p := range localPeers {
		if _, ok := expected[p.Address]; !ok {
			t.Fatalf("expected peers %+v; received %+v", expected, localPeers)
		}
	}
}

func TestGetLocalPeersMinimum(t *testing.T) {
	cfg := &ServerConfig{
		Name:       "test",
		ListenAddr: "10.0.1.5:9000",
	}

	peers := []*PeerInfo{
		{
			Name:    "node-20",
			Address: "10.0.1.20:8080",
		},
		{
			Name:    "node-10",
			Address: "10.0.1.10:8080",
		},
	}
	s, err := NewServer(cfg)
	if err != nil {
		t.Fatal(err)
	}

	localPeers := s.getLocalPeers(peers)
	expected := map[string]struct{}{
		"10.0.1.20:8080": struct{}{},
		"10.0.1.10:8080": struct{}{},
	}
	for _, p := range localPeers {
		if _, ok := expected[p.Address]; !ok {
			t.Fatalf("expected peers %+v; received %+v", expected, localPeers)
		}
	}
}
