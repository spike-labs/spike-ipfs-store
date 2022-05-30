package api

import "context"

type IPFSPin interface {
	// Pinning
	PinFileToIPFS(ctx context.Context, request *PinataRequest) error // Send JSON to Pinata for direct pinning to IPFS.
	Unpin() error                                                    // Have Pinata unpin content that you've pinned through the service.

	// Data
	TestAuthentication() error // Tests that you can authenticate with Pinata correctly
	PinList() error            // Retrieve pin records for your Pinata account
}
