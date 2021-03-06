Logger
==========



Logger is a high performing Go logging system that provides log trace level detection as a CLI.
The log level defaults to `TRACE`, `DEBUG`, `INFO`, `WARN`, `ERROR`, similar to a terraform log level and detects level traces from aggregated level defaults based on telemetry of log throughput.

```console
$ go get github.com/DanielPickens/Logger
```

## Example

```go
log.Env = "YOUR_APP_LOG"
log.SetOutput()

log.Printf("[INFO] your running logs")
```

```console
$ YOUR_APP_LOG=trace go run main.go
2020/01/24 20:49:35 [INFO] your running logs
```


## Logging Structure 
There are many forms of logging and many ways to record logs. They can be written to a file, sent to stdout/stderr, sent to a logging service, and more. Each of these is possible with this cli.

If you have a specific logger you want to use, you can write one that conforms to the Logger cli found in log.go. That logger can then be configured as documented in the previous section.


## Log Levels
The following log levels are currently active in log tracing defaults across the interface.

- Error
- Warn
- Info
- Debug
- Trace

## License
MIT


