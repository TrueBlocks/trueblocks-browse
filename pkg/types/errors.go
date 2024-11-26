package types

import "errors"

var (
	ErrNoConfigFolder     = errors.New("core config folder not found")
	ErrNoConfigFile       = errors.New("trueBlocks.toml file not found")
	ErrCantReadToml       = errors.New("can't read toml file")
	ErrChainNotConfigured = errors.New("chain not configured")
)
