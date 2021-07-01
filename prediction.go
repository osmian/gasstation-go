package gasstation

import (
	"net/http"
)

const (
	PredictionTableEndpoint = "predictTable.json"
)

// Returns estimated confirmation times for a range of gas prices from the EGS Prediction Page.
func (client *gasStationClient) GetPredictionTable() ([]Prediction, error) {
	predictions := []Prediction{}
	_, err := client.Request(http.MethodGet, PredictionTableEndpoint, &predictions)
	if err != nil {
		return nil, err
	}
	return predictions, nil
}
