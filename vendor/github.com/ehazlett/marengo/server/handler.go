package server

import (
	"encoding/json"
	"fmt"

	"github.com/ehazlett/libdiscover"
	"github.com/sirupsen/logrus"
)

func (s *Server) eventHandler(e libdiscover.Event) error {
	logrus.WithFields(logrus.Fields{
		"name":    e.Name,
		"payload": string(e.Payload),
		"data":    fmt.Sprintf("%+v", e.Data),
	}).Debug("cluster event")

	switch e.Name {
	case NodeJoinEvent:
	case NodeLeaveEvent:
		return s.handleLeaveEvent(e)
	case NodeInfoEvent:
		return s.handleInfoEvent(e)
	}

	return nil
}

func getNodeInfo(data []byte) (*PeerInfo, error) {
	var info *PeerInfo
	if err := json.Unmarshal(data, &info); err != nil {
		return info, err
	}

	return info, nil
}

func (s *Server) handleInfoEvent(e libdiscover.Event) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	info, err := getNodeInfo(e.Payload)
	if err != nil {
		return err
	}
	// don't add self
	if info.Name == s.config.Name {
		return nil
	}

	// update local peer info
	// TODO: add a TTL to detect dead peers (that haven't left cleanly)
	if _, ok := s.peers[info.Name]; !ok {
		s.peers[info.Name] = info
		logrus.WithFields(logrus.Fields{
			"node":    info.Name,
			"address": info.Address,
		}).Debug("added node to peers")
	}

	// update local IPs with used from other nodes
	for ip, container := range info.IPs {
		v, ok := s.ips[ip]
		if !ok {
			logrus.WithFields(logrus.Fields{
				"node":      info.Name,
				"ip":        ip,
				"container": container,
			}).Debug("reserving ip from cluster")
			s.ips[ip] = container
			continue
		}
		if ok && container != v {
			logrus.Errorf("ip allocated for different containers: %s local=%s %s=%s", ip, v, info.Name, container)
			continue
		}
	}
	return nil
}

func (s *Server) handleLeaveEvent(e libdiscover.Event) error {
	info, err := getNodeInfo(e.Payload)
	if err != nil {
		return err
	}
	// remove local peer
	delete(s.peers, info.Name)
	logrus.WithFields(logrus.Fields{
		"node": info.Name,
	}).Debug("removed node from peers")
	return nil
}
