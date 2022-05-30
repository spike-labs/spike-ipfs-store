package api

import "github.com/jooyyy/pinata-go/pkg/client"

type PinataRequest struct {
	PinataOptions  PinataOptions
	PinataMetaData client.PinataMetaData
}

type PinataOptions struct {
	CidVersion        int
	WrapWithDirectory bool
	CustomPinPolicy   client.CustomPinPolicy
}
