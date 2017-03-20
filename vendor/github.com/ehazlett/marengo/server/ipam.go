package server

import (
	"encoding/binary"
	"fmt"
	"math/rand"
	"net"
	"time"

	"github.com/ehazlett/marengo/api"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

// AllocateIP allocates a new IP address
func (s *Server) AllocateIP(ctx context.Context, req *api.IPAMRequest) (*api.IPAMResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	logrus.WithFields(logrus.Fields{
		"container": req.ContainerID,
		"subnet":    req.Subnet,
	}).Debug("allocating IP")

	var ip net.IP
	// retry until unique is received
	for i := 0; i < 10; i++ {
		i, err := getIP(req.Subnet)
		if err != nil {
			logrus.Errorf("error allocating ip: %s", err)
			continue
		}

		if _, ok := s.ips[i.String()]; !ok {
			ip = i
			break
		}
	}

	// reserve
	s.ips[ip.String()] = req.ContainerID

	logrus.WithFields(logrus.Fields{
		"ip":        ip.String(),
		"container": req.ContainerID,
	}).Info("allocated local IP")

	return &api.IPAMResponse{
		IP: &api.IPConfig{
			Version: "4", // TODO: ipv6?
			Address: ip.String(),
			// TODO: Gateway ?
		},
		DNS: nil, // TODO
	}, nil
}

// ReleaseIP releases an IP back to the pool
func (s *Server) ReleaseIP(ctx context.Context, req *api.IPReleaseRequest) (*api.IPReleaseResponse, error) {
	// TODO: release ip

	// TODO: broadcast removal to cluster
	return nil, nil
}

func getIP(subnet string) (net.IP, error) {
	_, sub, err := net.ParseCIDR(subnet)
	if err != nil {
		return nil, err
	}
	o := sub.IP.To4()
	// add new source; default is deterministic
	x := rand.NewSource(time.Now().UnixNano())
	r := rand.New(x)
	min := 2   // start at 2
	max := 254 // no greater than 254
	d := min + r.Intn(max-min)

	if len(o) < 3 {
		return nil, fmt.Errorf("error allocating ip for %v; invalid length", subnet)
	}
	ip := net.IPv4(o[0], o[1], o[2], byte(d))

	return ip, nil
}

func ipToInt(ip net.IP) uint32 {
	if len(ip) == 16 {
		return binary.BigEndian.Uint32(ip[12:16])
	}
	return binary.BigEndian.Uint32(ip)
}

func intToIP(n uint32) net.IP {
	ip := make(net.IP, 4)
	binary.BigEndian.PutUint32(ip, n)
	return ip
}
