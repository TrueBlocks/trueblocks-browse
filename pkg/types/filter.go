package types

import "encoding/json"

type Criteria string

type Filter struct {
	Criteria string `json:"criteria"`
}

func (f *Filter) String() string {
	bytes, _ := json.Marshal(f)
	return string(bytes)
}

func (f *Filter) HasCriteria() bool {
	return f.Criteria != ""
}

type Nothing struct {
	Unused string `json:"unused"`
}

