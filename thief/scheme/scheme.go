package scheme

type Scheme struct {
	Name       string `json:"name"`
	Identifier *Element `json:"identifier"`
	Structure  map[string]*Element `json:"structure"`
}


