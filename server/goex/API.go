package goex

// api interface

type API interface {
	LimitBuy(amount, price string, currency CurrencyPair, opt ...LimitOrderOptionalParameter) (*Order, error)
	LimitSell(amount, price string, currency CurrencyPair, opt ...LimitOrderOptionalParameter) (*Order, error)
	MarketBuy(amount, posSide string, currency CurrencyPair) (*Order, error)
	MarketSell(amount, price string, currency CurrencyPair) (*Order, error)
	CancelOrder(orderId string, currency CurrencyPair) (bool, error)
	GetOneOrder(orderId string, currency CurrencyPair) (*Order, error)
	GetUnfinishOrders(currency CurrencyPair) ([]Order, error)
	GetOrderHistorys(currency CurrencyPair, opt ...OptionalParameter) ([]Order, error)
	GetAccount() (*Account, error)
	ClosePosition(instId, mode, ccy string) (bool, error)
	GetTicker(currency CurrencyPair) (*Ticker, error)
	GetDepth(size int, currency CurrencyPair) (*Depth, error)
	GetKlineRecords(currency CurrencyPair, period KlinePeriod, size int, before, after int) ([]Kline, error)

	GetPositions(instId string) (*Positions, error)
	SetLeverage(instId, ccy string, lever string, mode string) (bool, error)
	GetMaxSize(instId, mode string) (*Max, error)
	GetExchangeName() string
	GetCoinInfo(currency CurrencyPair, instType string) (*Coin, error)
}
