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

func CmdCreateAuction() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-auction [billboard-id] [start] [end] [amount] [ad-url]",
		Short: "Broadcast message create-auction",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argBillboardId, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}
			argStart, err := cast.ToUint64E(args[1])
			if err != nil {
				return err
			}
			argEnd, err := cast.ToUint64E(args[2])
			if err != nil {
				return err
			}
			argAmount, err := sdk.ParseCoinNormalized(args[3])
			if err != nil {
				return err
			}
			argAdUrl := args[4]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateAuction(
				clientCtx.GetFromAddress().String(),
				argBillboardId,
				argStart,
				argEnd,
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
