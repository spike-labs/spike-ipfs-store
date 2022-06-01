package client

import "github.com/Fueav/spike-ipfs-store/pkg/tis/api"

// pinataclient options collector
type PinataOption struct {
	CidVersion        int
	WrapWithDirectory bool
	CustomPinPolicy   *api.CustomPinPolicy
	PinataMetaData    *api.PinataMetaData
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
func WithCustomPinPolicy(customPinPolicy *api.CustomPinPolicy) PinataOptions {
	return func(c *PinataOption) error {
		c.CustomPinPolicy = customPinPolicy
		return nil
	}
}

func WithPinataMetaData(pinataMetaData *api.PinataMetaData) PinataOptions {
	return func(c *PinataOption) error {
		c.PinataMetaData = pinataMetaData
		return nil
	}
}
