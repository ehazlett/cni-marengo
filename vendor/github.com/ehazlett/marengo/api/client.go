package api

import (
	"net"
	"time"

	"github.com/ehazlett/marengo/utils"
	"google.golang.org/grpc"
)

// NewClient returns a `NetworkManagerClient` for use with the API
func NewClient(addr string) (NetworkManagerClient, error) {
	proto, host, err := utils.GetProtoHost(addr)
	if err != nil {
		return nil, err
	}
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	opts = append(opts, grpc.WithDialer(
		func(addr string, timeout time.Duration) (net.Conn, error) {
			return net.DialTimeout(proto, host, timeout)
		},
	))
	conn, err := grpc.Dial(host, opts...)
	if err != nil {
		return nil, err
	}

	client := NewNetworkManagerClient(conn)
	return client, nil
}
