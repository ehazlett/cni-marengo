package server

import "time"

type ServerConfig struct {
	Name          string
	ListenAddr    string
	ControlSocket string
	BindAddr      string
	AdvertiseAddr string
	JoinAddr      string
	NodeTimeout   time.Duration
	Debug         bool
}
