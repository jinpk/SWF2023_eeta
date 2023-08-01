package simulation

import (
	"math/rand"

	"eeta/x/sto/keeper"
	"eeta/x/sto/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgCreateSto(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgCreateSto{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the CreateSto simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "CreateSto simulation not implemented"), nil, nil
	}
}
