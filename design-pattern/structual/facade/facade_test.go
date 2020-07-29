package facade

import "testing"

func Test_newWalletFacade(t *testing.T) {
	walletFacade := newWalletFacade("betty", 123)
	_ = walletFacade.addMoneyToWallet(walletFacade.account.id, 123, 100)
	_ = walletFacade.deductMoneyFromWallet(walletFacade.account.id, 123, 100)
}
