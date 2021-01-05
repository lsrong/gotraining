package main

import "github.com/learning_golang/19-goruntine/workpool"

func main() {
	// Goruntine
	//goruntine.TestGoruntine()
	//
	//goruntine.TestMultiGoruntine()
	//
	//goruntine.TestTimeGoruntine()

	// Channel
	//channel.Base()
	//
	//channel.BufChannel()
	//
	//channel.RuntineChannel()

	//channel.Sync()

	//channel.SingleChan()

	//channel.WaitGroup()

	// Workpool
	//job := &workpool.Job{
	//	Id: 1,
	//	Number: 224,
	//}
	//ret := make(chan *workpool.Result, 1)
	//workpool.Progress(job, ret)
	//data :=<- ret
	//fmt.Println(data)
	workpool.Start()
}
