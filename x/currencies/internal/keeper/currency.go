package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/dfinance/dnode/x/ccstorage"
)

// CreateCurrency redirects CreateCurrency request to the currencies storage.
func (k Keeper) CreateCurrency(ctx sdk.Context, denom string, params ccstorage.CurrencyParams) error {
	return k.ccsKeeper.CreateCurrency(ctx, denom, params)
}