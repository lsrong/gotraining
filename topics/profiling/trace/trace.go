package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
	"runtime/trace"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

// 执行一系列 IO 相关任务以更好地理解 Go 中的跟踪的示例程序。

type (
	item struct {
		XMLName     xml.Name `xml:"item"`
		Title       string   `xml:"title"`
		Description string   `xml:"description"`
	}
	channel struct {
		XMLName xml.Name `xml:"channel"`
		Items   []item   `xml:"item"`
	}
	document struct {
		XMLName xml.Name `xml:"rss"`
		Channel channel  `xml:"channel"`
	}
)

var (
	cpuProfile   string
	traceProfile string
	since        time.Time
)

func init() {
	flag.StringVar(&cpuProfile, "cpuprofile", "", "cpu profile out file")
	flag.StringVar(&traceProfile, "traceprofile", "", "trace out file")
	flag.Parse()

	since = time.Now()
}

func main() {
	if cpuProfile != "" {
		// 开启cpu 性能取样
		f, err := os.OpenFile(cpuProfile, os.O_CREATE|os.O_WRONLY, 0744)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatalf("StartCPUProfie failed: %v", err)
		}
		defer pprof.StopCPUProfile()
	} else if traceProfile != "" {
		// 开启cpu 性能取样
		f, err := os.OpenFile(traceProfile, os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		// 开启追踪
		if err := trace.Start(f); err != nil {
			log.Fatalf("Trace start failed: %v", err)
		}
		defer trace.Stop()
	}

	// 示例的io文件
	docs := make([]string, 4000)
	for i := range docs {
		docs[i] = fmt.Sprintf("newsfeed-%.4d.xml", i)
	}

	// 从文件中检索某个内容并统计出现的次数.
	search := "president"
	//n := freq(search, docs)	// 同步扫描
	//n := freqConcurrent(search, docs)	// 并发扫描
	//n := freqConcurrentSem(search, docs)	// 频率并发扫描
	//n := freqProcessors(search, docs) // 协程池扫描
	n := freqActor(search, docs) // 分布式扫描

	fmt.Printf("Searching %d files, found %s %d times, Using %.2fs \n", len(docs), search, n, time.Since(since).Seconds())

}

// freq 同步扫描
func freq(search string, docs []string) int {
	var found int
	for _, doc := range docs {
		file, err := os.Open(fmt.Sprintf("%s.xml", doc[:8]))
		if err != nil {
			log.Printf("Opening [%s], Error: %v \n", doc, err)
			return 0
		}
		data, err := io.ReadAll(file)
		file.Close()
		if err != nil {
			log.Printf("Reading [%s], Error: %v ", doc, err)
			return 0
		}
		var d document
		if err = xml.Unmarshal(data, &d); err != nil {
			log.Printf("Unmarshal Error: %s", err)
			return 0
		}
		for _, channelItem := range d.Channel.Items {
			if strings.Contains(channelItem.Title, search) {
				found++
			}

			if strings.Contains(channelItem.Description, search) {
				found++
			}
		}
	}
	return found
}

// freqConcurrent 并发扫描
func freqConcurrent(search string, docs []string) int {
	var found int64
	var wg sync.WaitGroup
	// 使用goroutine加速
	wg.Add(len(docs))
	for _, doc := range docs {
		go func(file string) {
			var lFound int64
			defer func() {
				atomic.AddInt64(&found, lFound)
				wg.Done()
			}()

			f, err := os.Open(fmt.Sprintf("%s.xml", file[:8]))
			if err != nil {
				log.Printf("Opening [%s], Error: %v \n", file, err)
				return
			}
			data, err := io.ReadAll(f)
			f.Close()
			if err != nil {
				log.Printf("Reading [%s], Error: %v ", file, err)
				return
			}
			var d document
			if err = xml.Unmarshal(data, &d); err != nil {
				log.Printf("Unmarshal Error: %s", err)
				return
			}

			for _, channelItem := range d.Channel.Items {
				if strings.Contains(channelItem.Title, search) {
					lFound++
				}

				if strings.Contains(channelItem.Description, search) {
					lFound++
				}
			}
		}(doc)
	}

	wg.Wait()

	return int(found)
}

