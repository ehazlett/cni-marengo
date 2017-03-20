package server

import (
	"github.com/ehazlett/marengo/utils"
	"github.com/sirupsen/logrus"
)

func (s *Server) Stop() error {
	logrus.Info("stopping server")
	defer func() {
		s.stopChan <- true
	}()

	// cleanup
	logrus.Info("removing tunnels")
	for name, tunnel := range s.tunnels {
		if err := utils.DeleteVxlan(tunnel); err != nil {
			logrus.Errorf("error removing tunnel %s (%s): %s", name, tunnel, err)
		}
	}

	return nil
}
