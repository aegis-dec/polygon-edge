package ibft

import (
	"github.com/aegis-dec/polygon-edge/command/helper"
	"github.com/aegis-dec/polygon-edge/command/ibft/candidates"
	"github.com/aegis-dec/polygon-edge/command/ibft/propose"
	"github.com/aegis-dec/polygon-edge/command/ibft/quorum"
	"github.com/aegis-dec/polygon-edge/command/ibft/snapshot"
	"github.com/aegis-dec/polygon-edge/command/ibft/status"
	_switch "github.com/aegis-dec/polygon-edge/command/ibft/switch"
	"github.com/spf13/cobra"
)

func GetCommand() *cobra.Command {
	ibftCmd := &cobra.Command{
		Use:   "ibft",
		Short: "Top level IBFT command for interacting with the IBFT consensus. Only accepts subcommands.",
	}

	helper.RegisterGRPCAddressFlag(ibftCmd)

	registerSubcommands(ibftCmd)

	return ibftCmd
}

func registerSubcommands(baseCmd *cobra.Command) {
	baseCmd.AddCommand(
		// ibft status
		status.GetCommand(),
		// ibft snapshot
		snapshot.GetCommand(),
		// ibft propose
		propose.GetCommand(),
		// ibft candidates
		candidates.GetCommand(),
		// ibft switch
		_switch.GetCommand(),
		// ibft quorum
		quorum.GetCommand(),
	)
}
