package gasstation

type GasPrice struct {
	Fast        float64 `json:"fast"`
	Fastest     float64 `json:"fastest"`
	SafeLow     float64 `json:"safeLow"`
	Average     float64 `json:"average"`
	BlockTime   float64 `json:"block_time"`
	BlockNumber float64 `json:"blockNum"`
	Speed       float64 `json:"speed"`
	SafeLowWait float64 `json:"safeLowWait"`
	AverageWait float64 `json:"avgWait"`
	FastWait    float64 `json:"fastWait"`
	FastestWait float64 `json:"fastestWait"`
}

type Prediction struct {
	GasPrice            float64 `json:"gasprice"`
	HashPowerAccepting  float64 `json:"hashpower_accepting"`
	HashPowerAccepting2 float64 `json:"hashpower_accepting2"`
	TxAtabove           int     `json:"tx_atabove"`
	Age                 float64 `json:"age"`
	PctRemaining5m      int     `json:"pct_remaining5m"`
	PctMined5m          int     `json:"pct_mined_5m"`
	TotalSeen5m         int     `json:"total_seen_5m"`
	PctRemaining30m     int     `json:"pct_remaining30m"`
	TotalSeen30m        int     `json:"total_seen_30m"`
	Average             int     `json:"average"`
	SafeLow             int     `json:"safelow"`
	NoMine              int     `json:"nomine"`
	AvgDiff             int     `json:"avgdiff"`
	Intercept           float64 `json:"intercept"`
	HpaCoef             float64 `json:"hpa_coef"`
	AvgDiffCoef         float64 `json:"avgdiff_coef"`
	TxAtaboveCoef       float64 `json:"tx_atabove_coef"`
	Int2                float64 `json:"int2"`
	HpaCoef2            float64 `json:"hpa_coef2"`
	Sum                 float64 `json:"sum"`
	ExpectedWait        float64 `json:"expectedWait"`
	ExpectedTime        float64 `json:"expectedTime"`
}
