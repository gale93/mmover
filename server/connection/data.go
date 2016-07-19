package connection

type data struct {
	Latitude  string `json:"lat"`
	Longitude string `json:"lng"`

	UpdateRatems int `json:"update_rate_ms"`
}

var app data
