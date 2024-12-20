package trading212

import "time"

type Exchange struct {
	ID               int               `json:"id"`
	Name             string            `json:"name"`
	WorkingSchedules []WorkingSchedule `json:"workingSchedules"`
}

type WorkingSchedule struct {
	ID         int         `json:"id"`
	TimeEvents []TimeEvent `json:"timeEvents"`
}

type TimeEvent struct {
	Date time.Time `json:"date"`
	Type string    `json:"type"`
}

func (c *Trading212Client) GetExchanges() ([]*Exchange, error) {
	var v []*Exchange
	err := c.getRequest(EndpointGetExchanges, &v)
	return v, err
}

type Instrument struct {
	AddedOn           time.Time `json:"addedOn"`
	CurrencyCode      string    `json:"currencyCode"`
	Isin              string    `json:"isin"`
	MaxOpenQuantity   int       `json:"maxOpenQuantity"`
	MinTradeQuantity  int       `json:"minTradeQuantity"`
	Name              string    `json:"name"`
	ShortName         string    `json:"shortName"`
	Ticker            string    `json:"ticker"`
	Type              string    `json:"type"`
	WorkingScheduleID int       `json:"workingScheduleId"`
}

func (c *Trading212Client) GetInstruments() ([]*Instrument, error) {
	var v []*Instrument
	err := c.getRequest(EndpointGetInstruments, &v)
	return v, err
}
