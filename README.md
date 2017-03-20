# CNI Marengo Plugin
This is a [CNI](https://github.com/containernetworking/cni) plugin for the
[Marengo](https://github.com/ehazlett/marengo) overlay manager.

# Build
Clone and `make`

# Usage
Build or download a release and put on your path like other CNI plugins.

You must either run this on a Marengo node or specify the `MARENGO_CONTROL_SOCK`
environment variable.  The format is `unix:///var/run/marengo.sock` or
`tcp://127.0.0.1:8080` for either a unix or tcp socket.

This is a complete plugin.  You do not need to specify IPAM as this plugin
uses Marengo for addressing.

# Configuration
The following is an example config:

```
{
    "cniVersion": "0.2.0",
    "name": "marengo",
    "type": "cni-marengo",
    "bridge": "marengo",
    "hairpin": true,
    "subnet": "172.50.0.0/24"
}
```

- `name`: user specified network name
- `type`: this is the plugin name
- `bridge`: this must be the marengo bridge (default: `marengo`)
- `hairpin`: enable hairpinning
- `subnet`: the subnet on which to request an IP

This plugin can be used with [Circuit](https://github.com/ehazlett/circuit)
for connecting / disconnecting containers from Marengo overlay networks.
