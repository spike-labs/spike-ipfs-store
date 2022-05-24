package client

// client options collector
type clientOptions struct {
	pinningService string
}

// ClientOption describes a functional parameter for the New constructor
type ClientOption func(*clientOptions) error

// WithPinService option
func WithPinService(pinningService string) ClientOption {
	return func(c *clientOptions) error {
		c.pinningService = pinningService
		return nil
	}
}
