package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"lottery-chain-nel/x/lottery/types"
)

var _ = strconv.Itoa(0)

func CmdSetupLottery() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "setup-lottery [entrance-fee]",
		Short: "execute first-time setup for the lottery and set entrance fee",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argEntranceFee := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSetupLottery(
				clientCtx.GetFromAddress().String(),
				argEntranceFee,
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
