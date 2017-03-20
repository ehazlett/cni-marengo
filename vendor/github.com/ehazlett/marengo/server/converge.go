package server

import (
	"fmt"
	"net"

	"github.com/ehazlett/marengo/types"
	"github.com/ehazlett/marengo/utils"
	"github.com/sirupsen/logrus"
)

func (s *Server) getConnectedPeers() []*PeerInfo {
	peers := []*PeerInfo{}
	for _, p := range s.peers {
		peers = append(peers, p)
	}

	return peers
}

// converge ensures that tunnels exist for peers
func (s *Server) converge() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if err := utils.CreateBridge(defaultBridgeName); err != nil && err != types.ErrBridgeExists {
		return err
	}

	// TODO: check peer tunnels
	// TODO: max peer tunnels
	peers := make([]*PeerInfo, 0, len(s.peers))
	for _, p := range s.peers {
		peers = append(peers, p)
	}
	localPeers := s.getLocalPeers(peers)
	localPeerInfo := map[string]*PeerInfo{}
	for _, p := range localPeers {
		localPeerInfo[p.Name] = p
	}

	for _, info := range localPeers {
		name := info.Name
		address := info.Address
		logrus.WithFields(logrus.Fields{
			"name": name,
			"addr": info,
		}).Debug("checking peer tunnel")

		port, err := s.getTunnelPort(address)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"name":    name,
				"address": address,
			}).Errorf("error getting tunnel port: %s", err)
			continue
		}

		host, _, err := getHostPort(address)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"name":    name,
				"address": address,
			}).Errorf("error getting host address for peer: %s", err)
			continue
		}

		// create tunnel
		ip := net.ParseIP(host)
		iface := fmt.Sprintf("ol-%s-%d", defaultBridgeName, port)
		ttl := 60

		// TODO: allow custom vxlan id
		if err := utils.CreateVxlan(iface, 1024, ip, port, ttl); err != nil && err != types.ErrTunnelExists {
			logrus.WithFields(logrus.Fields{
				"name":    name,
				"address": address,
			}).Errorf("error creating vxlan %s: %s", iface, err)
			continue
		}

		// update tunnels
		s.tunnels[name] = iface

		// connect to bridge
		if err := utils.ConnectToBridge(iface, defaultBridgeName); err != nil {
			logrus.WithFields(logrus.Fields{
				"name":    name,
				"address": address,
			}).Errorf("error connecting iface %s to bridge %s: %s", iface, defaultBridgeName, err)
			continue
		}
	}

	// remove tunnels that are no longer needed (peer is gone)
	for name, tunnel := range s.tunnels {
		logrus.WithFields(logrus.Fields{
			"name":   name,
			"tunnel": tunnel,
		}).Debug("checking current tunnel for removal")
		if _, ok := localPeerInfo[name]; !ok {
			logrus.WithFields(logrus.Fields{
				"name": name,
			}).Info("removing tunnel")
			if err := utils.DeleteVxlan(tunnel); err != nil {
				logrus.Errorf("error removing tunnel: %s", err)
				continue
			}

			delete(s.tunnels, name)
		}

	}

	logrus.WithFields(logrus.Fields{
		"peers":   localPeerInfo,
		"tunnels": s.tunnels,
		"ips":     s.ips,
	}).Debug("converged")

	return nil
}
