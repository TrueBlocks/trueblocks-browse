package app

import (
	"fmt"

	"github.com/TrueBlocks/trueblocks-browse/pkg/messages"
	sdk "github.com/TrueBlocks/trueblocks-sdk/v3"
)

func (a *App) GetExploreUrl(term string, google, dalle bool) string {
	opts := sdk.ExploreOptions{
		Terms:   []string{term},
		Google:  google,
		Dalle:   dalle,
		NoOpen:  true,
		Globals: a.globals,
	}

	if result, meta, err := opts.Explore(); err != nil {
		messages.Send(a.ctx, messages.Error, messages.NewErrorMsg(err))
		return ""
	} else if (result == nil) || (len(result) == 0) {
		messages.Send(a.ctx, messages.Error, messages.NewErrorMsg(fmt.Errorf("url not found")))
		return ""
	} else {
		a.meta = *meta
		return result[0].Url
	}
}
