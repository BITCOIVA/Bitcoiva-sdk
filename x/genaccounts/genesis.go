package genaccounts

import (
	"github.com/BITCOIVA/Bitcoiva-sdk/codec"
	sdk "github.com/BITCOIVA/Bitcoiva-sdk/types"
	authexported "github.com/BITCOIVA/Bitcoiva-sdk/x/auth/exported"
	"github.com/BITCOIVA/Bitcoiva-sdk/x/genaccounts/internal/types"
)

// InitGenesis initializes accounts and deliver genesis transactions
func InitGenesis(ctx sdk.Context, _ *codec.Codec, accountKeeper types.AccountKeeper, genesisState GenesisState) {
	genesisState.Sanitize()

	// load the accounts
	for _, gacc := range genesisState {
		acc := gacc.ToAccount()
		acc = accountKeeper.NewAccount(ctx, acc) // set account number
		accountKeeper.SetAccount(ctx, acc)
	}
}

// ExportGenesis exports genesis for all accounts
func ExportGenesis(ctx sdk.Context, accountKeeper types.AccountKeeper) GenesisState {

	// iterate to get the accounts
	accounts := []GenesisAccount{}
	accountKeeper.IterateAccounts(ctx,
		func(acc authexported.Account) (stop bool) {
			account, err := NewGenesisAccountI(acc)
			if err != nil {
				panic(err)
			}
			accounts = append(accounts, account)
			return false
		},
	)

	return accounts
}
