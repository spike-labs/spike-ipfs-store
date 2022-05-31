package client

import (
	"github.com/jooyyy/pinata-go/pkg/tis"
	"net/http"
)

type ClientCreateRequest struct {
	ps                    tis.PinningService
	pinningServiceBaseUrl string
	filePinBaseUrl        string
	bearerToken           string
	pinataApiKey          string
	pinataSecretApiKey    string
	httpClient            *http.Client
}

func NewClientRequest(ps tis.PinningService) ClientCreateRequest {
	request := ClientCreateRequest{ps: ps}
	request.pinningServiceBaseUrl = ps.GetPinningServiceBaseUrl()
	request.filePinBaseUrl = ps.GetFilePinBaseUrl()
	request.httpClient = ps.CreateHTTPClient()
	return request
}

func (r ClientCreateRequest) BearerToken(token string) ClientCreateRequest {
	r.bearerToken = token
	return r
}

func (r ClientCreateRequest) PinningServiceBaseUrl(url string) ClientCreateRequest {
	r.pinningServiceBaseUrl = url
	return r
}

func (r ClientCreateRequest) FilePinBaseUrl(url string) ClientCreateRequest {
	r.filePinBaseUrl = url
	return r
}

func (r ClientCreateRequest) PinataApiKey(key string) ClientCreateRequest {
	r.pinataApiKey = key
	return r
}

func (r ClientCreateRequest) PinataSecretApiKey(secretKey string) ClientCreateRequest {
	r.pinataSecretApiKey = secretKey
	return r
}

// HttpClient Users can initialize http.pinataclient replacement for automatic initialization
func (r ClientCreateRequest) HttpClient(client http.Client) ClientCreateRequest {
	r.httpClient = &client
	return r
}

func (r ClientCreateRequest) GetPinningServiceBaseUrl() string {
	return r.pinningServiceBaseUrl
}

func (r ClientCreateRequest) GetFilePinBaseUrl() string {
	return r.filePinBaseUrl
}

func (r ClientCreateRequest) GetPinataApiKey() string {
	return r.pinataApiKey
}

func (r ClientCreateRequest) GetPinataSecretApiKey() string {
	return r.pinataSecretApiKey
}

func (r ClientCreateRequest) GetBearerToken() string {
	return r.bearerToken
}

func (r ClientCreateRequest) GetHttpClient() *http.Client {
	return r.httpClient
}
