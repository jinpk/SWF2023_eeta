package cli

import (
	"strconv"

	"eeta/x/billboard/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdCreateBillboard() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-billboard [name] [description] [url] [board-type] [final-bid-price-per-minute]",
		Short: "Broadcast message create-billboard",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argName := args[0]
			argDescription := args[1]
			argUrl := args[2]
			argBoardType := args[3]
			argFinalBidPricePerMinute, err := sdk.ParseCoinNormalized(args[4])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateBillboard(
				clientCtx.GetFromAddress().String(),
				argName,
				argDescription,
				argUrl,
				argBoardType,
				argFinalBidPricePerMinute,
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
