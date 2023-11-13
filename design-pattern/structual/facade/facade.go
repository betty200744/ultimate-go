package facade

import (
	"fmt"
	"ultimate-go/design-pattern/creational/simple_factory"
)

type walletFacade struct {
	account      *account
	wallet       *wallet
	password     *password
	notification *notification
}

func newWalletFacade(name string, code int) *walletFacade {
	fmt.Println("Starting create account")
	walletFacacde := &walletFacade{
		account:      newAccount(name),
		password:     newPassword(code),
		wallet:       newWallet(),
		notification: &notification{},
	}
	fmt.Println("Account created")
	return walletFacacde
}

func (w *walletFacade) addMoneyToWallet(accountID int64, password int, amount int) error {
	fmt.Println("Starting add money to wallet")
	err := w.account.checkAccount(accountID)
	if err != nil {
		return err
	}
	err = w.password.checkPassword(password)
	if err != nil {
		return err
	}
	w.wallet.creditBalance(amount)
	w.notification.sendWalletCreditNotification()
	return nil
}

func (w *walletFacade) deductMoneyFromWallet(accountID int64, password int, amount int) error {
	fmt.Println("Starting debit money from wallet")
	err := w.account.checkAccount(accountID)
	if err != nil {
		return err
	}
	err = w.password.checkPassword(password)
	if err != nil {
		return err
	}
	err = w.wallet.debitBalance(amount)
	if err != nil {
		return err
	}
	w.notification.sendWalletDebitNotification()
	return nil
}

type account struct {
	id   int64
	name string
}

func newAccount(name string) *account {
	s, _ := simple_factory.NewIdGenerator()
	id := s.Generate()
	return &account{
		id:   id,
		name: "",
	}
}
func (a *account) checkAccount(id int64) error {
	if a.id != id {
		return fmt.Errorf("Account id is incorrect")
	}
	fmt.Println("Account Verified")
	return nil
}

type password struct {
	password int
}

func newPassword(pwd int) *password {
	return &password{password: pwd}
}
func (p *password) checkPassword(pwd int) error {
	if p.password != pwd {
		return fmt.Errorf("password name is incorrect")
	}
	fmt.Println("password Verified")
	return nil
}

type wallet struct {
	balance int
}

func newWallet() *wallet {
	return &wallet{balance: 0}
}
func (w *wallet) creditBalance(amount int) {
	w.balance += amount
	fmt.Println("Wallet balance added successfully")
	return
}

func (w *wallet) debitBalance(amount int) error {
	if w.balance < amount {
		return fmt.Errorf("Balance is not sufficient")
	}
	fmt.Println("Wallet balance is Sufficient")
	w.balance = w.balance - amount
	return nil
}

type notification struct {
}

func (n *notification) sendWalletCreditNotification() {
	fmt.Println("Sending wallet credit notification")
}

func (n *notification) sendWalletDebitNotification() {
	fmt.Println("Sending wallet debit notification")
}
