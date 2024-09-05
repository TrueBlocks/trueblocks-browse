package daemons

type Type string

const (
	ST_FileDaemon Type = "FileDaemon"
	ST_Scraper    Type = "Scraper"
	ST_Freshen    Type = "Freshen"
	ST_Api        Type = "Api"
	ST_Ipfs       Type = "Ipfs"
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
