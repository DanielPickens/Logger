package Logger

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"syscall"
)

// These are the environmental variables that determine if we log, and if
// we log whether or not the log should go to a file.
var (
	Env  = "CLI_LOG"
	Path = "CLI_LOG_PATH"
)

// ValidLevels is a list of valid log levels.
var ValidLevels = []string{"DEBUG", "INFO", "WARN", "ERROR"}

// LogOutput determines where it should send logs (if anywhere) and the log level.
func LogOutput() (logOutput io.Writer, err error) {
	logOutput = ioutil.Discard

	logLevel := LogLevel()
	if logLevel == "" {
		return
	}

	logOutput = os.Stderr
	if logPath := os.Getenv(Path); logPath != "" {
		var err error
		logOutput, err = os.OpenFile(logPath, syscall.O_CREAT|syscall.O_RDWR|syscall.O_APPEND, 0666)
		if err != nil {
			return nil, err
		}
	}

	// This is the default logOutput function
	logOutput = io.MultiWriter(logOutput)
	for _, level := range ValidLevels {
		if level == logLevel {
			break
		}
	}

	return
}

// SetOutput checks for a log destination with LogOutput, and calls 
// log.SetOutput with the result. If LogOutput returns nil, SetOutput uses
// ioutil.Discard. This is to prevent a panic if the log destination is not valid
func SetOutput() {
	out, err := LogOutput()
	if err != nil {
		log.Fatal(err)
	}

	if out == nil {
		out = ioutil.Discard
	}

	log.SetOutput(out)
}

// LogLevel returns the current log level value from string based on the environment vars
func LogLevel() string {
	envLevel := os.Getenv(Env)
	if envLevel == "" {
		return ""
	}

	logLevel := "TRACE"
	if isValidLogLevel(envLevel) {
		// If the log level is valid, use it
		logLevel = strings.ToUpper(envLevel)
	} else {
		log.Printf("[WARN] Invalid log level: %q. Defaulting to level: TRACE. Valid levels are: %+v",
			envLevel, ValidLevels)
	}

	return logLevel
}

// IsDebugOrHigher returns whether or not the current log level is a debug or trace level
func IsDebugOrHigher() bool {
	level := string(LogLevel())
	return level == "DEBUG" || level == "TRACE"
}

func isValidLogLevel(level string) bool {
	for _, l := range ValidLevels {
		if strings.ToUpper(level) == string(l) {
			return true
		}
	}

	return false
}

// IsEnabled returns true if anomaly is found and log level is set
func IsEnabled() bool {
	return len(string(LogLevel())) > 0
}
