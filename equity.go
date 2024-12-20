package trading212

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type TimeValidity string

const (
	TimeValidityDay            TimeValidity = "DAY"
	TimeValidityGoodTillCancel TimeValidity = "GOOD_TILL_CANCEL"
)

type MarketOrder struct {
	Quantity float64 `json:"quantity"`
	Ticker   string  `json:"ticker"`
}

func (c *Client) PlaceMarketOrder(order MarketOrder) (*EquityOrderResponse, error) {
	return c.placeEquityOrder(EndpointPlaceMarketOrder, order)
}

type LimitOrder struct {
	LimitPrice   float64      `json:"limitPrice"`
	Quantity     float64      `json:"quantity"`
	Ticker       string       `json:"ticker"`
	TimeValidity TimeValidity `json:"timeValidity"`
}

func (c *Client) PlaceLimitOrder(order LimitOrder) (*EquityOrderResponse, error) {
	return c.placeEquityOrder(EndpointPlaceLimitOrder, order)
}

type StopOrder struct {
	Quantity     float64 `json:"quantity"`
	StopPrice    float64 `json:"stopPrice"`
	Ticker       string  `json:"ticker"`
	TimeValidity string  `json:"timeValidity"`
}

func (c *Client) PlaceStopOrder(order StopOrder) (*EquityOrderResponse, error) {
	return c.placeEquityOrder(EndpointPlaceStopOrder, order)
}

type StopLimitOrder struct {
	LimitPrice   float64 `json:"limitPrice"`
	Quantity     float64 `json:"quantity"`
	StopPrice    float64 `json:"stopPrice"`
	Ticker       string  `json:"ticker"`
	TimeValidity string  `json:"timeValidity"`
}

func (c *Client) PlaceStopLimitOrder(order StopLimitOrder) (*EquityOrderResponse, error) {
	return c.placeEquityOrder(EndpointPlaceStopLimitOrder, order)
}

type EquityOrderResponse struct {
	CreationTime   time.Time `json:"creationTime"`
	FilledQuantity int       `json:"filledQuantity"`
	FilledValue    int       `json:"filledValue"`
	ID             int       `json:"id"`
	LimitPrice     int       `json:"limitPrice"`
	Quantity       int       `json:"quantity"`
	Status         string    `json:"status"`
	StopPrice      int       `json:"stopPrice"`
	Strategy       string    `json:"strategy"`
	Ticker         string    `json:"ticker"`
	Type           string    `json:"type"`
	Value          int       `json:"value"`
}

func (c *Client) GetOrder(orderID int) (*EquityOrderResponse, error) {
	var v EquityOrderResponse
	if err := c.getRequest(fmt.Sprintf("%s/%d", EndpointGetOrders, orderID), &v); err != nil {
		return nil, err
	}
	return &v, nil
}

func (c *Client) GetOrders() ([]*EquityOrderResponse, error) {
	var v []*EquityOrderResponse
	if err := c.getRequest(EndpointGetOrders, &v); err != nil {
		return nil, err
	}
	return v, nil
}

func (c *Client) placeEquityOrder(reqUrl string, order any) (*EquityOrderResponse, error) {
	data, err := json.Marshal(order)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", c.baseUrl+reqUrl, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", c.apiKey)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode == http.StatusBadRequest {
		var validationError ValidationError
		err := json.NewDecoder(res.Body).Decode(&validationError)
		if err != nil {
			return nil, err
		}
		return nil, &validationError
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var v EquityOrderResponse
	err = json.Unmarshal(body, &v)
	if err != nil {
		return nil, err
	}
	return &v, nil
}
