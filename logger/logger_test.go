package logger

import (
	"testing"
)

func TestInitFile(t *testing.T) {
	config := map[string]string{
		"path":  "D:\\test",
		"level": "warning",
	}
	logger, _ := NewFileLogger(config)
	t.Log(logger)
	logger.Debug("debug log[%s]", "hello world")
	logger.Trace("Trace log")
	logger.Info("Info log")
	logger.Warning("Warning log")
	logger.Error("Error log")
	logger.Fatal("Fatal log")
	logger.Close()
}
