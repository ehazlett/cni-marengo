package server

import (
	"encoding/json"

	"github.com/sirupsen/logrus"
)

// heartbeat broadcasts to the cluster the local GRPC connection info
func (s *Server) heartbeat() error {
	logrus.WithFields(logrus.Fields{
		"peers": s.peers,
	}).Debug("current peers")
	logrus.WithFields(logrus.Fields{
		"node":    s.config.Name,
		"address": s.config.ListenAddr,
		"ips":     s.ips,
	}).Debug("heartbeat")
	info := &PeerInfo{
		Name:    s.config.Name,
		Address: s.config.ListenAddr,
		IPs:     s.ips,
	}
	data, err := json.Marshal(info)
	if err != nil {
		return err
	}

	return s.discover.SendEvent(NodeInfoEvent, data, false)
}
