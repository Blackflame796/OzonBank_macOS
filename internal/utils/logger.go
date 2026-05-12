package utils

import (
	"log"
	"os"
)

// Logger provides basic logging functionality
type Logger struct {
	info  *log.Logger
	warn  *log.Logger
	err   *log.Logger
	debug *log.Logger
}

// NewLogger creates a new logger instance
func NewLogger() *Logger {
	return &Logger{
		info:  log.New(os.Stdout, "[INFO] ", log.LstdFlags),
		warn:  log.New(os.Stdout, "[WARN] ", log.LstdFlags),
		err:   log.New(os.Stderr, "[ERROR] ", log.LstdFlags),
		debug: log.New(os.Stdout, "[DEBUG] ", log.LstdFlags),
	}
}

// Info logs an info message
func (l *Logger) Info(msg string) {
	l.info.Println(msg)
}

// Warn logs a warning message
func (l *Logger) Warn(msg string) {
	l.warn.Println(msg)
}

// Error logs an error message
func (l *Logger) Error(msg string) {
	l.err.Println(msg)
}

// Debug logs a debug message
func (l *Logger) Debug(msg string) {
	l.debug.Println(msg)
}
