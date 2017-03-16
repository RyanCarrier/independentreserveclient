package independentreserveclient

import "encoding/json"

/*
GetPrimaryCurrencies() ([]string, error)
GetSecondaryCurrencies() ([]string, error)
GetTransactionCost(CurrencyFrom, CurrencyTo string) (cryptoclientgo.Cost, error)
GetWithdrawCost(Currency string) (cryptoclientgo.Cost, error)
GetDepositCost(Currency string) (cryptoclientgo.Cost, error)
*/

//GetPrimaryCurrencies gets the primary currencies available
func GetPrimaryCurrencies() ([]string, error) {
	got, err := getAndRead(Domain + "/Public/GetValidPrimaryCurrencyCodes")
	if err != nil {
		return []string{}, err
	}
	var l []string
	err = json.Unmarshal(got, l)
	return l, err
}
