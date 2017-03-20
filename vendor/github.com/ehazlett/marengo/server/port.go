package server

import (
	"fmt"
	"strconv"
	"strings"
)

func getHostPort(addr string) (string, int, error) {
	parts := strings.Split(addr, ":")
	if len(parts) != 2 {
		return "", -1, fmt.Errorf("expected peer address to be in <host>:<port> format: %s", addr)
	}

	port, err := strconv.Atoi(parts[1])
	if err != nil {
		return "", -1, err
	}

	return parts[0], port, nil
}

// getTunnelPort returns a predictable port based upon the server listen
// address and the peer address
func (s *Server) getTunnelPort(peerAddress string) (int, error) {
	peerAddr, _, err := getHostPort(peerAddress)
	if err != nil {
		return -1, err
	}
	nodeAddr, _, err := getHostPort(s.config.ListenAddr)
	if err != nil {
		return -1, err
	}

	peerParts := strings.Split(peerAddr, ".")
	nodeParts := strings.Split(nodeAddr, ".")

	if len(peerParts) != 4 {
		return -1, fmt.Errorf("expected peer host to be in x.x.x.x format: %s", peerAddr)
	}
	if len(nodeParts) != 4 {
		return -1, fmt.Errorf("expected node host to be in x.x.x.x format: %s", nodeAddr)
	}

	p, err := strconv.Atoi(peerParts[3])
	if err != nil {
		return -1, err
	}
	n, err := strconv.Atoi(nodeParts[3])
	if err != nil {
		return -1, err
	}

	port := basePort + p + n
	return port, nil
}
