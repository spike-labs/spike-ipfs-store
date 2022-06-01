package api

import (
	"context"
	"net/http"
)

type IPFSPin interface {
	// PinFileToIPFS Pinning
	PinFileToIPFS(ctx context.Context, request *PinataRequest, filePath string) (*http.Request, error)
	// PinJSONToIPFS Send JSON to Pinata for direct pinning to IPFS.
	PinJSONToIPFS(ctx context.Context, request *PinataRequest, json string) (*http.Request, error)
	Unpin() error // Have Pinata unpin content that you've pinned through the service.

	// Data
	TestAuthentication() (*http.Request, error) // Tests that you can authenticate with Pinata correctly
	PinList() error                             // Retrieve pin records for your Pinata account
}
