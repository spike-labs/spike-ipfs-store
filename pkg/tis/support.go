package tis

type PinningService string

const (
	Pinata PinningService = "pinata"
)

func (f PinningService) GetPinningServiceBaseUrl() string {
	switch f {
	case Pinata:
		return "https://api.pinata.cloud"
	}

	return ""
}

func (f PinningService) GetFilePinBaseUrl() string {
	switch f {
	case Pinata:
		return "https://api.pinata.cloud"
	}
	panic("unsupported file pin support")
}

func (f PinningService) String() string {
	return string(f)
}
