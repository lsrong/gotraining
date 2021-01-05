package logger

import (
	"testing"
)

// 测试文件日志
func TestFileLogger(t *testing.T) {
	config := map[string]string{
		"path":  "/Users/lsrong/Work/Project/Test",
		"level": "debug",
	}
	file, _ := NewFileLogger(config)
	//t.Log(logger)
	file.Debug("debug log[%s]", "hello world")
	file.Trace("Trace log")
	file.Info("Info log")
	file.Warning("Warning log")
	file.Error("Error log")
	file.Fatal("Fatal log")
	file.Close()
}

// 测试终端日志
func TestConsoleLogger(t *testing.T) {
	config := map[string]string{
		"level": "debug",
	}

	console, _ := NewConsoleLogger(config)
	console.Debug("debug log[%s]", "hello world")
	console.Trace("Trace log")
	console.Info("Info log")
	console.Warning("Warning log")
	console.Error("Error log")
	console.Fatal("Fatal log")
}
