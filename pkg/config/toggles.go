package config

import "encoding/json"

type Toggles struct {
	Header bool `json:"header"`
	Menu   bool `json:"menu"`
	Help   bool `json:"help"`
	Footer bool `json:"footer"`
}

func (t *Toggles) String() string {
	bytes, _ := json.Marshal(t)
	return string(bytes)
}
