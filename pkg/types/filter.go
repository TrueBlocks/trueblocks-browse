package types

import "encoding/json"

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

type Nothing struct{}
