package app

import "errors"

var (
	ErrLoadingNames    = errors.New("error loading names")
	ErrLoadingRpc      = errors.New("error loading rpc")
	ErrWindowSize      = errors.New("error fixing window size")
	ErrDaemonLoad      = errors.New("could not load daemon")
	ErrSavingProject   = errors.New("error saving project file")
	ErrOpeningProject  = errors.New("error opening file")
	ErrLoadingProject  = errors.New("error loading file")
	ErrProjectNotSaved = errors.New("project not saved")
)
