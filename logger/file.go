package logger

import (
	"errors"
	"fmt"
	"os"
	"time"
)

const ChanNum = 10000

type FileLogger struct {
	file  *os.File
	path  string
	level int
	data  chan *Data
}

// 构造文件日志处理类
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
		data:  make(chan *Data, ChanNum),
	}

	err := log.Init()

	if err != nil {
		return nil, errors.New("Failed to init file logger")
	}

	return log, nil
}

// 初始化操作
func (f *FileLogger) Init() error {
	filename := fmt.Sprintf("%s/golang-%s.log", f.path, time.Now().Format("2006-01-02"))
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		return fmt.Errorf("Failed to open log file[%s]", filename)
	}
	f.file = file

	// 后台写入
	go f.WriteBackground()
	return nil
}

// 文件统一写入入口
func (f *FileLogger) Log(level int, format string, args ...interface{}) {
	if f.level > level {
		return
	}
	// 日志数据
	data := LogData(level, format, args...)

	// 日志格式：fmt.Fprintf(file, "%s %s (%s:%s:%d) %s\n", nowStr, levelStr, fileName, funcName, lineNo, msg)
	_, _ = f.file.WriteString(fmt.Sprintf(Format, data.Datetime, data.Level, data.File, data.Func, data.Line, data.Message))

	// 放入日志数据管道
	//select {
	//case f.data <- data:
	//default:
	//}
}

// TODO 后台协程应用
func (f *FileLogger) WriteBackground() {
	for data := range f.data {
		_, _ = f.file.WriteString(fmt.Sprintf(Format, data.Datetime, data.Level, data.File, data.Func, data.Line, data.Message))
	}
}

// DEBUG 日志
func (f *FileLogger) Debug(format string, args ...interface{}) {
	f.Log(DebugLevel, format, args...)
}

// TRACE 日志
func (f *FileLogger) Trace(format string, args ...interface{}) {
	f.Log(TraceLevel, format, args...)
}

// INFO 日志
func (f *FileLogger) Info(format string, args ...interface{}) {
	f.Log(InfoLevel, format, args...)
}

// WARNING 日志
func (f *FileLogger) Warning(format string, args ...interface{}) {
	f.Log(WarningLevel, format, args...)
}

// ERROR 日志
func (f *FileLogger) Error(format string, args ...interface{}) {
	f.Log(ErrorLevel, format, args...)
}

// FATAL 日志
func (f *FileLogger) Fatal(format string, args ...interface{}) {
	f.Log(FatalLevel, format, args...)
}

// 关闭文件句柄
func (f *FileLogger) Close() {
	f.file.Close()
}
