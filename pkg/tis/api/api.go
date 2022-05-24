package api

type IPFSPin interface {
	// Pinning
	pinFileToIPFS() error // Send JSON to Pinata for direct pinning to IPFS.
	unpin() error         // Have Pinata unpin content that you've pinned through the service.

	// Data
	testAuthentication() error // Tests that you can authenticate with Pinata correctly
	pinList() error            // Retrieve pin records for your Pinata account
}
