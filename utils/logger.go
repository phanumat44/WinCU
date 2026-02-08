package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"time"
)

// LogLevel definition
type LogLevel int

const (
	INFO LogLevel = iota
	WARN
	ERROR
	DEBUG
)

// LogEntry represents a single log line in JSON format
type LogEntry struct {
	Timestamp string `json:"timestamp"`
	Level     string `json:"level"`
	Message   string `json:"message"`
	Data      any    `json:"data,omitempty"`
}

var (
	jsonOutput bool
	mu         sync.Mutex
)

// InitLogger configures the logger
func InitLogger(jsonFormat bool) {
	mu.Lock()
	defer mu.Unlock()
	jsonOutput = jsonFormat
}

func log(level LogLevel, msg string, data any) {
	mu.Lock()
	defer mu.Unlock()

	var levelStr string
	switch level {
	case ERROR:
		levelStr = "❌"
	case WARN:
		levelStr = "⚠️"
	case DEBUG:
		levelStr = "🔍"
	default:
		levelStr = "✅"
	}

	if jsonOutput {
		entry := LogEntry{
			Timestamp: time.Now().Format(time.RFC3339),
			Level:     levelStr,
			Message:   msg,
			Data:      data,
		}
		encoder := json.NewEncoder(os.Stdout)
		encoder.Encode(entry)
	} else {
		// Standard text format
		// [INFO] Message data...
		if data != nil {
			fmt.Printf("%s %s %v\n", levelStr, msg, data)
		} else {
			fmt.Printf("%s %s\n", levelStr, msg)
		}
	}
}

// Info logs an informational message
func Info(msg string, data ...any) {
	var d any
	if len(data) > 0 {
		d = data[0]
	}
	log(INFO, msg, d)
}

// Warn logs a warning message
func Warn(msg string, data ...any) {
	var d any
	if len(data) > 0 {
		d = data[0]
	}
	log(WARN, msg, d)
}

// Error logs an error message
func Error(msg string, data ...any) {
	var d any
	if len(data) > 0 {
		d = data[0]
	}
	log(ERROR, msg, d)
}

// Debug logs a debug message
func Debug(msg string, data ...any) {
	var d any
	if len(data) > 0 {
		d = data[0]
	}
	log(DEBUG, msg, d)
}
