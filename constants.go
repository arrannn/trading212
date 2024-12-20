package trading212

const (
	baseUrlDemo = "https://demo.trading212.com/api/v0"
	baseUrlLive = "https://live.trading212.com/api/v0/"
)

const (
	// Account endpoints
	EndpointGetAccountCash     = "/equity/account/cash"
	EndpointGetAccountMetadata = "/equity/account/info"

	// Market data endpoints
	EndpointGetExchanges   = "/equity/metadata/exchanges"
	EndpointGetInstruments = "/equity/metadata/instruments"

	// Order endpoints
	EndpointPlaceMarketOrder    = "/equity/orders/market"
	EndpointPlaceLimitOrder     = "/equity/orders/limit"
	EndpointPlaceStopOrder      = "/equity/orders/stop"
	EndpointPlaceStopLimitOrder = "/equity/orders/stop_limit"
	EndpointGetOrders           = "/equity/orders"
)
