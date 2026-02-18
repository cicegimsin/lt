package aur

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

const aurRPCURL = "https://aur.archlinux.org/rpc"

type Client struct {
	httpClient *http.Client
}

type RPCResponse struct {
	ResultCount int         `json:"resultcount"`
	Results     []AURPackage `json:"results"`
	Type        string      `json:"type"`
}

type AURPackage struct {
	Name        string  `json:"Name"`
	Version     string  `json:"Version"`
	Description string  `json:"Description"`
	URL         string  `json:"URL"`
	Popularity  float64 `json:"Popularity"`
	NumVotes    int     `json:"NumVotes"`
	Maintainer  string  `json:"Maintainer"`
	LastModified int64  `json:"LastModified"`
}

func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (c *Client) Search(query string) ([]AURPackage, error) {
	params := url.Values{}
	params.Set("v", "5")
	params.Set("type", "search")
	params.Set("arg", query)
	
	resp, err := c.httpClient.Get(aurRPCURL + "?" + params.Encode())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("AUR API hatası: %d", resp.StatusCode)
	}
	
	var rpcResp RPCResponse
	if err := json.NewDecoder(resp.Body).Decode(&rpcResp); err != nil {
		return nil, err
	}
	
	return rpcResp.Results, nil
}

func (c *Client) Info(pkgName string) (*AURPackage, error) {
	params := url.Values{}
	params.Set("v", "5")
	params.Set("type", "info")
	params.Set("arg", pkgName)
	
	resp, err := c.httpClient.Get(aurRPCURL + "?" + params.Encode())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	
	var rpcResp RPCResponse
	if err := json.NewDecoder(resp.Body).Decode(&rpcResp); err != nil {
		return nil, err
	}
	
	if len(rpcResp.Results) == 0 {
		return nil, fmt.Errorf("paket bulunamadı")
	}
	
	return &rpcResp.Results[0], nil
}
