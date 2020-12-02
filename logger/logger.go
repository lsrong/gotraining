package logger

const (
	DebugLevel = iota
	TraceLevel
	InfoLevel
	WarningLevel
	ErrorLevel
	FatalLevel
)

// 定义日志驱动接口
type loggerInterface interface {
	Write(level int, format string, args ...interface{})
	Debug(format string, args ...interface{})
	Trace(format string, args ...interface{})
	Info(format string, args ...interface{})
	Warning(format string, args ...interface{})
	Error(format string, args ...interface{})
	Fatal(format string, args ...interface{})
	Close()
}
