package api

type PinataClient struct {
}

func NewPinataClient() (*PinataClient, error) {
	return nil, nil
}

// HashMetaData  Modify stored file metadata.
func (p *PinataClient) HashMetaData() error {
	return nil
}

// HashPinPolicy  Allows the user to change the pin policy for an individual piece of content.
func (p *PinataClient) HashPinPolicy() error {
	return nil
}

//PinByHash  Adds a hash to Pinata's pin queue to be pinned asynchronously.
func (p *PinataClient) PinByHash() error {
	return nil
}

//PinJobs This endpoint allows users to search for the status of all hashes that are currently in Pinata's pin queue.
func (p *PinataClient) PinJobs() error {
	return nil
}

// PinFileToIPFS Send JSON to Pinata for direct pinning to IPFS.
func (p *PinataClient) PinFileToIPFS() error {
	return nil
}

//PinJSONToIPFS Send JSON to Pinata for direct pinning to IPFS.
func (p *PinataClient) PinJSONToIPFS() error {
	return nil
}

//Unpin   Have Pinata unpin content that you've pinned through the service.
func (p *PinataClient) Unpin() error {
	return nil
}

//TestAuthentication Tests that you can authenticate with Pinata correctly
func (p *PinataClient) TestAuthentication() error {
	return nil
}

//PinList  Retrieve pin records for your Pinata account
func (p *PinataClient) PinList() error {
	return nil
}

//UserPinnedDataTotal  Returns the total combined size (in bytes) of all content you currently have pinned on Pinata.
func (p *PinataClient) UserPinnedDataTotal() error {
	return nil
}
