Logger
==========



This program is a high performing Go logging system as a CLI tool.
The log level defaults to `TRACE`, `DEBUG`, `INFO`, `WARN`, `ERROR`, similar to a terraform log level.
This log tracing logger system uses tracing level defaults that detect level traces and aggregates levels defaults based on telemetry of log throughput.

```console
$ go get github.com/Logger
```

## Example

```go
clilog.Env = "YOUR_APP_LOG"
clilog.SetOutput()

log.Printf("[INFO] run main function")
```

```console
$ YOUR_APP_LOG=trace go run _example/main.go
2020/01/24 20:49:35 [INFO] run main function
```

## License
MIT

