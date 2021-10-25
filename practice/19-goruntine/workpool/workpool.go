package workpool

import (
	"fmt"
	"github.com/learning_golang/logger"
	"math/rand"
)

// 计算数字
type Job struct {
	Id     int
	Number int
}

// 计算结果
type Result struct {
	Job *Job
	Sum int
}

// 实际业务处理
func Progress(job *Job, retChan chan *Result) {
	var (
		sum    int
		number int
	)
	number = job.Number
	for number != 0 {
		sum = sum + number%10
		number /= 10
	}
	result := &Result{
		Job: job,
		Sum: sum,
	}
	retChan <- result
}

// 处理函数
func Worker(jobChan chan *Job, retChan chan *Result) {
	for job := range jobChan {
		Progress(job, retChan)
	}
}

// 开启线程池
func Workpool(workNum int, jobChan chan *Job, retChan chan *Result) {
	for i := 0; i < workNum; i++ {
		go Worker(jobChan, retChan)
	}
}

// 打印结果
func PrintResult(retChan chan *Result) {
	logConf := map[string]string{
		"path":  "/Users/lsrong/Work/Project/Test",
		"level": "debug",
	}
	log, _ := logger.NewFileLogger(logConf)
	for ret := range retChan {
		job := ret.Job
		fmt.Printf("Job:id=%d,number=%d; result=%d\n", job.Id, job.Number, ret.Sum)
		log.Debug("Job:id=%d,number=%d; result=%d\n", job.Id, job.Number, ret.Sum)
	}
	log.Close()
}

func Start() {
	jobChan := make(chan *Job, 1000)
	retChan := make(chan *Result, 1000)
	workNum := 64
	Workpool(workNum, jobChan, retChan)
	go PrintResult(retChan)
	var id int
	for {
		id++
		job := &Job{
			Id:     id,
			Number: rand.Int(),
		}
		jobChan <- job
	}
}
