package rotatelogger

import "testing"

func TestNewLogger(t *testing.T) {
	logger := NewLogger("test", "", "debug")
	logger.Debug("debug")
	logger.Info("info")
	logger.Notice("notice")
	logger.Warning("warning")
	logger.Error("error")
	logger.Critical("critical")
}
