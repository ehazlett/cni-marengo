package utils

import (
	"net"
	"strings"

	"github.com/ehazlett/marengo/types"
	"github.com/sirupsen/logrus"
	"github.com/vishvananda/netlink"
)

const (
	defaultVxlanMTU = 1450
)

// CreateVxlan creates a new VXLAN tunnel
func CreateVxlan(name string, vxlanID int, src net.IP, port int, ttl int) error {
	_, ifErr := net.InterfaceByName(name)
	if ifErr == nil {
		return types.ErrTunnelExists
	}

	if !strings.Contains(ifErr.Error(), "no such network interface") {
		return ifErr
	}

	attrs := netlink.NewLinkAttrs()
	attrs.Name = name
	attrs.MTU = defaultVxlanMTU

	t := &netlink.Vxlan{
		LinkAttrs: attrs,
		VxlanId:   vxlanID,
		Group:     src,
		TTL:       ttl,
		Port:      port,
	}

	if err := netlink.LinkAdd(t); err != nil {
		return err
	}

	if err := netlink.LinkSetUp(t); err != nil {
		return err
	}

	return nil
}

func DeleteVxlan(name string) error {
	logrus.WithFields(logrus.Fields{
		"name": name,
	}).Debug("removing vxlan")
	if _, err := net.InterfaceByName(name); err != nil {
		return err
	}

	attrs := netlink.NewLinkAttrs()
	attrs.Name = name

	t := &netlink.Vxlan{
		LinkAttrs: attrs,
	}

	if err := netlink.LinkSetDown(t); err != nil {
		return err
	}

	if err := netlink.LinkDel(t); err != nil {
		return err
	}

	return nil
}
