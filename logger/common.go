package logger

import "runtime"

// 获取日志级别
/**
DebugLevel
TraceLevel
InfoLevel
WarningLevel
ErrorLevel
FatalLevel
*/
func GetLevelInt(level string) int {
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

func Info() (filename string, function string, line int) {
	pc, file, line, ok := runtime.Caller(4)
	if !ok {
		return "", "", 0
	}

	filename = file
	function = runtime.FuncForPC(pc).Name()

	return filename, function, line
}
