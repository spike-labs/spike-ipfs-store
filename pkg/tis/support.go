package tis

import (
	"net"
	"net/http"
	"time"
)

type PinningService string

// Pinata third part IPFS service
const (
	Pinata PinningService = "pinata"
)

// http.Client config
const (
	MaxIdleConns        int = 100
	MaxIdleConnsPerHost int = 100
	IdleConnTimeout     int = 90
)

func (f PinningService) GetPinningServiceBaseUrl() string {
	switch f {
	case Pinata:
		return "https://api.pinata.cloud"
	}

	return ""
}

func (f PinningService) GetFilePinBaseUrl() string {
	switch f {
	case Pinata:
		return "https://api.pinata.cloud"
	}
	panic("unsupported file pin support")
}

func (f PinningService) String() string {
	return string(f)
}

func (f PinningService) CreateHTTPClient() *http.Client {
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
			MaxIdleConns:        MaxIdleConns,
			MaxIdleConnsPerHost: MaxIdleConnsPerHost,
			IdleConnTimeout:     time.Duration(IdleConnTimeout) * time.Second,
		},
		Timeout: 20 * time.Second,
	}
	return client
}
