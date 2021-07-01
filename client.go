package gasstation

/*
Wraps API for https://docs.ethgasstation.info/
*/

import (
	"encoding/json"
	"net/http"
	"net/url"
	"path"
)

type GasStationClient interface {
	GetPredictionTable() ([]Prediction, error)
	GetGasPrice() (*GasPrice, error)
}

type gasStationClient struct {
	baseURI    string
	httpClient *http.Client
	apiKey     string
}

const (
	DefaultBaseURI = "https://ethgasstation.info/api/"
)

func NewGasStationClient(apikey string) GasStationClient {
	return &gasStationClient{
		baseURI:    DefaultBaseURI,
		httpClient: http.DefaultClient,
		apiKey:     apikey,
	}
}

func (c *gasStationClient) Request(httpMethod string, endpoint string, result interface{}) (resp *http.Response, err error) {

	requestURL, err := createRequestURL(c.baseURI, endpoint, c.apiKey)
	if err != nil {
		return resp, err
	}

	req, err := http.NewRequest(httpMethod, requestURL, nil)
	if err != nil {
		return resp, err
	}

	resp, err = c.httpClient.Do(req)
	if err != nil {
		return resp, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return resp, err
	}

	decoder := json.NewDecoder(resp.Body)
	if err = decoder.Decode(result); err != nil {
		return resp, err
	}
	return resp, nil
}

func createRequestURL(baseURL, endpoint, apiKey string) (string, error) {
	u, err := url.Parse(baseURL)
	if err != nil {
		return "", err
	}
	u.Path = path.Join(u.Path, endpoint)
	q := u.Query()
	q.Set("api-key", apiKey)
	u.RawQuery = q.Encode()
	return u.String(), nil
}
