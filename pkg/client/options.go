package client

// client options collector
type PinataOption struct {
	CidVersion        int
	WrapWithDirectory bool
	CustomPinPolicy   CustomPinPolicy
	PinataMetaData    PinataMetaData
}

type CustomPinPolicy struct {
	Regions []Regions
}

type Regions struct {
	ID                      string
	DesiredReplicationCount int
}

type PinataMetaData struct {
	Name      string            `json:"name"`
	Keyvalues map[string]string `json:"keyvalues"`
}

// PinataOption  describes a functional parameter for the New constructor
type PinataOptions func(*PinataOption) error

// WithCidVersion  option
func WithCidVersion(cidVersion int) PinataOptions {
	return func(c *PinataOption) error {
		c.CidVersion = cidVersion
		return nil
	}
}

// WithWrapWithDirectory Wrap your content inside of a directory when adding to IPFS.
func WithWrapWithDirectory(wrapWithDirectory bool) PinataOptions {
	return func(c *PinataOption) error {
		c.WrapWithDirectory = wrapWithDirectory
		return nil
	}
}

// WithCustomPinPolicy  a custom pin policy for the piece of content being pinned.
func WithCustomPinPolicy(customPinPolicy CustomPinPolicy) PinataOptions {
	return func(c *PinataOption) error {
		c.CustomPinPolicy = customPinPolicy
		return nil
	}
}

func WithPinataMetaData(pinataMetaData PinataMetaData) PinataOptions {
	return func(c *PinataOption) error {
		c.PinataMetaData = pinataMetaData
		return nil
	}
}
