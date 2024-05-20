package logging

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

const (
	debugLevel = iota
	infoLevel
	warningLevel
	errorLevel
)

var (
	Logger *log.Logger
	logLevel = infoLevel // Default log level
)

func init() {
	// Create a new logger instance that writes to a log file
	logFile, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}

	Logger = log.New(io.MultiWriter(os.Stdout, logFile), "", log.LstdFlags|log.Lshortfile)
}

func SetLogLevel(level int) {
    logLevel = level
}

func Debug(format string, args ...interface{}) {
	if logLevel <= debugLevel {
		logWithContext(debugLevel, format, args...)
	}
}

func Info(format string, args ...interface{}) {
	if logLevel <= infoLevel {
		logWithContext(infoLevel, format, args...)
	}
}

func Warning(format string, args ...interface{}) {
	if logLevel <= warningLevel {
		logWithContext(warningLevel, format, args...)
	}
}

func Error(format string, args ...interface{}) {
	if logLevel <= errorLevel {
		logWithContext(errorLevel, format, args...)
	}
}

func logWithContext(level int, format string, args ...interface{}) {
	prefix := getLogLevelPrefix(level)
	_, file, line, _ := runtime.Caller(2)
	callerWithinPackage := getCallerWithinPackage(file)
	message := fmt.Sprintf(format, args...)
	Logger.Printf("%s %s:%d %s", prefix, callerWithinPackage, line, message)
}

func getLogLevelPrefix(level int) string {
	switch level {
	case debugLevel:
		return "DEBUG"
	case infoLevel:
		return "INFO"
	case warningLevel:
		return "WARNING"
	case errorLevel:
		return "ERROR"
	default:
		return "UNKNOWN"
	}
}

func getCallerWithinPackage(file string) string {
	_, filename := filepath.Split(file)
	packageParts := strings.Split(filename, "/")
	if len(packageParts) > 1 {
		return packageParts[len(packageParts)-2] + "/" + packageParts[len(packageParts)-1]
	}
	return filename
}