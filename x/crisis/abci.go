package crisis

import (
	sdk "github.com/BITCOIVA/Bitcoiva-sdk/types"
)

// check all registered invariants
func EndBlocker(ctx sdk.Context, k Keeper) {
	if k.InvCheckPeriod() == 0 || ctx.BlockHeight()%int64(k.InvCheckPeriod()) != 0 {
		// skip running the invariant check
		return
	}
	k.AssertInvariants(ctx)
}
