package usecase

import "fmt"

type WalletFacade struct {
	Account      *Account
	Wallet       *Wallet
	SecurityCode *SecurityCode
	Notification *Notification
	Ledger       *Ledger
}

func NewWalletFacade(accountID string, code int) *WalletFacade {
	fmt.Println("Starting create account")
	walletFacacde := &WalletFacade{
		Account:      NewAccount(accountID),
		SecurityCode: NewSecurityCode(code),
		Wallet:       NewWallet(),
		Notification: &Notification{},
		Ledger:       &Ledger{},
	}
	fmt.Println("Account created")
	return walletFacacde
}

func (w *WalletFacade) AddMoneyToWallet(accountID string, securityCode int, amount int) error {
	fmt.Println("Starting add money to wallet")
	err := w.Account.CheckAccount(accountID)
	if err != nil {
		return err
	}
	err = w.SecurityCode.CheckCode(securityCode)
	if err != nil {
		return err
	}
	w.Wallet.CreditBalance(amount)
	w.Notification.SendWalletCreditNotification()
	w.Ledger.MakeEntry(accountID, "credit", amount)
	return nil
}

func (w *WalletFacade) DeductMoneyFromWallet(accountID string, securityCode int, amount int) error {
	fmt.Println("Starting debit money from wallet")
	err := w.Account.CheckAccount(accountID)
	if err != nil {
		return err
	}

	err = w.SecurityCode.CheckCode(securityCode)
	if err != nil {
		return err
	}
	err = w.Wallet.DebitBalance(amount)
	if err != nil {
		return err
	}
	w.Notification.SendWalletDebitNotification()
	w.Ledger.MakeEntry(accountID, "credit", amount)
	return nil
}
