package structs

type Balances struct {
	Success bool `json:"success"`
	Result  []struct {
		Coin  string  `json:"coin"`
		Free  float64 `json:"free"`
		Total float64 `json:"total"`
	} `json:"result"`
}
