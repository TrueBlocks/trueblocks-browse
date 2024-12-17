all:
	@wails build

build: all
	@open build/bin/TrueBlocks\ Browse.app/Contents/MacOs/TrueBlocks\ Browse


lint:
	@cd frontend && yarn lint && cd -

update:
	@go get github.com/TrueBlocks/trueblocks-sdk/v4@latest
	@go get github.com/TrueBlocks/trueblocks-core/src/apps/chifra@latest

open:
	@open build/bin/TrueBlocks\ Browse.app/Contents/MacOs/TrueBlocks\ Browse

run:
	@yarn dev

generate:
	@echo building goMaker...
	@cd ~/Development/trueblocks-core/build ; make -j 12 goMaker ; cd - && goMaker

generate-go:
	@echo building goMaker...
	@cd ~/Development/trueblocks-core/build ; make -j 12 goMaker ; cd - && TB_GENERATOR_FILTER=go.tmpl goMaker
