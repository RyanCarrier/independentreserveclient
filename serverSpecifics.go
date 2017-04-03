package independentreserveclient

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/RyanCarrier/cryptoclientgo"
)

//GetWithdrawCost gets the cost of withdrawing the specific currency
func (c Client) GetWithdrawCost(Currency string) (cryptoclientgo.Cost, error) {
	if lookup(c.secondaryCurrencies, Currency) >= 0 {
		if Currency == "AUD" {
			return cryptoclientgo.Cost{Flat: 0, Percent: 0}, nil
		}
		return cryptoclientgo.Cost{Flat: 20 * Multiplier, Percent: 0}, nil
	}
	switch strings.ToUpper(Currency) {
	case "XBT":
		return cryptoclientgo.Cost{Flat: int64(0.0005 * float64(Multiplier)), Percent: 0}, nil
	case "ETH":
		return cryptoclientgo.Cost{Flat: int64(0.01 * float64(Multiplier)), Percent: 0}, nil
	default:
		return cryptoclientgo.Cost{}, errors.New("Currency " + Currency + " not found")
	}
}

//GetDepositCost gets the cost of deposit the specific currency
func (c Client) GetDepositCost(Currency string) (cryptoclientgo.Cost, error) {
	if lookup(c.primaryCurrencies, Currency) >= 0 {
		return cryptoclientgo.Cost{Flat: 0, Percent: 0}, nil
	}
	return cryptoclientgo.Cost{Flat: 1 * Multiplier, Percent: 0}, nil
}

func (c Client) getCurrencies(primary bool) ([]string, error) {
	var URL string
	if primary {
		URL = Domain + "/Public/GetValidPrimaryCurrencyCodes"
	} else {
		URL = Domain + "/Public/GetValidSecondaryCurrencyCodes"
	}
	got, err := getAndRead(URL)
	if err != nil {
		return []string{}, err
	}
	var l []string
	err = json.Unmarshal(got, l)
	return l, err
}

//GetPrimaryCurrencies gets the primary currencies available
func (c Client) GetPrimaryCurrencies() ([]string, error) {
	return c.primaryCurrencies, nil
}

//GetSecondaryCurrencies gets the Secondary currencies available
func (c Client) GetSecondaryCurrencies() ([]string, error) {
	return c.secondaryCurrencies, nil
}

//GetTransactionCost gets the transaction cost between two currencies
// if trading between currencies is not allowed, the cost will be 100%
func (c Client) GetTransactionCost(CurrencyFrom, CurrencyTo string) (cryptoclientgo.Cost, error) {
	if (lookup(c.primaryCurrencies, CurrencyFrom) >= 0 && lookup(c.secondaryCurrencies, CurrencyTo) >= 0) ||
		(lookup(c.primaryCurrencies, CurrencyTo) >= 0 && lookup(c.secondaryCurrencies, CurrencyFrom) >= 0) {
		//100%=1;10%=0.1;1%=0.01;0.5%=0.005
		return cryptoclientgo.Cost{Flat: 0, Percent: int64(0.005 * float64(Multiplier))}, nil
	}
	return cryptoclientgo.Cost{Flat: 0, Percent: 1 * Multiplier}, nil
}
