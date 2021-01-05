package channel

import (
	"fmt"
	"sync"
	"time"
)

func Base(){
	var c chan int
	fmt.Printf("c = %v \n", c)
	c = make(chan int, 1)
	fmt.Printf("c = %v \n", c)

	// 入队
	c <- 100

	// 出队
	data := <-c
	fmt.Printf("out channel data=%d \n", data)
}
func BufChannel(){
	ch := make(chan string, 3)
	//s := <- ch
	ch <- "hello"
	ch <- "world"
	ch <- "!!"

	s1 := <- ch
	s2 := <- ch
	s3 := <- ch

	fmt.Println(s1,s2,s3)
}

func Producer(ch chan int){
	for i:=0; i < 10;i++{
		ch <- i
		fmt.Printf("Wite %d successful \n", i)
	}
	close(ch)
}


func RuntineChannel(){
	ch := make(chan int)
	go Producer(ch)
	time.Sleep(1 * time.Second)
	for v := range ch{
		fmt.Printf("Read value %d \n", v)
	}
}

func Hello(exit chan bool){
	fmt.Println("goruntine sync used bool channel")
	time.Sleep(5*time.Second)
	exit <- true
}
// 协程同步
func Sync(){
	exit := make(chan bool)
	go Hello(exit)
	<- exit
	fmt.Println("exit")
}

// 只进不出
func SendCh(ch chan <- int){
	ch <- 100
}
// 不进只出
func ReadCh(ch <- chan int){
	data :=<-ch
	fmt.Printf("data=%d \n",data)
}
// 单向队列
func SingleChan(){
	ch := make(chan int)
	go SendCh(ch)
	ReadCh(ch)
}

var ret int
func process(i int, wg *sync.WaitGroup){
	ret += i
	wg.Done()
}
// 同步操作
func WaitGroup(){
	var wg sync.WaitGroup
	wg.Wait()
	for i:=0;i<10;i++{
		wg.Add(1)
		go process(2, &wg)
	}
	wg.Wait()
	fmt.Printf("Ret=%d", ret)

}




