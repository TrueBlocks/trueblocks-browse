package config

type Layout struct {
	Header bool `json:"header"`
	Menu   bool `json:"menu"`
	Help   bool `json:"help"`
	Footer bool `json:"footer"`
}

type Headers struct {
	Project   bool `json:"project"`
	History   bool `json:"history"`
	Monitors  bool `json:"monitors"`
	Names     bool `json:"names"`
	Abis      bool `json:"abis"`
	Indexes   bool `json:"indexes"`
	Manifests bool `json:"manifests"`
	Status    bool `json:"status"`
	Settings  bool `json:"settings"`
}

type Daemons struct {
	Freshen bool `json:"freshen"`
	Scraper bool `json:"scraper"`
	Ipfs    bool `json:"ipfs"`
}

type Toggles struct {
	Layout  Layout  `json:"layout"`
	Headers Headers `json:"headers"`
	Daemons Daemons `json:"daemons"`
}

func (t *Toggles) IsOn(which string) bool {
	if which == "" {
		which = "project"
	}
	switch which {
	case "header":
		return t.Layout.Header
	case "menu":
		return t.Layout.Menu
	case "help":
		return t.Layout.Help
	case "footer":
		return t.Layout.Footer
	case "project":
		return t.Headers.Project
	case "history":
		return t.Headers.History
	case "monitors":
		return t.Headers.Monitors
	case "names":
		return t.Headers.Names
	case "abis":
		return t.Headers.Abis
	case "indexes":
		return t.Headers.Indexes
	case "manifests":
		return t.Headers.Manifests
	case "status":
		return t.Headers.Status
	case "settings":
		return t.Headers.Settings
	case "freshen":
		return t.Daemons.Freshen
	case "scraper":
		return t.Daemons.Scraper
	case "ipfs":
		return t.Daemons.Ipfs
	}
	return false
}

func (t *Toggles) SetState(which string, onOff bool) {
	if which == "" {
		which = "project"
	}
	switch which {
	case "header":
		t.Layout.Header = onOff
	case "menu":
		t.Layout.Menu = onOff
	case "help":
		t.Layout.Help = onOff
	case "footer":
		t.Layout.Footer = onOff
	case "project":
		t.Headers.Project = onOff
	case "history":
		t.Headers.History = onOff
	case "monitors":
		t.Headers.Monitors = onOff
	case "names":
		t.Headers.Names = onOff
	case "abis":
		t.Headers.Abis = onOff
	case "indexes":
		t.Headers.Indexes = onOff
	case "manifests":
		t.Headers.Manifests = onOff
	case "status":
		t.Headers.Status = onOff
	case "settings":
		t.Headers.Settings = onOff
	case "freshen":
		t.Daemons.Freshen = onOff
	case "scraper":
		t.Daemons.Scraper = onOff
	case "ipfs":
		t.Daemons.Ipfs = onOff
	}
}
