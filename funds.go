package independentreserveclient

	GetBalance(Currency string) (cryptoclientgo.AccountBalance, error)
	GetBalances() (cryptoclientgo.AccountBalances, error)
	GetPrimaryCurrencyDepositAddress(Currency string) (cryptoclientgo.CurrencyAddress, error)
	WithdrawCurrency(Currency, to string, amount int64) error
