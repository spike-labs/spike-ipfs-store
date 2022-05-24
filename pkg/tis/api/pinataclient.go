package api

type PinataClient struct {
}

func NewPinataClient() (*PinataClient, error) {
	return nil, nil
}

//hashMetaData() error  // Modify stored file metadata.
func (p *PinataClient) hashMetaData() error {
	return nil
}

//hashPinPolicy() Allows the user to change the pin policy for an individual piece of content.
func (p *PinataClient) hashPinPolicy() error {
	return nil
}

//pinByHash()  Adds a hash to Pinata's pin queue to be pinned asynchronously.
func (p *PinataClient) pinByHash() error {
	return nil
}

//pinJobs()  This endpoint allows users to search for the status of all hashes that are currently in Pinata's pin queue.
func (p *PinataClient) pinJobs() error {
	return nil
}

// Send JSON to Pinata for direct pinning to IPFS.
func (p *PinataClient) pinFileToIPFS() error {
	return nil
}

//pinJSONToIPFS() Send JSON to Pinata for direct pinning to IPFS.
func (p *PinataClient) pinJSONToIPFS() error {
	return nil
}

//unpin()  Have Pinata unpin content that you've pinned through the service.
func (p *PinataClient) unpin() error {
	return nil
}

//testAuthentication() Tests that you can authenticate with Pinata correctly
func (p *PinataClient) testAuthentication() error {
	return nil
}

//pinList()  Retrieve pin records for your Pinata account
func (p *PinataClient) pinList() error {
	return nil
}

//userPinnedDataTotal()  Returns the total combined size (in bytes) of all content you currently have pinned on Pinata.
func (p *PinataClient) userPinnedDataTotal() error {
	return nil
}
