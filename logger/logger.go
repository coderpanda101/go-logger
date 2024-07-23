package logger

import (
	"fmt"
	"log"
	"os"
)

var (
	LogLevel uint8
)

// Logger interface with different log levels
type Logger interface {
	Dev(v ...interface{})   // 5
	Debug(v ...interface{}) // 4
	Info(v ...interface{})  // 3
	Warn(v ...interface{})  // 2
	Error(v ...interface{}) // 1
}

// CustomLogger struct to handle both console and file logging
type CustomLogger struct {
	dev   *log.Logger // 5
	debug *log.Logger // 4
	info  *log.Logger // 3
	warn  *log.Logger // 2
	error *log.Logger // 1
}

// Initializes and returns a CustomLogger
func GetLogger(path string, level uint8) (*CustomLogger, error) {
	if level > 5 || level < 1 {
		fmt.Println("Invalid log level, should be in range from 1 to 5, using default level: 4")
		LogLevel = 4
	} else {
		LogLevel = level
	}

	// Open the log file in append mode, create if not exists
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("**** Error initiating Logger!! ****", err)
		return nil, err
	}

	return &CustomLogger{
		dev:   log.New(file, "[DEV] ", log.LstdFlags),   // 5
		debug: log.New(file, "[DEBUG] ", log.LstdFlags), // 4
		info:  log.New(file, "[INFO] ", log.LstdFlags),  // 3
		warn:  log.New(file, "[WARN] ", log.LstdFlags),  // 2
		error: log.New(file, "[ERROR] ", log.LstdFlags), // 1
	}, nil
}

// Dev logs dev messages to file and console, to be used only for devlopment logs
func (l *CustomLogger) Dev(v ...interface{}) {
	if LogLevel > 4 {
		log.New(os.Stdout, "DEVELOPMENT: ", log.LstdFlags).Println(v...)
		l.debug.Println(v...)
	}
}

// Debug logs debug messages to file and console
func (l *CustomLogger) Debug(v ...interface{}) {
	if LogLevel > 3 {
		log.New(os.Stdout, "DEBUG: ", log.LstdFlags).Println(v...)
		l.debug.Println(v...)
	}
}

// Info logs info messages to file and console
func (l *CustomLogger) Info(v ...interface{}) {
	if LogLevel > 2 {
		log.New(os.Stdout, "INFO: ", log.LstdFlags).Println(v...)
		l.info.Println(v...)
	}
}

// Warn logs warn messages to file and console
func (l *CustomLogger) Warn(v ...interface{}) {
	if LogLevel > 1 {
		log.New(os.Stdout, "WARN: ", log.LstdFlags).Println(v...)
		l.warn.Println(v...)
	}
}

// Error logs error messages to file and console
func (l *CustomLogger) Error(v ...interface{}) {
	if LogLevel >= 1 {
		log.New(os.Stderr, "ERROR: ", log.LstdFlags).Println(v...)
		l.error.Println(v...)
	}
}
