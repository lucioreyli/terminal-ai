# Terminal AI

It's a zero-dependency simple terminal-based interface for GPT using Golang.

## Features

- [x] Support for completions
- [x] Support for temperature
- [ ] Support for response formatting
- [ ] Support for images generation
- [ ] Support for multiple prompts
- [ ] Support for multiple files
- [ ] Support for threads
- [ ] Support for streaming

## SETUP
```sh
GPT_TOKEN='sk-...' # get it from https://platform.openai.com/account/api-keys

# if you're using fish
set -Ux GPT_TOKEN "sk-..."
```

## Usage

```bash
go run cmd/main.go 'What is cybertruck?' # default model is gpt-3.5-turbo

# check all available models at https://platform.openai.com/docs/models
go run cmd/main.go --model gpt-4o 'What is the meaning of life?'
```

### Build
```bash
go build -o bin/terminal_ai cmd/main.go
```

### Help
```bash
go run cmd/main.go -h # or ./bin/terminal_ai -h
go run cmd/main.go --help # or ./bin/terminal_ai --help
```

## License

MIT

## References
- [https://platform.openai.com/docs/api-reference/chat](https://platform.openai.com/docs/api-reference/chat)
- [https://platform.openai.com/docs/guides/text-generation](https://platform.openai.com/docs/guides/text-generation)
- [https://platform.openai.com/docs/models](https://platform.openai.com/docs/models)
