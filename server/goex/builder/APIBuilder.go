package builder

import (
	"context"
	"errors"
	"fmt"

	. "goex"
	//"github.com/nntaoli-project/goex/coin58"
	"net"
	"net/http"
	"net/url"
	"time"

	"goex/okex"
	okexV5 "goex/okex/v5"
)

type APIBuilder struct {
	HttpClientConfig *HttpClientConfig
	client           *http.Client
	httpTimeout      time.Duration
	apiKey           string
	secretkey        string
	clientId         string
	apiPassphrase    string
	futuresEndPoint  string
	endPoint         string
	futuresLever     float64
}

type HttpClientConfig struct {
	HttpTimeout  time.Duration
	Proxy        *url.URL
	MaxIdleConns int
}

func (c HttpClientConfig) String() string {
	return fmt.Sprintf("{ProxyUrl:\"%s\",HttpTimeout:%s,MaxIdleConns:%d}", c.Proxy, c.HttpTimeout.String(), c.MaxIdleConns)
}

func (c *HttpClientConfig) SetHttpTimeout(timeout time.Duration) *HttpClientConfig {
	c.HttpTimeout = timeout
	return c
}

func (c *HttpClientConfig) SetProxyUrl(proxyUrl string) *HttpClientConfig {
	if proxyUrl == "" {
		return c
	}
	proxy, err := url.Parse(proxyUrl)
	if err != nil {
		return c
	}
	c.Proxy = proxy
	return c
}

func (c *HttpClientConfig) SetMaxIdleConns(max int) *HttpClientConfig {
	c.MaxIdleConns = max
	return c
}

var (
	DefaultHttpClientConfig = &HttpClientConfig{
		Proxy:        nil,
		HttpTimeout:  5 * time.Second,
		MaxIdleConns: 10}
	DefaultAPIBuilder = NewAPIBuilder()
)

func NewAPIBuilder() (builder *APIBuilder) {
	return NewAPIBuilder2(DefaultHttpClientConfig)
}

func NewAPIBuilder2(config *HttpClientConfig) *APIBuilder {
	if config == nil {
		config = DefaultHttpClientConfig
	}

	return &APIBuilder{
		HttpClientConfig: config,
		client: &http.Client{
			Timeout: config.HttpTimeout,
			Transport: &http.Transport{
				Proxy: func(request *http.Request) (*url.URL, error) {
					return config.Proxy, nil
				},
				MaxIdleConns:          config.MaxIdleConns,
				IdleConnTimeout:       5 * config.HttpTimeout,
				MaxConnsPerHost:       2,
				MaxIdleConnsPerHost:   2,
				TLSHandshakeTimeout:   config.HttpTimeout,
				ResponseHeaderTimeout: config.HttpTimeout,
				ExpectContinueTimeout: config.HttpTimeout,
				DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
					return net.DialTimeout(network, addr, config.HttpTimeout)
				}},
		}}
}

func NewCustomAPIBuilder(client *http.Client) (builder *APIBuilder) {
	return &APIBuilder{client: client}
}

func (builder *APIBuilder) GetHttpClientConfig() *HttpClientConfig {
	return builder.HttpClientConfig
}

func (builder *APIBuilder) GetHttpClient() *http.Client {
	return builder.client
}

func (builder *APIBuilder) HttpProxy(proxyUrl string) (_builder *APIBuilder) {
	if proxyUrl == "" {
		return builder
	}
	proxy, err := url.Parse(proxyUrl)
	if err != nil {
		return builder
	}
	builder.HttpClientConfig.Proxy = proxy
	transport := builder.client.Transport.(*http.Transport)
	transport.Proxy = http.ProxyURL(proxy)
	return builder
}

func (builder *APIBuilder) HttpTimeout(timeout time.Duration) (_builder *APIBuilder) {
	builder.HttpClientConfig.HttpTimeout = timeout
	builder.httpTimeout = timeout
	builder.client.Timeout = timeout
	transport := builder.client.Transport.(*http.Transport)
	if transport != nil {
		//transport.ResponseHeaderTimeout = timeout
		//transport.TLSHandshakeTimeout = timeout
		transport.IdleConnTimeout = timeout
		transport.DialContext = func(ctx context.Context, network, addr string) (net.Conn, error) {
			return net.DialTimeout(network, addr, timeout)
		}
	}
	return builder
}

func (builder *APIBuilder) APIKey(key string) (_builder *APIBuilder) {
	builder.apiKey = key
	return builder
}

func (builder *APIBuilder) APISecretkey(key string) (_builder *APIBuilder) {
	builder.secretkey = key
	return builder
}

func (builder *APIBuilder) ClientID(id string) (_builder *APIBuilder) {
	builder.clientId = id
	return builder
}

func (builder *APIBuilder) ApiPassphrase(apiPassphrase string) (_builder *APIBuilder) {
	builder.apiPassphrase = apiPassphrase
	return builder
}

func (builder *APIBuilder) FuturesEndpoint(endpoint string) (_builder *APIBuilder) {
	builder.futuresEndPoint = endpoint
	return builder
}

func (builder *APIBuilder) Endpoint(endpoint string) (_builer *APIBuilder) {
	builder.endPoint = endpoint
	return builder
}

func (builder *APIBuilder) FuturesLever(lever float64) (_builder *APIBuilder) {
	builder.futuresLever = lever
	return builder
}

func (builder *APIBuilder) Build(exName string) (api API) {
	var _api API
	switch exName {
	case OKEX:
		_api = okexV5.NewOKExV5Spot(&APIConfig{
			HttpClient:    builder.client,
			ApiKey:        builder.apiKey,
			ApiSecretKey:  builder.secretkey,
			ApiPassphrase: builder.apiPassphrase,
			Endpoint:      builder.endPoint,
		})
	default:
		println("exchange name error [" + exName + "].")

	}
	return _api
}

func (builder *APIBuilder) BuildFuture(exName string) (api FutureRestAPI) {
	switch exName {
	default:
		println(fmt.Sprintf("%s not support future", exName))
		return nil
	}
}

func (builder *APIBuilder) BuildFuturesWs(exName string) (FuturesWsApi, error) {
	switch exName {
	case OKEX_V3, OKEX, OKEX_FUTURE:
		return okex.NewOKExV3FuturesWs(okex.NewOKEx(&APIConfig{
			HttpClient: builder.client,
			Endpoint:   builder.futuresEndPoint,
		})), nil
	}
	return nil, errors.New("not support the exchange " + exName)
}

func (builder *APIBuilder) BuildSpotWs(exName string) (SpotWsApi, error) {
	switch exName {
	case OKEX_V3, OKEX:
		return okex.NewOKExSpotV3Ws(nil), nil
	}
	return nil, errors.New("not support the exchange " + exName)
}

func (builder *APIBuilder) BuildWallet(exName string) (WalletApi, error) {
	switch exName {
	case OKEX_V3, OKEX:
		return okex.NewOKEx(&APIConfig{
			HttpClient:    builder.client,
			ApiKey:        builder.apiKey,
			ApiSecretKey:  builder.secretkey,
			ApiPassphrase: builder.apiPassphrase,
		}).OKExWallet, nil
	}
	return nil, errors.New("not support the wallet api for  " + exName)
}
