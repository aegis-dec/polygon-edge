package root

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/aegis-dec/polygon-edge/command/backup"
	"github.com/aegis-dec/polygon-edge/command/bridge"
	"github.com/aegis-dec/polygon-edge/command/genesis"
	"github.com/aegis-dec/polygon-edge/command/helper"
	"github.com/aegis-dec/polygon-edge/command/ibft"
	"github.com/aegis-dec/polygon-edge/command/license"
	"github.com/aegis-dec/polygon-edge/command/monitor"
	"github.com/aegis-dec/polygon-edge/command/peers"
	"github.com/aegis-dec/polygon-edge/command/polybft"
	"github.com/aegis-dec/polygon-edge/command/polybftsecrets"
	"github.com/aegis-dec/polygon-edge/command/regenesis"
	"github.com/aegis-dec/polygon-edge/command/rootchain"
	"github.com/aegis-dec/polygon-edge/command/secrets"
	"github.com/aegis-dec/polygon-edge/command/server"
	"github.com/aegis-dec/polygon-edge/command/status"
	"github.com/aegis-dec/polygon-edge/command/txpool"
	"github.com/aegis-dec/polygon-edge/command/version"
)

type RootCommand struct {
	baseCmd *cobra.Command
}

func NewRootCommand() *RootCommand {
	rootCommand := &RootCommand{
		baseCmd: &cobra.Command{
			Short: "Polygon Edge is a framework for building Ethereum-compatible Blockchain networks",
		},
	}

	helper.RegisterJSONOutputFlag(rootCommand.baseCmd)

	rootCommand.registerSubCommands()

	return rootCommand
}

func (rc *RootCommand) registerSubCommands() {
	rc.baseCmd.AddCommand(
		version.GetCommand(),
		txpool.GetCommand(),
		status.GetCommand(),
		secrets.GetCommand(),
		peers.GetCommand(),
		rootchain.GetCommand(),
		monitor.GetCommand(),
		ibft.GetCommand(),
		backup.GetCommand(),
		genesis.GetCommand(),
		server.GetCommand(),
		license.GetCommand(),
		polybftsecrets.GetCommand(),
		polybft.GetCommand(),
		bridge.GetCommand(),
		regenesis.GetCommand(),
	)
}

func (rc *RootCommand) Execute() {
	if err := rc.baseCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)

		os.Exit(1)
	}
}
