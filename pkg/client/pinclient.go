package client

import (
	"context"
	logging "github.com/ipfs/go-log/v2"
	"github.com/jooyyy/pinata-go/pkg/tis"
	"github.com/jooyyy/pinata-go/pkg/tis/api"
	"github.com/pkg/errors"
)

var logger = logging.Logger("ipfs-store")

type Client struct {
	pinningService string
	ctx            context.Context
}

func New(ctx context.Context, opts ...ClientOption) (*Client, error) {
	return nil, nil
}

func (c *Client) PinFileToIPFS(ctx context.Context) error {
	tisclient, _ := newTISClient(c.ctx, c.pinningService)
	tisclient.PinFileToIPFS()
	return nil
}

func newTISClient(ctx context.Context, pinningService string) (api.IPFSPin, error) {
	var err error
	var tisClient api.IPFSPin
	switch pinningService {
	case tis.Pinata.String():
		tisClient, err = api.NewPinataClient()
		if err != nil {
			return nil, errors.WithMessage(err, "failed to create Pinata Client")
		}
	default:
		logger.Fatalf("only pinata supported for file upload")
	}
	return tisClient, nil
}
