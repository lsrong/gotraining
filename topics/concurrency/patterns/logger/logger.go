package logger

import (
	"fmt"
	"io"
	"sync"
)

type Logger struct {
	write chan string
	wg    sync.WaitGroup
}

func New(w io.Writer, capacity int) *Logger {
	l := Logger{
		write: make(chan string, capacity),
	}

	l.wg.Add(1)

	// begin to write.
	go func() {
		for data := range l.write {
			_, err := fmt.Fprintln(w, data)
			if err != nil {
				fmt.Println(err)
			}
		}

		fmt.Println("shutdown Closing")

		l.wg.Done()
	}()

	return &l
}

func (l *Logger) Shutdown() {
	close(l.write)

	l.wg.Wait()
}

func (l *Logger) Write(data string) {
	// select 如果通道操作阻塞则会执行
	select {
	case l.write <- data:
	default:
		fmt.Println("Dropping the write")

	}
}
