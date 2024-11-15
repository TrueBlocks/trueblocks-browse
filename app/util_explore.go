package app

import (
	"fmt"
	"os"

	sdk "github.com/TrueBlocks/trueblocks-sdk/v3"
)

func (a *App) GetExploreUrl(term string, google, dalle bool) string {
	if len(term) == 0 {
		return ""
	} else if len(term) != 42 {
		google = false
		dalle = false
	}

	opts := sdk.ExploreOptions{
		Terms:   []string{term},
		Google:  google,
		Dalle:   dalle,
		NoOpen:  true,
		Globals: a.getGlobals(),
	}

	// TODO: Expose this to the user and/or put it in trueBlocks.toml
	os.Setenv("TB_DALLE_SERIES", "five-tone-postal-protozoa")
	if result, meta, err := opts.Explore(); err != nil {
		a.emitErrorMsg(err, nil)
		return ""
	} else if (result == nil) || (len(result) == 0) {
		err := fmt.Errorf("url not found")
		a.emitErrorMsg(err, nil)
		return ""
	} else {
		a.meta = *meta
		return result[0].Url
	}
}
