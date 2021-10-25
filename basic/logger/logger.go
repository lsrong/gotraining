package logger

const (
	DebugLevel = iota
	TraceLevel
	InfoLevel
	WarningLevel
	ErrorLevel
	FatalLevel
)

const (
	Format       string = "%s %s (%s:%s:%d) %s\n"
	DefaultLevel string = "debug"
)

type Data struct {
	Level    string `json:"level"`
	Datetime string `json:"datetime"`
	Message  string `json:"message"`
	File     string `json:"file"`
	Line     int    `json:"line"`
	Func     string `json:"func"`
}

// 定义日志驱动接口
type loggerInterface interface {
	Log(level int, format string, args ...interface{})
	Debug(format string, args ...interface{})
	Trace(format string, args ...interface{})
	Info(format string, args ...interface{})
	Warning(format string, args ...interface{})
	Error(format string, args ...interface{})
	Fatal(format string, args ...interface{})
	Close()
}
