package goruntine

import (
	"fmt"
	"time"
)

func Hello() {
	fmt.Println("hello goroutine")
}

func TestGoruntine() {
	go Hello()
	fmt.Println("End exec go runtine")
	time.Sleep(time.Second)
}

func TestMultiGoruntine() {
	for i := 0; i < 10; i++ {
		go Hello()
	}
	fmt.Println("End exec go runtine")
	time.Sleep(time.Second)
}

func numbers() {
	for i := 0; i < 5; i++ {
		fmt.Printf("number %d\n", i)
		time.Sleep(time.Millisecond * 250)
	}
}

func alph() {
	for i := 'a'; i < 'f'; i++ {
		fmt.Printf("alph %d \n", i)
		time.Sleep(time.Millisecond * 450)
	}
}

func TestTimeGoruntine() {
	go numbers()
	go alph()
	time.Sleep(time.Millisecond * 3000)
	fmt.Println("terminated")
}
