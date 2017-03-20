package utils

import (
	"net"
	"strings"

	"github.com/ehazlett/marengo/types"
	"github.com/sirupsen/logrus"
	"github.com/vishvananda/netlink"
)

const (
	DefaultMTU = 1500
)

// CreateBridge creates a new local bridge interface
func CreateBridge(name string) error {
	_, brErr := net.InterfaceByName(name)
	if brErr == nil {
		return types.ErrBridgeExists
	}

	if !strings.Contains(brErr.Error(), "no such network interface") {
		return brErr
	}

	attrs := netlink.NewLinkAttrs()
	attrs.Name = name
	attrs.MTU = DefaultMTU

	br := &netlink.Bridge{
		LinkAttrs: attrs,
	}

	if err := netlink.LinkAdd(br); err != nil {
		return err
	}

	if err := netlink.LinkSetUp(br); err != nil {
		return err
	}

	return nil
}

func ConnectToBridge(ifaceName string, bridgeName string) error {
	if _, err := net.InterfaceByName(ifaceName); err != nil {
		return err
	}

	if _, err := net.InterfaceByName(bridgeName); err != nil {
		return err
	}

	attrs := netlink.NewLinkAttrs()
	attrs.Name = ifaceName
	t := &netlink.Vxlan{
		LinkAttrs: attrs,
	}

	brAttrs := netlink.NewLinkAttrs()
	brAttrs.Name = bridgeName
	br := &netlink.Bridge{
		LinkAttrs: brAttrs,
	}

	if err := netlink.LinkSetMaster(t, br); err != nil {
		return err
	}

	return nil
}

func DeleteBridge(name string) error {
	logrus.WithFields(logrus.Fields{
		"name": name,
	}).Debug("removing bridge")

	if _, err := net.InterfaceByName(name); err != nil {
		return err
	}

	attrs := netlink.NewLinkAttrs()
	attrs.Name = name

	br := &netlink.Bridge{
		LinkAttrs: attrs,
	}

	if err := netlink.LinkSetDown(br); err != nil {
		return err
	}

	if err := netlink.LinkDel(br); err != nil {
		return err
	}

	return nil
}
