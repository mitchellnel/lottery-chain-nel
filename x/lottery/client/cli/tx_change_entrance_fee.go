package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"lottery-chain-nel/x/lottery/types"
)

var _ = strconv.Itoa(0)

func CmdChangeEntranceFee() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "change-entrance-fee [new-entrance-fee]",
		Short: "change the entrance fee for the lottery -- lottery must be closed",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argNewEntranceFee, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgChangeEntranceFee(
				clientCtx.GetFromAddress().String(),
				argNewEntranceFee,
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
