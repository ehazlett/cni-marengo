package server

import "sort"

type PeerInfo struct {
	Name    string
	Address string
	IPs     map[string]string
}

// getLocalPeers returns the neighboring peers in the cluster
func (s *Server) getLocalPeers(peers []*PeerInfo) []*PeerInfo {
	// if there are only two peers, return
	if len(peers) <= 2 {
		return peers
	}

	// "insert" self node info to sort and get peers
	peers = append(peers, &PeerInfo{
		Name:    s.config.Name,
		Address: s.config.ListenAddr,
	})
	// convert peer info into map to sort
	info := map[string]*PeerInfo{}
	keys := make([]string, 0, len(peers))
	for _, p := range peers {
		info[p.Address] = p
		keys = append(keys, p.Address)
	}

	sort.Strings(keys)

	localPeers := make([]*PeerInfo, 0, 2)
	for k, v := range keys {
		if v == s.config.ListenAddr {
			// get previous and next local peers
			next := k + 1
			prev := k - 1

			// handle edges
			if next >= len(peers) {
				next = 0
			}
			if prev < 0 {
				prev = len(peers) - 1
			}

			n := keys[next]
			p := keys[prev]

			nextInfo := info[n]
			prevInfo := info[p]
			localPeers = append(localPeers, []*PeerInfo{nextInfo, prevInfo}...)
			break
		}
	}

	// TODO: return previous and next peers as "local" peers
	return localPeers
}
