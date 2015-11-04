package rotatelogger

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/op/go-logging"
)

const (
	CRITICAL logging.Level = logging.CRITICAL
	ERROR                  = logging.ERROR
	WARNING                = logging.WARNING
	NOTICE                 = logging.NOTICE
	INFO                   = logging.INFO
	DEBUG                  = logging.DEBUG
)

var consoleFormat = logging.MustStringFormatter(
	"[%{color}%{level:s}%{color:reset}][%{time:2006-01-02 15:04:05.000} %{shortfile}: %{longfunc}][%{id}] %{message}",
)

var fileFormat = logging.MustStringFormatter(
	"[%{level:s}][%{time:2006-01-02 15:04:05.000} %{shortfile}: %{longfunc}][%{id}] %{message}",
)

func createConsoleBackend(level logging.Level) logging.LeveledBackend {
	consoleBackend := logging.NewLogBackend(os.Stderr, "", 0)
	consoleBackendFormatter := logging.NewBackendFormatter(consoleBackend, consoleFormat)
	consoleBackendLeveled := logging.AddModuleLevel(consoleBackendFormatter)
	consoleBackendLeveled.SetLevel(level, "")
	return consoleBackendLeveled
}

func createFileBackend(fullFileName string, level logging.Level) logging.LeveledBackend {
	file := &Rotator{}
	file.Create(fullFileName, HourlyRotation)
	fileBackend := logging.NewLogBackend(file, "", 0)
	fileBackendFormatter := logging.NewBackendFormatter(fileBackend, fileFormat)
	fileBackendLeveled := logging.AddModuleLevel(fileBackendFormatter)
	fileBackendLeveled.SetLevel(level, "")
	return fileBackendLeveled
}

func NewLogger(module string, logDir string, level string) *logging.Logger {
	var logLevel logging.Level
	switch strings.ToLower(level) {
	case "debug":
		logLevel = logging.DEBUG
	case "info":
		logLevel = logging.INFO
	case "notice":
		logLevel = logging.NOTICE
	case "warning":
		logLevel = logging.WARNING
	case "error":
		logLevel = logging.ERROR
	case "critical":
		logLevel = logging.CRITICAL
	default:
		logLevel = logging.DEBUG
	}
	logger := logging.MustGetLogger(module)
	backends := make([]logging.Backend, 0)
	backends = append(backends, createConsoleBackend(logLevel))
	if strings.TrimSpace(logDir) != "" {
		if absPath, err := filepath.Abs(filepath.Join(logDir, module+".log")); err == nil {
			backends = append(backends, createFileBackend(absPath, logLevel))
			fmt.Println("log file:" + absPath)
		} else {
			fmt.Println("fail to use file log: " + err.Error())
		}
	}
	logger.SetBackend(logging.MultiLogger(backends...))
	return logger
}
