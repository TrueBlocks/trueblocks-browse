# trueblocks-browse

Web 3.0 Account Browser built on TrueBlocks

## Installing

```[bash]
go mod tidy
cd frontend && yarn install && cd -
```

## Running

```[bash]
wails dev
```

## Building

```[bash]
wails build
```

## Api Keys

If you intend to use features that require OpenAI, rename the `.env.example` file to `.env` and add your OpenAI API key. The features will not work otherwise.
