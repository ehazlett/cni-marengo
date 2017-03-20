// copyright 2017 Evan Hazlett <ejhazlett@gmail.com>
package main

import (
	"fmt"
	"net"
	"os"
	"syscall"

	"github.com/containernetworking/cni/pkg/ip"
	"github.com/containernetworking/cni/pkg/ns"
	"github.com/containernetworking/cni/pkg/types/current"
	"github.com/ehazlett/marengo/api"
	"github.com/vishvananda/netlink"
)

func getClient() (api.NetworkManagerClient, error) {
	control := defaultControlSocket
	ctrl := os.Getenv("MARENGO_CONTROL_SOCKET")
	if ctrl != "" {
		control = ctrl
	}
	return api.NewClient(control)
}

func getBridge() (*netlink.Bridge, error) {
	l, err := netlink.LinkByName(marengoBridgeName)
	if err != nil {
		return nil, fmt.Errorf("could not lookup %q: %v", marengoBridgeName, err)
	}
	br, ok := l.(*netlink.Bridge)
	if !ok {
		return nil, fmt.Errorf("%q already exists but is not a bridge", marengoBridgeName)
	}
	return br, nil
}

func setupVeth(netns ns.NetNS, br *netlink.Bridge, ifName string, addr *net.IPNet, mtu int, hairpinMode bool) (*current.Interface, *current.Interface, error) {
	contIface := &current.Interface{}
	hostIface := &current.Interface{}

	err := netns.Do(func(hostNS ns.NetNS) error {
		hostVeth, containerVeth, err := ip.SetupVeth(ifName, mtu, hostNS)
		if err != nil {
			return err
		}
		contIface.Name = containerVeth.Attrs().Name
		contIface.Mac = containerVeth.Attrs().HardwareAddr.String()
		contIface.Sandbox = netns.Path()
		hostIface.Name = hostVeth.Attrs().Name

		// setup addr
		cVeth, err := netlink.LinkByName(contIface.Name)
		if err != nil {
			return err
		}
		ipAddr := &netlink.Addr{IPNet: addr, Label: ""}
		if err := netlink.AddrAdd(cVeth, ipAddr); err != nil {
			return fmt.Errorf("error assigning ip to container peer interface: s", err)
		}

		return nil
	})
	if err != nil {
		return nil, nil, err
	}

	hostVeth, err := netlink.LinkByName(hostIface.Name)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to lookup %q: %v", hostIface.Name, err)
	}
	hostIface.Mac = hostVeth.Attrs().HardwareAddr.String()

	if err := netlink.LinkSetMaster(hostVeth, br); err != nil {
		return nil, nil, fmt.Errorf("failed to connect %q to bridge %v: %v", hostVeth.Attrs().Name, br.Attrs().Name, err)
	}

	if err = netlink.LinkSetHairpin(hostVeth, hairpinMode); err != nil {
		return nil, nil, fmt.Errorf("failed to setup hairpin mode for %v: %v", hostVeth.Attrs().Name, err)
	}

	return hostIface, contIface, nil
}

func getContainerIPs(netns ns.NetNS) ([]netlink.Addr, error) {
	var addrs []netlink.Addr
	err := netns.Do(func(hostNS ns.NetNS) error {
		cVeth, err := netlink.LinkByName("eth0")
		if err != nil {
			return err
		}
		a, err := netlink.AddrList(cVeth, syscall.AF_INET)
		if err != nil {
			return fmt.Errorf("error getting container addresses: %s", err)
		}

		addrs = a

		return nil
	})
	if err != nil {
		return nil, err
	}

	return addrs, nil
}

func teardownVeth(netns ns.NetNS) error {
	err := netns.Do(func(hostNS ns.NetNS) error {
		cVeth, err := netlink.LinkByName("eth0")
		if err != nil {
			return err
		}
		if err := netlink.LinkDel(cVeth); err != nil {
			return fmt.Errorf("error removing container interface: %s", err)
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}
