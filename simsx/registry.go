package simsx

import (
	"context"
	"iter"
	"math/rand"
	"slices"
	"strings"
	"time"

	"golang.org/x/exp/maps"

	"cosmossdk.io/core/address"
	"cosmossdk.io/core/log"

	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

type (
	// Registry is an abstract entry point to register message factories with weights
	Registry interface {
		Add(weight uint32, f SimMsgFactoryX)
	}
	// FutureOpsRegistry register message factories for future blocks
	FutureOpsRegistry interface {
		Add(blockTime time.Time, f SimMsgFactoryX)
	}

	// AccountSourceX Account and Module account
	AccountSourceX interface {
		AccountSource
		ModuleAccountSource
	}
)

// WeightedProposalMsgIter iterator for weighted gov proposal payload messages
type WeightedProposalMsgIter = iter.Seq2[uint32, SimMsgFactoryX]

var _ Registry = &WeightedOperationRegistryAdapter{}

// common types for abstract registry without generics
type regCommon struct {
	reporter     SimulationReporter
	ak           AccountSourceX
	bk           BalanceSource
	addressCodec address.Codec
	txConfig     client.TxConfig
	logger       log.Logger
}

func (c regCommon) newChainDataSource(ctx context.Context, r *rand.Rand, accs ...simtypes.Account) *ChainDataSource {
	return NewChainDataSource(ctx, r, c.ak, c.bk, c.addressCodec, accs...)
}

type AbstractRegistry[T any] struct {
	regCommon
	legacyObjs []T
}

// ToLegacyObjects returns the legacy properties of the SimsRegistryAdapter as a slice of type T.
func (l *AbstractRegistry[T]) ToLegacyObjects() []T {
	return l.legacyObjs
}

// WeightedOperationRegistryAdapter is an implementation of the Registry interface that provides adapters to use the new message factories
// with the legacy simulation system
type WeightedOperationRegistryAdapter struct {
	AbstractRegistry[simtypes.WeightedOperation]
}

// NewSimsMsgRegistryAdapter creates a new instance of SimsRegistryAdapter for WeightedOperation types.
func NewSimsMsgRegistryAdapter(
	reporter SimulationReporter,
	ak AccountSourceX,
	bk BalanceSource,
	txConfig client.TxConfig,
	logger log.Logger,
) *WeightedOperationRegistryAdapter {
	return &WeightedOperationRegistryAdapter{
		AbstractRegistry: AbstractRegistry[simtypes.WeightedOperation]{
			regCommon: regCommon{
				reporter:     reporter,
				ak:           ak,
				bk:           bk,
				txConfig:     txConfig,
				addressCodec: txConfig.SigningContext().AddressCodec(),
				logger:       logger,
			},
		},
	}
}

// Add adds a new weighted operation to the collection
func (l *WeightedOperationRegistryAdapter) Add(weight uint32, fx SimMsgFactoryX) {
	if fx == nil {
		panic("message factory must not be nil")
	}
	if weight == 0 {
		return
	}
	obj := simulation.NewWeightedOperation(int(weight), legacyOperationAdapter(l.regCommon, fx))
	l.legacyObjs = append(l.legacyObjs, obj)
}

type HasFutureOpsRegistry interface {
	SetFutureOpsRegistry(FutureOpsRegistry)
}

func legacyOperationAdapter(l regCommon, fx SimMsgFactoryX) simtypes.Operation {
	return func(
		r *rand.Rand, app AppEntrypoint, ctx sdk.Context,
		accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		xCtx, done := context.WithCancel(ctx)
		ctx = sdk.UnwrapSDKContext(xCtx)
		testData := l.newChainDataSource(ctx, r, accs...)
		reporter := l.reporter.WithScope(fx.MsgType(), SkipHookFn(func(args ...any) { done() }))
		fOpsReg := NewFutureOpsRegistry(l)
		if fx, ok := fx.(HasFutureOpsRegistry); ok {
			fx.SetFutureOpsRegistry(fOpsReg)
		}
		from, msg := SafeRunFactoryMethod(ctx, testData, reporter, fx.Create())
		futOps := fOpsReg.legacyObjs
		weightedOpsResult := DeliverSimsMsg(ctx, reporter, app, r, l.txConfig, l.ak, chainID, msg, fx.DeliveryResultHandler(), from...)
		err := reporter.Close()
		return weightedOpsResult, futOps, err
	}
}

func NewFutureOpsRegistry(l regCommon) *FutureOperationRegistryAdapter {
	return &FutureOperationRegistryAdapter{regCommon: l}
}

type FutureOperationRegistryAdapter AbstractRegistry[simtypes.FutureOperation]

func (l *FutureOperationRegistryAdapter) Add(blockTime time.Time, fx SimMsgFactoryX) {
	if fx == nil {
		panic("message factory must not be nil")
	}
	if blockTime.IsZero() {
		return
	}
	obj := simtypes.FutureOperation{
		BlockTime: blockTime,
		Op:        legacyOperationAdapter(l.regCommon, fx),
	}
	l.legacyObjs = append(l.legacyObjs, obj)
}

var _ Registry = &UniqueTypeRegistry{}

type UniqueTypeRegistry map[string]WeightedFactory

func NewUniqueTypeRegistry() UniqueTypeRegistry {
	return make(UniqueTypeRegistry)
}

func (s UniqueTypeRegistry) Add(weight uint32, f SimMsgFactoryX) {
	msgType := f.MsgType()
	msgTypeURL := sdk.MsgTypeURL(msgType)
	if _, exists := s[msgTypeURL]; exists {
		panic("type is already registered: " + msgTypeURL)
	}
	s[msgTypeURL] = WeightedFactory{Weight: weight, Factory: f}
}

// Iterator returns an iterator function for a Go for loop sorted by weight desc.
func (s UniqueTypeRegistry) Iterator() iter.Seq2[uint32, SimMsgFactoryX] {
	x := maps.Values(s)
	slices.SortFunc(x, func(a, b WeightedFactory) int {
		return a.Compare(b)
	})
	return func(yield func(uint32, SimMsgFactoryX) bool) {
		for _, v := range x {
			if !yield(v.Weight, v.Factory) {
				return
			}
		}
	}
}

// WeightedFactory is a Weight Factory tuple
type WeightedFactory struct {
	Weight  uint32
	Factory SimMsgFactoryX
}

// Compare compares the WeightedFactory f with another WeightedFactory b.
// The comparison is done by comparing the weight of f with the weight of b.
// If the weight of f is greater than the weight of b, it returns 1.
// If the weight of f is less than the weight of b, it returns -1.
// If the weights are equal, it compares the TypeURL of the factories using strings.Compare.
// Returns an integer indicating the comparison result.
func (f WeightedFactory) Compare(b WeightedFactory) int {
	switch {
	case f.Weight > b.Weight:
		return 1
	case f.Weight < b.Weight:
		return -1
	default:
		return strings.Compare(sdk.MsgTypeURL(f.Factory.MsgType()), sdk.MsgTypeURL(b.Factory.MsgType()))
	}
}
