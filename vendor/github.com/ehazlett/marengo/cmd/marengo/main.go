package main

import (
	"os"

	"github.com/ehazlett/marengo/cmd/marengo/commands"
	"github.com/ehazlett/marengo/version"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = version.Name()
	app.Usage = version.Description()
	app.Version = version.FullVersion()
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "debug, D",
			Usage: "Enable debug logging",
		},
	}
	app.Commands = []cli.Command{
		commands.ServerCmd,
	}

	app.Run(os.Args)
}
