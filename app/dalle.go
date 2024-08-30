package app

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
)

func (a *App) GetDalle() (string, error) {
	addr := a.GetLastSub("/history")
	url := "http://192.34.63.136:8080/dalle/simple/" + addr

	resp, err := http.Get(url)
	if err != nil {
		logger.Info("Failed to fetch content:", err)
		return "", err
	}
	defer resp.Body.Close()

	contentType := resp.Header.Get("Content-Type")
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Info("Failed to read response body:", err)
		return "", err
	}

	if contentType == "image/png" || contentType == "image/jpeg" {
		encodedImage := base64.StdEncoding.EncodeToString(bytes)
		return fmt.Sprintf("data:%s;base64,%s", contentType, encodedImage), nil
	}

	return string(bytes), nil
}
