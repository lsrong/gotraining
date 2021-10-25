package logger

import (
	"fmt"
	"os"
)

type ConsoleLogger struct {
	level int
}

// 初始化日志操作类
func NewConsoleLogger(config map[string]string) (*ConsoleLogger, error) {
	logLevel, ok := config["level"]
	if !ok {
		logLevel = DefaultLevel
	}

	level := GetLevelInt(logLevel)

	logger := &ConsoleLogger{
		level: level,
	}

	return logger, nil
}

// 写日志入口
func (c *ConsoleLogger) Log(level int, format string, args ...interface{}) {
	if c.level > level {
		return
	}
	// 日志数据
	data := LogData(level, format, args...)

	// 打印
	_, _ = fmt.Fprintf(os.Stdout, Format, data.Datetime, data.Level, data.File, data.Func, data.Line, data.Message)
}

// DEBUG 日志
func (c *ConsoleLogger) Debug(format string, args ...interface{}) {
	c.Log(DebugLevel, format, args...)
}

// TRACE 日志
func (c *ConsoleLogger) Trace(format string, args ...interface{}) {
	c.Log(DebugLevel, format, args...)
}

// INFO 日志
func (c *ConsoleLogger) Info(format string, args ...interface{}) {
	c.Log(DebugLevel, format, args...)
}

// WARNING 日志
func (c *ConsoleLogger) Warning(format string, args ...interface{}) {
	c.Log(DebugLevel, format, args...)
}

// NOTICE 日志
func (c *ConsoleLogger) Notice(format string, args ...interface{}) {
	c.Log(DebugLevel, format, args...)
}

// ERROR 日志
func (c *ConsoleLogger) Error(format string, args ...interface{}) {
	c.Log(DebugLevel, format, args...)
}

// FATAL 日志
func (c *ConsoleLogger) Fatal(format string, args ...interface{}) {
	c.Log(DebugLevel, format, args...)
}

func (c *ConsoleLogger) Close() {
}
