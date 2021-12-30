package blocking

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
	"testing"
	"time"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

// 实现大数据文件阻塞任务
var data []byte

// 初始化读取大文件。
func init() {
	file, err := os.Open("data.bytes")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	data, err = io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("data.Bytes: %d \n", len(data))
}

// TestLatency 测试单个延迟
func TestLatency(t *testing.T) {
	bufSize := 0
	dur := stream(bufSize)

	t.Logf("bufSuize: %d, second: %.2fs", bufSize, float64(dur)/float64(time.Second))
}

// TestLatencies 演示不同有缓冲channel与无缓冲的阻塞提升百分比曲线
func TestLatencies(t *testing.T) {
	// 阻塞分布曲线图
	var bufSize int
	var first time.Duration
	var idx int
	xys := make(plotter.XYs, 20)
	for {
		since := stream(bufSize)
		if bufSize == 0 {
			first = since
		}
		dec := (float64(first) - float64(since)) / float64(first) * 100
		xys[idx].X = float64(bufSize)
		xys[idx].Y = dec
		idx++
		t.Logf("BufSize: %d\t %v\t %.2f%%\n", bufSize, since, dec)
		if bufSize < 10 {
			bufSize++
			continue
		}
		if bufSize == 100 {
			break
		}

		bufSize += 10
	}

	err := makePlot(xys)
	if err != nil {
		t.Error(err)
	}
}

// stream 启动一个goroutine模拟读取文件。
func stream(bufSize int) time.Duration {
	ch := make(chan int, bufSize)
	input := bytes.NewBuffer(data)
	var wg sync.WaitGroup

	first := time.Now()

	wg.Add(1)
	go func() {
		recv(ch)
		wg.Done()
	}()

	send(input, ch)
	close(ch)
	wg.Wait()

	return time.Since(first)
}

// recv 接受channel
func recv(ch chan int) {
	var total int
	for n := range ch {
		total += n
	}
}

// send 发送channel
func send(r io.Reader, ch chan int) {
	buf := make([]byte, 1)

	for {
		n, err := r.Read(buf)
		if err != nil || n == 0 {
			break
		}
		ch <- int(buf[0])
	}
}

// makePlot 绘制阻塞图，对比无缓冲channels和有缓冲channel之间的差值。
func makePlot(xys plotter.XYs) error {
	p := plot.New()
	p.X.Label.Text = "Buffer Length"
	p.Y.Label.Text = "Latency"
	p.Title.Text = "Latencies(difference from the unbuffered channel)"

	if err := plotutil.AddLinePoints(p, "Latencies", xys); err != nil {
		return err
	}

	return p.Save(10*vg.Inch, 5*vg.Inch, "latencies.png")
}
