Logger
==========



This program is a high performing Go logging system that provides log trace level detection as a CLI tool.
The log level defaults to `TRACE`, `DEBUG`, `INFO`, `WARN`, `ERROR`, similar to a terraform log level.
The application detects level traces and aggregates levels defaults based on telemetry of log throughput.

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


## Logging methods
There are many loggers and many ways to record logs. They can be written to a file, sent to stdout/stderr, sent to a logging service, and more. Each of these is possible with this package.

If you have a specific logger you want to use, you can write one that conforms to the Logger interface found in logging.go. That logger can then be configured as documented in the previous section.


## Log Levels
The following log levels are currently active in log tracing defaults across the interface.

- Error
- Warn
- Info
- Debug
- Trace

## License
MIT


## Todo

 - Needs to write messages that in some situations are sent to console output and in other situations are sent to logs. 
 
