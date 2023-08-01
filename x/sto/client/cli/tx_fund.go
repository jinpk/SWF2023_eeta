package cli

import (
	"strconv"

	"eeta/x/sto/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdFund() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "fund [billboard-id] [sto-id] [amount]",
		Short: "Broadcast message fund",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argBillboardId, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}

			argStoId, err := cast.ToUint64E(args[1])
			if err != nil {
				return err
			}
			argAmount, err := sdk.ParseCoinNormalized(args[2])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgFund(
				clientCtx.GetFromAddress().String(),
				argBillboardId,
				argStoId,
				argAmount,
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
