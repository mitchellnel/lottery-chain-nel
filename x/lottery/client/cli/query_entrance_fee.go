package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/mitchellnel/lottery-chain-nel/x/lottery/types"
	"github.com/spf13/cobra"
)

func CmdShowEntranceFee() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-entrance-fee",
		Short: "shows entrance-fee",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetEntranceFeeRequest{}

			res, err := queryClient.EntranceFee(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
