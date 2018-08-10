package log

import (
	"fmt"
	"log"
	"time"
)

// Logger with level, output logs to std output

var (
	// LevelFatal fatal 1
	LevelFatal = 1

	// LevelError error 2
	LevelError = 2

	// LevelWarn warning 3
	LevelWarn = 3

	// LevelInfo info 4
	LevelInfo = 4

	// LevelDebug debug 5
	LevelDebug = 5

	// default log level is info
	//level = LevelInfo
	level = LevelDebug
)

// SetLevel set log level
func SetLevel(l int) {
	level = l
}

// GetLevel get log level
func GetLevel() int {
	return level
}

// log base log output function
func _log(l int, format string, a ...interface{}) {
	prefix := []string{
		"",
		"FATAL",
		"ERROR",
		"WARN ",
		"INFO ",
		"DEBUG",
	}
	if l <= level {
		str := fmt.Sprintf(format, a...)
		if l == LevelFatal {
			log.Fatalf("[%s] %s", prefix[l], str)
		} else {
			log.Printf("[%s] %s", prefix[l], str)
		}
	}
}

// Fatal log fatal message
func Fatal(format string, a ...interface{}) {
	_log(LevelFatal, format, a...)
}

// Error log error message
func Error(format string, a ...interface{}) {
	_log(LevelError, format, a...)
}

// Warn log warning message
func Warn(format string, a ...interface{}) {
	_log(LevelWarn, format, a...)
}

// Info log info message
func Info(format string, a ...interface{}) {
	_log(LevelInfo, format, a...)
}

// Debug log debug message
func Debug(format string, a ...interface{}) {
	_log(LevelDebug, format, a...)
}

// Profiling print profile log, valid in debug level
// Should use defer call at the beginning of function
// Example:
// func abc() {
//	  st := time.Now()
//	  defer logger.Profile(st, xxx, ...)
//	  ...
// }
func Profiling(begin time.Time, format string, a ...interface{}) {
	if level >= LevelDebug {
		//_log(LevelDebug, "[% 5d ms]%s", (time.Now().UnixNano()-begin.UnixNano())/int64(time.Millisecond), fmt.Sprintf(format, a...))
		_log(LevelDebug, "[Time Elapsed %.04fs]%s", time.Since(begin).Seconds(), fmt.Sprintf(format, a...))
	}
}
