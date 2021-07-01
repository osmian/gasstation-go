package gasstation

import "net/http"

const (
	GasPriceEndpoint = "ethgasAPI.json"
)

// Returns the current recommended fast, standard and safe low gas prices on the Ethereum network, along with the current block and wait times for each "speed".
func (client *gasStationClient) GetGasPrice() (*GasPrice, error) {
	gasPrice := &GasPrice{}
	_, err := client.Request(http.MethodGet, GasPriceEndpoint, gasPrice)
	if err != nil {
		return nil, err
	}
	return gasPrice, nil
}
