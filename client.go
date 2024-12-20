package trading212

import (
	"encoding/json"
	"io"
	"net/http"
)

type Trading212Client struct {
	apiKey  string
	baseUrl string
}

func NewTrading212Client(apiKey string, demo bool) *Trading212Client {
	if demo {
		return &Trading212Client{apiKey, baseUrlDemo}
	}
	return &Trading212Client{apiKey, baseUrlLive}
}

func (c *Trading212Client) getRequest(reqUrl string, dest any) error {
	req, err := http.NewRequest("GET", c.baseUrl+reqUrl, nil)
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", c.apiKey)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, dest)
}
