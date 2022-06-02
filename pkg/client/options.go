package client

import (
	"github.com/Fueav/spike-ipfs-store/pkg/tis/pinataclient"
)

// pinataclient options collector
type PinataOption struct {
	CidVersion        int
	WrapWithDirectory bool
	CustomPinPolicy   *pinataclient.CustomPinPolicy
	PinataMetaData    *pinataclient.PinataMetaData
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
func WithCustomPinPolicy(customPinPolicy *pinataclient.CustomPinPolicy) PinataOptions {
	return func(c *PinataOption) error {
		c.CustomPinPolicy = customPinPolicy
		return nil
	}
}

func WithPinataMetaData(pinataMetaData *pinataclient.PinataMetaData) PinataOptions {
	return func(c *PinataOption) error {
		c.PinataMetaData = pinataMetaData
		return nil
	}
}
