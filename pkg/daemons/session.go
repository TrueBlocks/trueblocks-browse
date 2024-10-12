package daemons

// Toggles stores the disable/enable state of the Daemons
type Toggles struct {
	Freshen bool `json:"freshen"`
	Scraper bool `json:"scraper"`
	Ipfs    bool `json:"ipfs"`
}
