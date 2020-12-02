package logger

import (
	"errors"
	"fmt"
	"os"
	"time"
)

const (
	DefaultLevel = "debug"
)

type data struct {
	Message string `json:"message"`
	File    string `json:"file"`
	Line    string `json:"line"`
	Func    string `json:"func"`
	Level   string `json:"level"`
}

type FileLogger struct {
	file  *os.File
	path  string
	level int
	//data  chan *data
}

func NewFileLogger(config map[string]string) (*FileLogger, error) {
	// 日志路径
	path, ok := config["path"]
	if !ok {
		err := errors.New("Empty path config")

		return nil, err
	}

	// 日志级别
	levelConfig, ok := config["level"]
	if !ok {
		levelConfig = DefaultLevel
	}
	level := GetLevelInt(levelConfig)

	log := &FileLogger{
		level: level,
		path:  path,
	}

	err := log.Init()

	if err != nil {
		return nil, errors.New("Failed to init file logger")
	}

	return log, nil
}

func (f *FileLogger) Init() error {
	filename := fmt.Sprintf("%s/golang-%s.log", f.path, time.Now().Format("2006-01-02"))
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		return fmt.Errorf("Failed to open log file[%s]", filename)
	}
	f.file = file
	return nil
}

func (f *FileLogger) checkLevel(level int) bool {
	if f.level <= level {
		return false
	}
	return true
}

func (f *FileLogger) Write(level int, format string, args ...interface{}) {
	filename, function, line := Info()
	datetime := time.Now().Format("2006-01-02 15:04:05")
	levelStr := GetLevelStr(level)
	message := fmt.Sprintf(format, args...)
	// fmt.Fprintf(file, "%s %s (%s:%s:%d) %s\n", nowStr, levelStr, fileName, funcName, lineNo, msg)

	_, _ = f.file.WriteString(fmt.Sprintf("%s %s (%s:%s:%d) %s\n", datetime, levelStr, filename, function, line, message))
}

func (f *FileLogger) Debug(format string, args ...interface{}) {
	if f.checkLevel(DebugLevel) {
		return
	}
	f.Write(DebugLevel, format, args...)
}

func (f *FileLogger) Trace(format string, args ...interface{}) {
	if f.checkLevel(TraceLevel) {
		return
	}
	f.Write(TraceLevel, format, args...)
}

func (f *FileLogger) Info(format string, args ...interface{}) {
	if f.checkLevel(InfoLevel) {
		return
	}
	f.Write(InfoLevel, format, args...)
}

func (f *FileLogger) Warning(format string, args ...interface{}) {
	if f.checkLevel(WarningLevel) {
		return
	}

	f.Write(WarningLevel, format, args...)
}
func (f *FileLogger) Error(format string, args ...interface{}) {
	if f.checkLevel(ErrorLevel) {
		return
	}
	f.Write(ErrorLevel, format, args...)
}

func (f *FileLogger) Fatal(format string, args ...interface{}) {
	if f.checkLevel(FatalLevel) {
		return
	}

	f.Write(FatalLevel, format, args...)
}

func (f *FileLogger) Close() {
	f.file.Close()
}
