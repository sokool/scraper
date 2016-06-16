package model

type Element struct {
	Selector   string `json:"selector"`
	Action     string `json:"action"`
	Filters    string `json:"filters"`
	Validators string `json:"validators"`
}
