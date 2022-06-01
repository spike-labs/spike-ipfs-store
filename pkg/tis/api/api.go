package api

import (
	"context"
	"net/http"
)

type IPFSPin interface {
	// Pinning
	PinFileToIPFS(ctx context.Context, request *PinataRequest, filePath string) (*http.Request, error) // Send JSON to Pinata for direct pinning to IPFS.
	Unpin() error                                                                                      // Have Pinata unpin content that you've pinned through the service.

	// Data
	TestAuthentication() (*http.Request, error) // Tests that you can authenticate with Pinata correctly
	PinList() error                             // Retrieve pin records for your Pinata account
}
