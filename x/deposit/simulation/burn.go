package simulation

import (
	"math/rand"

	"eeta/x/deposit/keeper"
	"eeta/x/deposit/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgBurn(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgBurn{
			Sender: simAccount.Address.String(),
		}

		// TODO: Handling the Burn simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "Burn simulation not implemented"), nil, nil
	}
}