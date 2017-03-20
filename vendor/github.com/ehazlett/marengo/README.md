```
_______ _________ __________________________ ______
__  __ `__ \  __ `/_  ___/  _ \_  __ \_  __ `/  __ \
_  / / / / / /_/ /_  /   /  __/  / / /  /_/ // /_/ /
/_/ /_/ /_/\__,_/ /_/    \___//_/ /_/_\__, / \____/
                                     /____/
```

Overlay Network Manager

Marengo is an overlay network manager.  It forms an overlay mesh (VXLAN)
automatically as nodes join and leave.  Upon joining, node information will
be broadcast throughout the cluster via [Gossip](https://www.serf.io/docs/internals/gossip.html).
The nodes will each "converge" creating an overlay network.  Marengo uses
standard functionality included in modern (3.18+ but newer the better) Linux
kernels.  Marengo exposes a simple bridge that can be used itself to access
the overlay or as a master interface for others to use (i.e. containers).

Marengo is stateless and each node creates and manages its tunnels predictively
so there is no need for a datastore.  If a node crashes, the cluster will
converge and automatically re-configure the tunnels needed to keep the overlay
stable.  Once it returns, tunnels are re-calculated the overlay is extended.

Here is an example of running `iperf` while a node is taken down and the
overlay converges:

```
[  3] local 172.42.0.10 port 45818 connected with 172.42.0.20 port 5001
[ ID] Interval       Transfer     Bandwidth
[  3]  0.0- 1.0 sec  44.6 MBytes   374 Mbits/sec
[  3]  1.0- 2.0 sec  64.0 MBytes   537 Mbits/sec
[  3]  2.0- 3.0 sec  74.5 MBytes   625 Mbits/sec
[  3]  3.0- 4.0 sec  77.4 MBytes   649 Mbits/sec
[  3]  4.0- 5.0 sec  73.8 MBytes   619 Mbits/sec
[  3]  5.0- 6.0 sec  22.9 MBytes   192 Mbits/sec
[  3]  6.0- 7.0 sec  0.00 Bytes  0.00 bits/sec
[  3]  7.0- 8.0 sec  0.00 Bytes  0.00 bits/sec
[  3]  8.0- 9.0 sec  0.00 Bytes  0.00 bits/sec
[  3]  9.0-10.0 sec  0.00 Bytes  0.00 bits/sec
[  3] 10.0-11.0 sec  0.00 Bytes  0.00 bits/sec
[  3] 11.0-12.0 sec  12.5 MBytes   105 Mbits/sec
[  3] 12.0-13.0 sec   137 MBytes  1.15 Gbits/sec
[  3] 13.0-14.0 sec   144 MBytes  1.20 Gbits/sec
[  3] 14.0-15.0 sec   151 MBytes  1.27 Gbits/sec
```

Notice how there was a five second delay while the overlay converged.

# Usage
Marengo is a single binary.  You will need a modern Linux kernel (the
newer the better).

## Start Initial Node

```
$> marengo -D server
```

This should automatically find the default IP that is used to access public
but if you want to use another IP address, run `marengo server -h` for all
options (specifically you will want `--bind-addr` and `--advertise-addr`).

## Connect Nodes
Since Marengo uses Gossip, it is not necessary to know all of the nodes in
the cluster; it will find them itself.  To join another node, simply specify
a single known member:

```
$> marengo -D server -j <ip>:<port>
```

For example:

```
$> marengo -D server -j 10.0.0.1:7946
```

This will join the node to the cluster and within a few seconds the cluster
should converge:

```
$> marengo -D server
INFO[0000] marengo                                       controlSocket="unix:///var/run/marengo.sock" listenAddr="10.255.0.1:8080" version="0.1.0 (7e2a210)"
DEBU[0000] current peers                                 peers=map[]
DEBU[0000] cluster event                                 data="map[addr:10.255.0.1:7946 name:xps]" name=node-join payload="{\"addr\":\"10.255.0.1:7946\",\"name\":\"xps\"}"
DEBU[0000] cluster event                                 data="map[Name:xps Address:10.255.0.1:8080]" name=marengo-node-info payload="{\"Name\":\"xps\",\"Address\":\"10.255.0.1:8080\"}"
DEBU[0000] converge interval                             interval=6s
DEBU[0000] converged                                     peers=map[] tunnels=map[]
DEBU[0001] cluster event                                 data="map[Name:node-00 Address:10.255.0.171:8080]" name=marengo-node-info payload="{\"Name\":\"node-00\",\"Address\":\"10.255.0.171:8080\"}"
DEBU[0001] added node to peers                           address="10.255.0.171:8080" node=node-00
DEBU[0001] cluster event                                 data="map[addr:10.255.0.171:7946 name:node-00]" name=node-join payload="{\"addr\":\"10.255.0.171:7946\",\"name\":\"node-00\"}"
DEBU[0005] current peers                                 peers=map[node-00:0xc4200d9980]
DEBU[0005] cluster event                                 data="map[Name:xps Address:10.255.0.1:8080]" name=marengo-node-info payload="{\"Name\":\"xps\",\"Address\":\"10.255.0.1:8080\"}"
DEBU[0006] checking peer tunnel                          addr=&{node-00 10.255.0.171:8080} name=node-00
DEBU[0006] checking current tunnel for removal           name=node-00 tunnel=ol-marengo-9172
DEBU[0006] converged                                     peers=map[node-00:0xc4200d9980] tunnels=map[node-00:ol-marengo-9172]
```

This will create a bridge and an overlay tunnel:

```
344: marengo: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1450 qdisc noqueue state UP group default qlen 1000
    link/ether 5e:3c:35:53:1c:22 brd ff:ff:ff:ff:ff:ff
    inet6 fe80::1c33:b2ff:fe73:84d4/64 scope link
       valid_lft forever preferred_lft forever
346: ol-marengo-9172: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1450 qdisc noqueue master marengo state UNKNOWN group default qlen 1000
    link/ether 5e:3c:35:53:1c:22 brd ff:ff:ff:ff:ff:ff
    inet6 fe80::5c3c:35ff:fe53:1c22/64 scope link
       valid_lft forever preferred_lft forever
```

You can either use the `marengo` bridge as a master or assign an IP to use.
If you want to use the bridge directly, you will need to assign IP addresses
to the `marengo` bridge on each node (using different IPs).

For example:

```
$> ip addr add 172.42.0.10/24 dev marengo
```

Once IPs are assigned you should be able to ping between the two.

Marengo will automatically cleanup and remove tunnels when either a node
leaves or it is shut down.

Note: Marengo will not remove the bridge as it can be useful to leave for
external automation (i.e. containers).

# Scale
Marengo was built with scale in mind, however it has not been tested
over a few dozen nodes.  There are optimizations in place for using at scale.
  For example, each node will enable a maximum of two tunnels once the cluster
reaches three nodes.  You can also adjust the `--node-timeout` parameter
to control the communication interval if you have a large cluster.
