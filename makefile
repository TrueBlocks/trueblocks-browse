all:
	@wails build

build: all
	@open build/bin/TrueBlocks\ Browse.app/Contents/MacOs/TrueBlocks\ Browse

update:
	@go get github.com/TrueBlocks/trueblocks-core/sdk/v3@latest
	@go get github.com/TrueBlocks/trueblocks-core/src/apps/chifra@latest

open:
	@open build/bin/TrueBlocks\ Browse.app/Contents/MacOs/TrueBlocks\ Browse

run:
	@yarn dev

generate:
	@goMaker

