package independentreservego

PlaceLimitBuyOrder(CurrencyFrom, CurrencyTo string, amount int64, price int64) (cryptoclientgo.PlacedOrder, error)
PlaceMarketBuyOrder(CurrencyFrom, CurrencyTo string, amount int64) (cryptoclientgo.PlacedOrder, error)
PlaceLimitSellOrder(CurrencyFrom, CurrencyTo string, amount int64, price int64) (cryptoclientgo.PlacedOrder, error)
PlaceMarketSellOrder(CurrencyFrom, CurrencyTo string, amount int64) (cryptoclientgo.PlacedOrder, error)
CancelOrder(OrderID int) error
GetOrderDetails(OrderID int) (cryptoclientgo.OrderDetails, error)
GetOpenOrders() (cryptoclientgo.OrdersDetails, error)
