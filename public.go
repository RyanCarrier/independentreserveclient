package independentreserveclient

import "github.com/RyanCarrier/cryptoclientgo"

/*
Tick(CurrencyFrom, CurrencyTo string) (cryptoclientgo.Tick, error)
GetOrderBook(CurrencyFrom, CurrencyTo string) (cryptoclientgo.OrderBook, error)
GetRecentTrades(CurrencyFrom, CurrencyTo string, historyAmount int) (cryptoclientgo.RecentTrades, error)
ExpectedMarketValueBuy(PrimaryCurrency, SecondaryCurrency string, amountOfToCurrency int64) (int64, error)
ExpectedMarketValueSell(PrimaryCurrency, SecondaryCurrency string, amountOfFromCurrency int64) (int64, error)
*/

func (c Client) Tick(CurrencyFrom, CurrencyTo string) (cryptoclientgo.Tick, error) {

}
