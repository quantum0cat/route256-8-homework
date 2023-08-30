package loms

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type Client struct {
	name string
	path string
}

type StocksRequest struct {
	SKU uint32 `json:"sku,omitempty"`
}

type StocksResponse struct {
	Count uint64 `json:"count,omitempty"`
}

const stocksPath = "stocks"

func New(name string, basePath string) (*Client, error) {
	path, err := url.JoinPath(basePath, stocksPath)
	if err != nil {
		return nil, fmt.Errorf("%s: incorrect base path: %w", name, err)
	}
	return &Client{
		name: name,
		path: path,
	}, nil
}

func (c Client) GetStocks(ctx context.Context, sku uint32) (uint64, error) {
	request := StocksRequest{
		SKU: sku,
	}
	data, err := json.Marshal(request)
	if err != nil {
		return 0, fmt.Errorf("%s: failed to encode request %w", c.name, err)
	}
	httpRequest, err := http.NewRequestWithContext(ctx, http.MethodPost, c.path, bytes.NewBuffer(data))
	if err != nil {
		return 0, fmt.Errorf("%s: failed to create HTTP request: %w", c.name, err)
	}
	httpResponse, err := http.DefaultClient.Do(httpRequest)
	if err != nil {
		return 0, fmt.Errorf("%s: failed to execute HTTP request: %w", c.name, err)
	}
	defer func() {
		_ = httpResponse.Body.Close()
	}()
	if httpResponse.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("%s: HTTP request responded with: %d", c.name, httpResponse.StatusCode)
	}
	response := &StocksResponse{}
	err = json.NewDecoder(httpResponse.Body).Decode(response)
	if err != nil {
		return 0, fmt.Errorf("%s: failed to encode response %w", c.name, err)
	}
	return response.Count, nil
}
