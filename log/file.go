package log

import (
	"os"
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
	perm  os.FileMode
	path  string
	name  string
	level int
	data  chan *data
}

func New(config map[string]string) (log logger, e error) {
	log = &FileLogger{}

	return
}

func (f *FileLogger) Init() {

}
func (f *FileLogger) SetLevel(level int) {

}

func (f *FileLogger) Debug(format string, args ...interface{}) {

}

func (f *FileLogger) Trace(format string, args ...interface{}) {

}

func (f *FileLogger) Info(format string, args ...interface{}) {

}

func (f *FileLogger) Warning(format string, args ...interface{}) {

}
func (f *FileLogger) Error(format string, args ...interface{}) {

}

func (f *FileLogger) Fatal(format string, args ...interface{}) {

}

func (f *FileLogger) Close() {

}
