package api

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/jooyyy/pinata-go/pkg/tis/pinataclient"
	"io"
	"net/http"
)

type PinataClient struct {
	bearerToken           string
	pinningServiceBaseUrl string
	filePinBaseUrl        string
	pinataApiKey          string
	pinataSecretApiKey    string
}

// todo Pass in the information of building req to build the pinataclient.
func NewPinataClient(pinataClientInfo *pinataclient.PinataClientRequest) (*PinataClient, error) {
	pinataClient := &PinataClient{
		bearerToken:           pinataClientInfo.BearerToken,
		pinningServiceBaseUrl: pinataClientInfo.PinningServiceBaseUrl,
		filePinBaseUrl:        pinataClientInfo.FilePinBaseUrl,
		pinataApiKey:          pinataClientInfo.PinataApiKey,
		pinataSecretApiKey:    pinataClientInfo.PinataSecretApiKey,
	}
	return pinataClient, nil
}

func (p *PinataClient) NewRequestWithHeaders(method, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("pinata_api_key", p.pinataApiKey)
	req.Header.Set("pinata_secret_api_key", p.pinataSecretApiKey)
	return req, nil
}

// HashMetaData  Modify stored file metadata.
func (p *PinataClient) HashMetaData() error {
	return nil
}

// HashPinPolicy  Allows the user to change the pin policy for an individual piece of content.
func (p *PinataClient) HashPinPolicy() error {
	return nil
}

//PinByHash  Adds a hash to Pinata's pin queue to be pinned asynchronously.
func (p *PinataClient) PinByHash() error {
	return nil
}

//PinJobs This endpoint allows users to search for the status of all hashes that are currently in Pinata's pin queue.
func (p *PinataClient) PinJobs() error {
	return nil
}

// PinFileToIPFS Send JSON to Pinata for direct pinning to IPFS.
func (p *PinataClient) PinFileToIPFS(ctx context.Context, request *PinataRequest) (*http.Request, error) {
	po, err := json.Marshal(request.PinataOptions)
	if err != nil {
		return nil, err
	}
	pm, err := json.Marshal(request.PinataMetaData)
	if err != nil {
		return nil, err
	}
	b := append(po, pm...)
	req, err := p.NewRequestWithHeaders("POST", p.pinningServiceBaseUrl+"/pinning/pinHashToIPFS", bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	return req, nil
}

//PinJSONToIPFS Send JSON to Pinata for direct pinning to IPFS.
func (p *PinataClient) PinJSONToIPFS() error {
	return nil
}

//Unpin   Have Pinata unpin content that you've pinned through the service.
func (p *PinataClient) Unpin() error {
	return nil
}

//TestAuthentication Tests that you can authenticate with Pinata correctly
func (p *PinataClient) TestAuthentication() (*http.Request, error) {
	req, err := p.NewRequestWithHeaders("GET", p.pinningServiceBaseUrl+"/data/testAuthentication", nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

//PinList  Retrieve pin records for your Pinata account
func (p *PinataClient) PinList() error {
	return nil
}

//UserPinnedDataTotal  Returns the total combined size (in bytes) of all content you currently have pinned on Pinata.
func (p *PinataClient) UserPinnedDataTotal() error {
	return nil
}
