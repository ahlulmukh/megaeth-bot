package ethclient

import (
	"context"
	"net/http"
	"net/url"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

type Client struct {
	*ethclient.Client
	rpcClient *rpc.Client
}

func New(ctx context.Context, rpcURL string, proxyURL string) (*Client, error) {
	var httpClient *http.Client

	if proxyURL != "" {
		proxy, err := url.Parse(proxyURL)
		if err != nil {
			return nil, err
		}

		transport := &http.Transport{
			Proxy: http.ProxyURL(proxy),
		}

		httpClient = &http.Client{
			Transport: transport,
		}
	} else {
		httpClient = http.DefaultClient
	}

	rpcClient, err := rpc.DialOptions(ctx, rpcURL, rpc.WithHTTPClient(httpClient))
	if err != nil {
		return nil, err
	}

	client := ethclient.NewClient(rpcClient)

	return &Client{
		Client:    client,
		rpcClient: rpcClient,
	}, nil
}

func (c *Client) Close() {
	c.rpcClient.Close()
}
