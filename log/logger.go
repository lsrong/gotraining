package log

const (
	DEBUG_LEVEL = iota
	TRAVE_LEVEL
	INFO_LEVEL
	WARNING_LEVEL
	ERROR_LEVEL
	FATAL_LEVEL
)

// 定义日志驱动接口
type logger interface {
	Init()
	SetLevel(level int)
	Debug(format string, args ...interface{})
	Trace(format string, args ...interface{})
	Info(format string, args ...interface{})
	Warning(format string, args ...interface{})
	Error(format string, args ...interface{})
	Fatal(format string, args ...interface{})
	Close()
}
