package trading212

import (
	"encoding/json"
	"io"
	"net/http"
)

type Client struct {
	apiKey  string
	baseUrl string
}

func NewClient(apiKey string, demo bool) *Client {
	if demo {
		return &Client{apiKey, baseUrlDemo}
	}
	return &Client{apiKey, baseUrlLive}
}

func (c *Client) getRequest(reqUrl string, dest any) error {
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
