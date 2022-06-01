package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/Fueav/spike-ipfs-store/pkg/tis/pinataclient"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type PinataClient struct {
	bearerToken           string
	pinningServiceBaseUrl string
	filePinBaseUrl        string
	pinataApiKey          string
	pinataSecretApiKey    string
}

func NewPinataClient(pinataClientInfo *pinataclient.PinataClientRequest) (*PinataClient, error) {
	pinataClient := &PinataClient{
		bearerToken:           pinataClientInfo.BearerToken,
		pinningServiceBaseUrl: pinataClientInfo.PinningServiceBaseUrl,
		filePinBaseUrl:        pinataClientInfo.FilePinBaseUrl,
		pinataApiKey:          pinataClientInfo.PinataApiKey,
		pinataSecretApiKey:    pinataClientInfo.PinataSecretApiKey,
	}
	return pinataClient, nil
}

func (p *PinataClient) NewRequestWithHeaders(method, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("pinata_api_key", p.pinataApiKey)
	req.Header.Set("pinata_secret_api_key", p.pinataSecretApiKey)
	return req, nil
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
func (p *PinataClient) PinFileToIPFS(ctx context.Context, request *PinataRequest, filePath string) (*http.Request, error) {
	var b bytes.Buffer
	var err error
	w := multipart.NewWriter(&b)
	localVarFormParams := url.Values{}

	if request.PinataOptions != nil {
		paramJson, err := parameterToJson(*request.PinataOptions)
		if err != nil {
			return nil, err
		}
		localVarFormParams.Add("pinataOptions", paramJson)
	}
	if request.PinataMetaData != nil {
		paramJson, err := parameterToJson(*request.PinataMetaData)
		if err != nil {
			return nil, err
		}
		localVarFormParams.Add("pinataMetadata", paramJson)
	}

	err = createMultipartFormData(localVarFormParams, w)
	if err != nil {
		return nil, err
	}

	err = createMultipartFormFile(filePath, w)
	if err != nil {
		return nil, err
	}

	err = w.Close()
	if err != nil {
		return nil, err
	}
	req, err := p.NewRequestWithHeaders("POST", p.pinningServiceBaseUrl+"/pinning/pinFileToIPFS", &b)

	req.Header.Set("Content-Type", w.FormDataContentType())
	fmt.Printf("%+v", localVarFormParams)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func createMultipartFormFile(filePath string, w *multipart.Writer) error {
	var err error
	var fw io.Writer

	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	if fw, err = w.CreateFormFile("file", formatFilename(filePath)); err != nil {
		return err
	}

	if _, err = io.Copy(fw, file); err != nil {
		return err
	}
	return nil
}

func createMultipartFormData(formParams url.Values, w *multipart.Writer) error {

	if len(formParams) != 0 {
		for k, v := range formParams {
			for _, iv := range v {
				// form value
				err := w.WriteField(k, iv)
				if err != nil {
					fmt.Println("add param error", err)
					return err
				}
			}
		}
	}

	return nil
}

func formatFilename(path string) string {
	items := strings.Split(path, "/")
	if len(items) > 0 {
		return items[len(items)-1]
	}
	return ""
}

// helper for converting interface{} parameters to json strings
func parameterToJson(obj interface{}) (string, error) {
	jsonBuf, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}
	return string(jsonBuf), err
}

//PinJSONToIPFS Send JSON to Pinata for direct pinning to IPFS.
func (p *PinataClient) PinJSONToIPFS(ctx context.Context, request *PinataRequest, jsonString string) (*http.Request, error) {

	paramJson, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	fmt.Println("-------------------")
	fmt.Printf("%+v", string(paramJson))
	fmt.Println("\n-------------------")

	req, err := p.NewRequestWithHeaders("POST", p.pinningServiceBaseUrl+"/pinning/pinJSONToIPFS", bytes.NewBuffer(paramJson))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	if err != nil {
		return nil, err
	}
	return req, nil
}

//Unpin   Have Pinata unpin content that you've pinned through the service.
func (p *PinataClient) Unpin() error {
	return nil
}

//TestAuthentication Tests that you can authenticate with Pinata correctly
func (p *PinataClient) TestAuthentication() (*http.Request, error) {
	req, err := p.NewRequestWithHeaders("GET", p.pinningServiceBaseUrl+"/data/testAuthentication", nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

//PinList  Retrieve pin records for your Pinata account
func (p *PinataClient) PinList() error {
	return nil
}

//UserPinnedDataTotal  Returns the total combined size (in bytes) of all content you currently have pinned on Pinata.
func (p *PinataClient) UserPinnedDataTotal() error {
	return nil
}
