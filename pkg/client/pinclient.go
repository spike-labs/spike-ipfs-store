package client

import (
	"context"
	logging "github.com/ipfs/go-log/v2"
	"github.com/jooyyy/pinata-go/pkg/tis"
	"github.com/jooyyy/pinata-go/pkg/tis/api"
	"github.com/jooyyy/pinata-go/pkg/tis/pinataclient"
	"github.com/pkg/errors"
	"net/http"
)

var logger = logging.Logger("ipfs-store")

type Client struct {
	httpClient *http.Client
	tisClient  api.IPFSPin
	ctx        context.Context
}

func New(request ClientCreateRequest) (*Client, error) {
	tisclient, _ := newTISClient(request)

	return &Client{httpClient: request.GetHttpClient(), tisClient: tisclient}, nil
}

func (c *Client) PinFileToIPFS(ctx context.Context, filePath string, opts ...PinataOptions) (*http.Response, error) {

	options, err := processPinataOptions(opts...)
	if err != nil {
		return nil, err
	}

	request := &api.PinataRequest{
		PinataOptions: api.PinataOptions{
			CidVersion:        options.CidVersion,
			WrapWithDirectory: options.WrapWithDirectory,
			CustomPinPolicy: api.CustomPinPolicy{
				Regions: options.CustomPinPolicy.Regions,
			},
		},
		PinataMetaData: options.PinataMetaData,
	}
	// todo The bottom method packaging req returns to the upper layer for processing.
	req, err := c.tisClient.PinFileToIPFS(ctx, request)
	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func newTISClient(pinningService ClientCreateRequest) (api.IPFSPin, error) {
	var err error
	var tisClient api.IPFSPin
	switch pinningService.ps.String() {
	case tis.Pinata.String():
		request := &pinataclient.PinataClientRequest{
			BearerToken:           pinningService.GetBearerToken(),
			PinningServiceBaseUrl: pinningService.GetPinningServiceBaseUrl(),
			FilePinBaseUrl:        pinningService.GetFilePinBaseUrl(),
			PinataApiKey:          pinningService.GetPinataApiKey(),
			PinataSecretApiKey:    pinningService.GetPinataSecretApiKey(),
		}
		tisClient, err = api.NewPinataClient(request)
		if err != nil {
			return nil, errors.WithMessage(err, "failed to create Pinata Client")
		}
	default:
		logger.Fatalf("only pinata supported for file upload")
	}
	return tisClient, nil
}

func processPinataOptions(opts ...PinataOptions) (*PinataOption, error) {
	options := PinataOption{}
	for _, param := range opts {
		err := param(&options)
		if err != nil {
			return nil, errors.WithMessage(err, "failed to create CA Client")
		}
	}
	return &options, nil
}
