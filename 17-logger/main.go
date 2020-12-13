package main

import "github.com/learning_golang/logger"

func main() {
	config := map[string]string{
		"path":  "D:\\test",
		"level": "debug",
	}
	file, _ := logger.NewFileLogger(config)

	file.Debug("debug log[%s]", "hello world")
	file.Trace("Trace log")
	file.Info("Info log")
	file.Warning("Warning log")
	file.Error("Error log")
	file.Fatal("Fatal log")

	//defer file.Close()
}
