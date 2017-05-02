// copyright 2017 Evan Hazlett <ejhazlett@gmail.com>

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net"

	"github.com/containernetworking/cni/pkg/ns"
	"github.com/containernetworking/cni/pkg/skel"
	"github.com/containernetworking/cni/pkg/types"
	"github.com/containernetworking/cni/pkg/types/current"
	"github.com/containernetworking/cni/pkg/version"
	"github.com/ehazlett/marengo/api"
)

const (
	defaultControlSocket = "unix:///var/run/marengo.sock"
	marengoBridgeName    = "marengo"
	defaultMTU           = 1450
)

type PluginConf struct {
	types.NetConf
	Subnet  string `json:"subnet"`
	Hairpin bool   `json:"hairpin"`
}

func loadConfig(stdin []byte) (*PluginConf, error) {
	conf := PluginConf{}
	if err := json.Unmarshal(stdin, &conf); err != nil {
		return nil, fmt.Errorf("failed to parse network configuration: %v", err)
	}

	return &conf, nil
}

func cmdAdd(args *skel.CmdArgs) error {
	conf, err := loadConfig(args.StdinData)
	if err != nil {
		return err
	}

	c, err := getClient()
	if err != nil {
		return err
	}

	ctx := context.Background()
	resp, err := c.AllocateIP(ctx, &api.IPAMRequest{
		ContainerID: args.ContainerID,
		Subnet:      conf.Subnet,
	})
	if err != nil {
		return err
	}

	_, ipnet, err := net.ParseCIDR(conf.Subnet)
	if err != nil {
		return err
	}

	netns, err := ns.GetNS(args.Netns)
	if err != nil {
		return fmt.Errorf("failed to open netns %q: %v", args.Netns, err)
	}
	defer netns.Close()

	br, err := getBridge()
	if err != nil {
		return err
	}

	ip := net.ParseIP(resp.IP.Address)
	ipn := &net.IPNet{
		IP:   ip,
		Mask: ipnet.Mask,
	}
	_, _, err = setupVeth(netns, br, args.IfName, ipn, defaultMTU, conf.Hairpin)
	if err != nil {
		return err
	}

	result := &current.Result{
		IPs: []*current.IPConfig{{
			Version: resp.IP.Version,
			Address: net.IPNet{
				IP:   ip,
				Mask: ipnet.Mask,
			},
			//Gateway: l.Gateway(),
		}},
		//Routes: []*types.Routes{
		//    {
		//	Dst:
		//    }
		//}
	}

	return types.PrintResult(result, conf.CNIVersion)
}

func cmdDel(args *skel.CmdArgs) error {
	netns, err := ns.GetNS(args.Netns)
	if err != nil {
		return fmt.Errorf("failed to open netns %q: %v", args.Netns, err)
	}
	defer netns.Close()

	// get container ips
	addrs, err := getContainerIPs(netns)
	if err != nil {
		return err
	}

	if err := teardownVeth(netns); err != nil {
		return err
	}

	c, err := getClient()
	if err != nil {
		return err
	}

	ctx := context.Background()
	for _, addr := range addrs {
		if _, err := c.ReleaseIP(ctx, &api.IPReleaseRequest{
			ContainerID: args.ContainerID,
			Address:     addr.IP.String(),
		}); err != nil {
			return err
		}
	}

	return nil
}

func main() {
	skel.PluginMain(cmdAdd, cmdDel, version.PluginSupports("", "0.1.0", "0.2.0", version.Current()))
}
