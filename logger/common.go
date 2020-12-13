package logger

import (
	"fmt"
	"path"
	"runtime"
	"strings"
	"time"
)

// 日志级别
func GetLevelInt(level string) int {
	level = strings.ToLower(level)
	switch level {
	case "debug":
		return DebugLevel
	case "trace":
		return TraceLevel
	case "info":
		return InfoLevel
	case "warning":
		return WarningLevel
	case "error":
		return ErrorLevel
	case "fatal":
		return FatalLevel
	}
	return DebugLevel
}

// 级别名称
func GetLevelStr(level int) string {
	switch level {
	case DebugLevel:
		return "Debug"
	case TraceLevel:
		return "Trace"
	case InfoLevel:
		return "Info"
	case WarningLevel:
		return "Warning"
	case ErrorLevel:
		return "Error"
	case FatalLevel:
		return "Fatal"
	}
	return "Debug"
}

// 日志位置信息
func LogPosition() (filename string, function string, line int) {
	pc, file, line, ok := runtime.Caller(4)
	if !ok {
		return "", "", 0
	}

	filename = path.Base(file)
	function = path.Base(runtime.FuncForPC(pc).Name())

	return filename, function, line
}

// 构建日志数据
func LogData(level int, format string, args ...interface{}) *Data {
	file, function, line := LogPosition()
	datetime := time.Now().Format("2006-01-02 15:04:05")
	levelStr := GetLevelStr(level)
	message := fmt.Sprintf(format, args...)

	return &Data{
		Level:    levelStr,
		Datetime: datetime,
		Message:  message,
		File:     file,
		Line:     line,
		Func:     function,
	}
}
