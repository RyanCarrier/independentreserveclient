package independentreservego

GetPrimaryCurrencies() ([]string, error)
GetSecondaryCurrencies() ([]string, error)
GetTransactionCost(CurrencyFrom, CurrencyTo string) (cryptoclientgo.Cost, error)
GetWithdrawCost(Currency string) (cryptoclientgo.Cost, error)
GetDepositCost(Currency string) (cryptoclientgo.Cost, error)