// freqConcurrentSem 控制一定并发频率次数扫描
func freqConcurrentSem(search string, docs []string) int {
	var found int64
	var wg sync.WaitGroup
	wg.Add(len(docs))
	ch := make(chan bool, runtime.GOMAXPROCS(0))
	for _, doc := range docs {
		go func(doc string) {
			// 阻塞等待，实现goroutine并发频率控制
			ch <- true
			{
				var lFound int64
				defer func() {
					atomic.AddInt64(&found, lFound)
					wg.Done()
				}()
				f, err := os.Open(fmt.Sprintf("%s.xml", doc[:8]))
				if err != nil {
					fmt.Printf("Opening %s Error: %v", doc, err)
					return
				}
				data, err := io.ReadAll(f)
				f.Close()
				if err != nil {
					fmt.Printf("Reading %s Error: %v", doc, err)
					return
				}
				var d document
				if err := xml.Unmarshal(data, &d); err != nil {
					fmt.Printf("Unable to xml.Unmarshal: %v", err)
					return
				}
				for _, i := range d.Channel.Items {
					if strings.Contains(i.Title, search) {
						lFound++
					}
					if strings.Contains(i.Description, search) {
						lFound++
					}
				}
			}
			<-ch
		}(doc)
	}

	wg.Wait()

	return int(found)
}

// freqProcessors 使用协程池方式处理任务
func freqProcessors(search string, docs []string) int {
	var wg sync.WaitGroup
	var found int64
	grs := runtime.GOMAXPROCS(0)
	docCh := make(chan string, grs)
	wg.Add(grs)
	for i := 0; i < grs; i++ {
		go func() {
			var lFound int64
			defer func() {
				atomic.AddInt64(&found, lFound)
				wg.Done()
			}()
			for doc := range docCh {
				f, err := os.Open(fmt.Sprintf("%s.xml", doc[:8]))
				if err != nil {
					fmt.Printf("Opening %s Error: %v", doc, err)
					return
				}
				data, err := io.ReadAll(f)
				f.Close()
				if err != nil {
					fmt.Printf("Reading %s Error: %v", doc, err)
					return
				}
				var d document
				if err := xml.Unmarshal(data, &d); err != nil {
					fmt.Printf("Unable to xml.Unmarshal: %v", err)
					return
				}
				for _, i := range d.Channel.Items {
					if strings.Contains(i.Title, search) {
						lFound++
					}
					if strings.Contains(i.Description, search) {
						lFound++
					}
				}
			}
		}()
	}

	for _, doc := range docs {
		docCh <- doc
	}
	close(docCh)
	wg.Wait()
	return int(found)
}

// freqActor 分部操作，每个任务一个goroutine
func freqActor(search string, docs []string) int {
	// 文件操作
	files := make(chan *os.File, 100)
	go func() {
		for _, doc := range docs {
			f, err := os.Open(fmt.Sprintf("%s.xml", doc[:8]))
			if err != nil {
				fmt.Printf("Opening [%s] Error: %v", doc, err)
				return
			}
			files <- f
		}

		// 写入操作结束关闭通道
		close(files)
	}()

	// 数据读取
	data := make(chan []byte, 100)
	go func() {
		for f := range files {
			d, err := io.ReadAll(f)
			if err != nil {
				fmt.Printf("Reading Error: %v", err)
				return
			}
			data <- d
		}
		close(data)
	}()

	// 序列化
	rss := make(chan *document, 100)
	go func() {
		for dd := range data {
			var d document
			if err := xml.Unmarshal(dd, &d); err != nil {
				fmt.Printf("xml.Unmarshal Error: %v", err)
				return
			}

			rss <- &d
		}
		close(rss)
	}()

	// 处理业务统计
	var found int
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for d := range rss {
			for _, i := range d.Channel.Items {
				if strings.Contains(i.Title, search) {
					found++
				}

				if strings.Contains(i.Description, search) {
					found++
				}
			}
		}
		wg.Done()
	}()
	wg.Wait()
	return found
}
