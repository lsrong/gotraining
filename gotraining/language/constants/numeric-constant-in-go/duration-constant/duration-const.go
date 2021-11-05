package main

import (
	"fmt"
	"time"
)

const fiveSecond = 5 * time.Second

func main() {
	now := time.Now()
	lessFiveNanoSeconds := now.Add(-5)
	lessFiveSeconds := now.Add(-fiveSecond)

	fmt.Printf("Now     : %v\n", now)
	fmt.Printf("Nano    : %v\n", lessFiveNanoSeconds)
	fmt.Printf("Seconds : %v\n", lessFiveSeconds)

	/**
	OUTPUT:
	Now     : 2021-10-28 10:34:51.210002556 +0800 CST m=+0.000037267
	Nano    : 2021-10-28 10:34:51.210002551 +0800 CST m=+0.000037262
	Seconds : 2021-10-28 10:34:46.210002556 +0800 CST m=-4.99996273
	*/
	//
	//var difference int = -5
	//var lessFiveNano = now.Add(difference)

}
