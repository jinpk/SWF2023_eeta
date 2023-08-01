package types

import (
	"fmt"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

var _ paramtypes.ParamSet = (*Params)(nil)

var (
	KeyWaitBlock = []byte("WaitBlock")
	// TODO: Determine the default value
	DefaultWaitBlock uint64 = 0
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(
	waitBlock uint64,
) Params {
	return Params{
		WaitBlock: waitBlock,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(
		DefaultWaitBlock,
	)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyWaitBlock, &p.WaitBlock, validateWaitBlock),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validateWaitBlock(p.WaitBlock); err != nil {
		return err
	}

	return nil
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

// validateWaitBlock validates the WaitBlock param
func validateWaitBlock(v interface{}) error {
	waitBlock, ok := v.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	// TODO implement validation
	_ = waitBlock

	return nil
}
