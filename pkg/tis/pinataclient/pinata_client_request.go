package pinataclient

type PinataClientRequest struct {
	BearerToken           string
	PinningServiceBaseUrl string
	FilePinBaseUrl        string
	PinataApiKey          string
	PinataSecretApiKey    string
}
