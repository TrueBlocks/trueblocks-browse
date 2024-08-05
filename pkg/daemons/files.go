package daemons

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
)

type DaemonFile struct {
	Port   int `json:"port"`
	Daemon `json:"daemon"`
}

func NewFileDaemon(name string, port int, sleep time.Duration) *DaemonFile {
	return &DaemonFile{
		Port: port,
		Daemon: Daemon{
			Name:    name,
			Sleep:   sleep,
			Color:   "green",
			State:   Paused,
			Started: time.Now(),
		},
	}
}

func (s *DaemonFile) Run() {
	logger.Info(fmt.Sprintf("Serving files from (%s): %d\n", s.Name, s.Port))

	http.HandleFunc("/files/", func(w http.ResponseWriter, r *http.Request) {
		path := strings.TrimPrefix(r.URL.Path, "/files/")
		parts := strings.Split(path, "&")
		series, address := "", ""
		if len(parts) > 2 {
			series = parts[0]
			address = parts[1]
		}
		if series == "" || address == "" {
			http.Error(w, "Series or address not provided to file daemon", http.StatusBadRequest)
			return
		}
		cwd, err := os.Getwd()
		if err != nil {
			http.Error(w, "Error getting current working directory", http.StatusInternalServerError)
			return
		}
		filePath := filepath.Join(cwd, "output", series, "annotated", address+".png")
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			msg := fmt.Sprintf("File not found at %s", filePath)
			http.Error(w, msg, http.StatusNotFound)
			return
		}
		s.Daemon.Notify(filePath)

		w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
		w.Header().Set("Pragma", "no-cache")
		w.Header().Set("Expires", "0")
		http.ServeFile(w, r, filePath)
	})

	if err := http.ListenAndServe(fmt.Sprintf(":%d", s.Port), nil); err != nil {
		logger.Error("File daemon error:", err)
	}
}

func (s *DaemonFile) Stop() error {
	return s.Daemon.Stop()
}

func (s *DaemonFile) Pause() error {
	return s.Daemon.Pause()
}

func (s *DaemonFile) Toggle() error {
	return s.Daemon.Toggle()
}

func (s *DaemonFile) Tick() int {
	return s.Daemon.Tick()
}
