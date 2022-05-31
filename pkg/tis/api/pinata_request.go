package api

type PinataRequest struct {
	PinataOptions  PinataOptions  `json:"pinataOptions"`
	PinataMetaData PinataMetaData `json:"pinataMetadata"`
}

type PinataOptions struct {
	CidVersion        int             `json:"cidVersion"`
	WrapWithDirectory bool            `json:"wrapWithDirectory"`
	CustomPinPolicy   CustomPinPolicy `json:"customPinPolicy"`
}

type PinataMetaData struct {
	Name      string            `json:"name"`
	Keyvalues map[string]string `json:"keyvalues"`
}

type CustomPinPolicy struct {
	Regions []Regions `json:"regions"`
}

type Regions struct {
	ID                      string `json:"id"`
	DesiredReplicationCount int    `json:"desiredReplicationCount"`
}
