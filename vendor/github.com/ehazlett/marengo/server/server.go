package server

import (
	"sync"

	"github.com/ehazlett/libdiscover"
)

const (
	NodeJoinEvent  = "node-join"
	NodeLeaveEvent = "node-leave"
	NodeInfoEvent  = "marengo-node-info"

	defaultBridgeName = "marengo"
	basePort          = 9000
)

type Server struct {
	config   *ServerConfig
	stopChan chan bool
	discover *libdiscover.Discover
	// peers is a map of name -> address
	peers map[string]*PeerInfo
	// tunnels is a map to peer name -> tunnel interface
	tunnels map[string]string
	// IPs are for the global IPAM service IP -> container id
	ips map[string]string
	mu  *sync.Mutex
}

func NewServer(cfg *ServerConfig) (*Server, error) {
	ch := make(chan bool)
	return &Server{
		stopChan: ch,
		config:   cfg,
		peers:    make(map[string]*PeerInfo),
		tunnels:  make(map[string]string),
		ips:      make(map[string]string),
		mu:       &sync.Mutex{},
	}, nil
}
