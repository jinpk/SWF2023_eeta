package cli

import (
	"strconv"

	"eeta/x/sto/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdCreateSto() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-sto [billboard-id] [organizer-url] [organizer-image-url] [name] [start] [end] [user-share] [organizer-share]",
		Short: "Broadcast message create-sto",
		Args:  cobra.ExactArgs(8),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argBillboardId, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}
			argOrganizerUrl := args[1]
			argOrganizerImageUrl := args[2]
			argName := args[3]
			argStart, err := cast.ToInt32E(args[4])
			if err != nil {
				return err
			}
			argEnd, err := cast.ToInt32E(args[5])
			if err != nil {
				return err
			}
			argUserShare, err := cast.ToInt32E(args[6])
			if err != nil {
				return err
			}
			argOrganizerShare, err := cast.ToInt32E(args[7])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateSto(
				clientCtx.GetFromAddress().String(),
				argBillboardId,
				argOrganizerUrl,
				argOrganizerImageUrl,
				argName,
				argStart,
				argEnd,
				argUserShare,
				argOrganizerShare,
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
