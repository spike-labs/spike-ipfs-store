package client

// client options collector
type PinataOptions struct {
	CidVersion        int
	WrapWithDirectory bool
	CustomPinPolicy   CustomPinPolicy `json:"customPinPolicy"`
}

type CustomPinPolicy struct {
	Regions []Regions `json:"regions"`
}

type Regions struct {
	ID                      string `json:"id"`
	DesiredReplicationCount int    `json:"desiredReplicationCount"`
}

// PinataOption  describes a functional parameter for the New constructor
type PinataOption func(*PinataOptions) error

// WithCidVersion  option
func WithCidVersion(cidVersion int) PinataOption {
	return func(c *PinataOptions) error {
		c.CidVersion = cidVersion
		return nil
	}
}

// WithWrapWithDirectory Wrap your content inside of a directory when adding to IPFS.
func WithWrapWithDirectory(wrapWithDirectory bool) PinataOption {
	return func(c *PinataOptions) error {
		c.WrapWithDirectory = wrapWithDirectory
		return nil
	}
}

// WithCustomPinPolicy  a custom pin policy for the piece of content being pinned.
func WithCustomPinPolicy(customPinPolicy CustomPinPolicy) PinataOption {
	return func(c *PinataOptions) error {
		c.CustomPinPolicy = customPinPolicy
		return nil
	}
}

type PinataMetaData struct {
	Name      string            `json:"name"`
	Keyvalues map[string]string `json:"keyvalues"`
}

type PinataMetaOptions func(*PinataMetaData) error

func WithName(name string) PinataMetaOptions {
	return func(c *PinataMetaData) error {
		c.Name = name
		return nil
	}
}

// WithKeyValues todo map初始化的问题
func WithKeyValues(metaData map[string]string) PinataMetaOptions {
	return func(c *PinataMetaData) error {
		c.Keyvalues = metaData
		return nil
	}
}
