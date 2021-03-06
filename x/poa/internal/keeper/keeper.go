// PoA module keeper stores validators meta data and multi signature params.
package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/dfinance/dnode/helpers/perms"
	"github.com/dfinance/dnode/x/poa/internal/types"
)

// Module keeper object.
type Keeper struct {
	cdc         *codec.Codec
	storeKey    sdk.StoreKey
	paramStore  params.Subspace
	modulePerms perms.ModulePermissions
}

// GetLogger gets logger with keeper context.
func (k Keeper) GetLogger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", "x/"+types.ModuleName)
}

// Create new currency keeper.
func NewKeeper(
	cdc *codec.Codec,
	storeKey sdk.StoreKey,
	paramStore params.Subspace,
	permsRequesters ...perms.RequestModulePermissions,
) Keeper {
	k := Keeper{
		cdc:         cdc,
		storeKey:    storeKey,
		paramStore:  paramStore.WithKeyTable(types.ParamKeyTable()),
		modulePerms: types.NewModulePerms(),
	}
	for _, requester := range permsRequesters {
		k.modulePerms.AutoAddRequester(requester)
	}

	return k
}
