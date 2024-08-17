package screen

import "fmt"

type Resolution struct {
	Width, Height int
}

func (r *Resolution) String() string {
	if r == nil {
		return ""
	}
	return fmt.Sprintf("%dx%d", r.Width, r.Height)
}

func GetPrimary() *Resolution {
	return getPrimary()
}
