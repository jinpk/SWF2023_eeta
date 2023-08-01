package cli

import (
	"strconv"

	"eeta/x/bid/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdBidding() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "bidding [billboard-id] [auction-id] [amount] [ad-url]",
		Short: "Broadcast message bidding",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argBillboardId, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}
			argAuctionId, err := cast.ToUint64E(args[1])
			if err != nil {
				return err
			}
			argAmount, err := sdk.ParseCoinNormalized(args[2])
			if err != nil {
				return err
			}
			argAdUrl := args[3]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgBidding(
				clientCtx.GetFromAddress().String(),
				argBillboardId,
				argAuctionId,
				argAmount,
				argAdUrl,
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
