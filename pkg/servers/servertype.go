package servers

type Type int

const (
	ST_FileServer Type = iota
	ST_Scraper
	ST_Freshen
	ST_Api
	ST_Ipfs
)

func (s Type) String() string {
	m := map[Type]string{
		ST_FileServer: "FileServer",
		ST_Scraper:    "Scraper",
		ST_Freshen:    "Freshen",
		ST_Api:        "Api",
		ST_Ipfs:       "IPFS",
	}
	return m[s]
}

var Types = []struct {
	Value  Type
	TSName string
}{
	{ST_FileServer, "FILESERVER"},
	{ST_Scraper, "SCRAPER"},
	{ST_Freshen, "FRESHEN"},
	{ST_Api, "API"},
	{ST_Ipfs, "IPFS"},
}
