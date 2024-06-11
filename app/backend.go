package app

import "fmt"

func (a *App) GetNames(page int) []string {
	first := page
	last := first + 20
	if len(a.namesArray) < last {
		return []string{"No names loaded"}
	}
	n := a.namesArray[first:last]
	var ret []string
	for _, name := range n {
		ret = append(ret, fmt.Sprintf("%s: %s", name.Address.Hex(), name.Name))
	}
	return ret
}
