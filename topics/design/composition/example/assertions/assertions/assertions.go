package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Sample program demonstrating that type assertions are a runtime and
// not compile time construct.

// 类型断言是运行时而非编译时构造的示例程序。

type car struct{}

func (car) String() string {
	return "Vroom!"
}

type cloud struct{}

func (cloud) String() string {
	return "Big Data"
}

func main() {
	rand.Seed(time.Now().UnixNano())
	mvs := []fmt.Stringer{
		car{},
		cloud{},
	}
	for i := 0; i < 10; i++ {
		rn := rand.Intn(2)

		// 运行时执行
		// 执行类型断言，我们在随机选择的接口值中有一个具体的云类型。
		if v, is := mvs[rn].(cloud); is {
			fmt.Println("Get Lucky: ", v)
			continue
		}
		fmt.Println("Got Unlucky")
	}
}
