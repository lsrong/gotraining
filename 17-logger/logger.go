package _7_logger

type Logger interface {
	Debug(message string)
	Trace(message string)
	Info(message string)
	Warn(message string)
	Error(message string)
	Fatal(message string)
}
