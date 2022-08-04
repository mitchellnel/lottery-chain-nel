package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/mitchellnel/lottery-chain-nel/x/lottery/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdChangeOwner() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "change-owner [new-owner]",
		Short: "transfer ownership of the lottery to another account",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argNewOwner := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgChangeOwner(
				clientCtx.GetFromAddress().String(),
				argNewOwner,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
