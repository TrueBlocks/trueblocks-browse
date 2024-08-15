package daemons

type Type int

const (
	ST_FileDaemon Type = iota
	ST_Scraper
	ST_Freshen
	ST_Api
	ST_Ipfs
)

func (s Type) String() string {
	m := map[Type]string{
		ST_FileDaemon: "FileDaemon",
		ST_Scraper:    "Scraper",
		ST_Freshen:    "Freshen",
		ST_Api:        "Api",
		ST_Ipfs:       "IPFS",
	}
	return m[s]
}
