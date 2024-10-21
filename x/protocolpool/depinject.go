package protocolpool

import (
	modulev1 "cosmossdk.io/api/cosmos/protocolpool/module/v1"
	"cosmossdk.io/core/address"
	"cosmossdk.io/core/appmodule"
	"cosmossdk.io/core/moduleaccounts"
	"cosmossdk.io/depinject"
	"cosmossdk.io/depinject/appconfig"
	"cosmossdk.io/x/protocolpool/keeper"
	"cosmossdk.io/x/protocolpool/simulation"
	"cosmossdk.io/x/protocolpool/types"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/simsx"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

var _ depinject.OnePerModuleType = AppModule{}

// IsOnePerModuleType implements the depinject.OnePerModuleType interface.
func (am AppModule) IsOnePerModuleType() {}

func init() {
	appconfig.RegisterModule(
		&modulev1.Module{},
		appconfig.Provide(ProvideModule),
	)
}

type ModuleInputs struct {
	depinject.In

	Config                *modulev1.Module
	Codec                 codec.Codec
	Environment           appmodule.Environment
	ModuleAccountsService moduleaccounts.Service
	AddressCdc            address.Codec

	BankKeeper    types.BankKeeper
	StakingKeeper types.StakingKeeper
}

type ModuleOutputs struct {
	depinject.Out

	Keeper         keeper.Keeper
	Module         appmodule.AppModule
	ModuleAccounts []runtime.ModuleAccount
}

func ProvideModule(in ModuleInputs) ModuleOutputs {
	// default to governance authority if not provided
	authority := authtypes.NewModuleAddress(types.GovModuleName)
	if in.Config.Authority != "" {
		authority = authtypes.NewModuleAddressOrBech32Address(in.Config.Authority)
	}

	authorityAddr, err := in.AddressCdc.BytesToString(authority)
	if err != nil {
		panic(err)
	}

	k := keeper.NewKeeper(
		in.Codec,
		in.Environment,
		in.BankKeeper,
		in.StakingKeeper,
		authorityAddr,
		in.AddressCdc,
		in.ModuleAccountsService,
	)
	m := NewAppModule(in.Codec, k, in.BankKeeper)

	return ModuleOutputs{
		Keeper: k,
		Module: m,
		ModuleAccounts: []runtime.ModuleAccount{
			runtime.NewModuleAccount(types.ModuleName),
			runtime.NewModuleAccount(types.ProtocolPoolDistrAccount),
			runtime.NewModuleAccount(types.StreamAccount),
		},
	}
}

// ____________________________________________________________________________

// AppModuleSimulation functions

// GenerateGenesisState creates a randomized GenState of the protocolpool module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
}

// RegisterStoreDecoder registers a decoder for protocolpool module's types
func (am AppModule) RegisterStoreDecoder(sdr simtypes.StoreDecoderRegistry) {
}

// ProposalMsgsX returns all the protocolpool msgs used to simulate governance proposals.
func (am AppModule) ProposalMsgsX(weight simsx.WeightSource, reg simsx.Registry) {
	reg.Add(weight.Get("msg_community_pool_spend", 50), simulation.MsgCommunityPoolSpendFactory())
}

func (am AppModule) WeightedOperationsX(weight simsx.WeightSource, reg simsx.Registry) {
	reg.Add(weight.Get("msg_fund_community_pool", 50), simulation.MsgFundCommunityPoolFactory())
}
