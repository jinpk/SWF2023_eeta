package sto

import (
	"math/rand"

	"eeta/testutil/sample"
	stosimulation "eeta/x/sto/simulation"
	"eeta/x/sto/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = stosimulation.FindAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
	_ = rand.Rand{}
)

const (
	opWeightMsgCreateSto = "op_weight_msg_create_sto"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateSto int = 100

	opWeightMsgFund = "op_weight_msg_fund"
	// TODO: Determine the simulation weight value
	defaultWeightMsgFund int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	stoGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&stoGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// ProposalContents doesn't return any content functions for governance proposals.
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateSto int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateSto, &weightMsgCreateSto, nil,
		func(_ *rand.Rand) {
			weightMsgCreateSto = defaultWeightMsgCreateSto
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateSto,
		stosimulation.SimulateMsgCreateSto(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgFund int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgFund, &weightMsgFund, nil,
		func(_ *rand.Rand) {
			weightMsgFund = defaultWeightMsgFund
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgFund,
		stosimulation.SimulateMsgFund(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateSto,
			defaultWeightMsgCreateSto,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				stosimulation.SimulateMsgCreateSto(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgFund,
			defaultWeightMsgFund,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				stosimulation.SimulateMsgFund(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
