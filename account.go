package trading212

type AccountCash struct {
	Blocked  float32 `json:"blocked"`
	Free     float32 `json:"free"`
	Invested float32 `json:"invested"`
	PieCash  float32 `json:"pieCash"`
	PPL      float32 `json:"ppl"`
	Result   float32 `json:"result"`
	Total    float32 `json:"total"`
}

func (c *Trading212Client) GetAccountCash() (*AccountCash, error) {
	var v AccountCash
	err := c.getRequest(EndpointGetAccountCash, &v)
	return &v, err
}

type Account struct {
	CurrencyCode string `json:"currencyCode"`
	ID           int    `json:"id"`
}

func (c *Trading212Client) GetAccount() (*Account, error) {
	var v Account
	err := c.getRequest(EndpointGetAccountMetadata, &v)
	return &v, err
}
