package simulation

import (
	"math/rand"

	"eeta/x/bid/keeper"
	"eeta/x/bid/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgBidding(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgBidding{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the Bidding simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "Bidding simulation not implemented"), nil, nil
	}
}
