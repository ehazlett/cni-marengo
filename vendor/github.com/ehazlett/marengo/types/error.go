package types

import "errors"

var (
	ErrBridgeExists = errors.New("bridge already exists with that name")
	ErrTunnelExists = errors.New("tunnel already exists with that name")
)
