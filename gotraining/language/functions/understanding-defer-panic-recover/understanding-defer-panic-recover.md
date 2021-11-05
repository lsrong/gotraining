# Understanding Defer, Panic and Recover -- William Kennedy

一开始了解如何在应用程序中使用 Defer 和 Recover 可能有点棘手，尤其是在您习惯使用 try/catch 块的情况下。您可以实现一种模式来在 Go 中提供相同类型的 try/catch 保护。在我向您展示这个之前，您需要了解 defer、panic 和 recover 的工作原理。

首先，您需要了解关键字 defer 的复杂性。从这段代码开始：
```go
package main

import (
  "fmt"
)

func main() {
  test()
}

func mimicError(key string) error {
  return fmt.Errorf("Mimic Error : %s", key)
}

func test() {
  fmt.Println( "开始测试")

  err := mimicError("1")

  defer func() {
    fmt.Println("Start Defer")
 
    if err != nil {
      fmt.Println("Defer Error:", err )
    }
  }()

  fmt.Println("结束测试")
}
```

MimicError 函数是一个测试函数，用于模拟错误。它遵循 Go 约定，即使用错误类型在出现问题时返回指示。

在 Go 中，错误类型被定义为一个接口：

```go
type Error interface {
	Error() string
}
```

如果您目前不了解 Go 接口是什么，那么现在这可能会有所帮助。任何实现了Error()函数的类型都实现了这个接口，并且可以用作这个类型的变量。MimicError 函数使用errors.New(string) 创建错误类型变量。错误类型可以在错误包中找到。

测试函数产生以下输出：
```
Start Test
End Test
Start Defer
Defer Error : Mimic Error : 1
```

当您研究输出时，您会看到 Test 函数开始和结束。然后就在 Test 函数永久终止之前，调用了内联 defer 函数。这里正在发生两件有趣的事情。首先， defer 关键字将内联函数的执行推迟到 Test 函数结束。其次，由于 Go 支持闭包，内联函数可以访问 err 变量，并将其消息“Mimic Error : 1”写入标准输出。

您可以随时在函数内部定义延迟函数。如果 defer 函数需要状态，就像在这种情况下使用 err 变量一样，那么它必须在 defer 函数被定义之前存在。

现在稍微改变一下测试功能：

```go
func test() {
  fmt.Println("Start Test")

  err := mimicError("1")

  defer func() {
    fmt.Println("Start Defer")

    if err != nil {
      fmt.Println("Defer Error:", err)
    }
  }()

  err = mimicError("2")

  fmt.Println("End Test")
}
```

这次代码在创建内联 defer 函数后第二次调用 MimicError 函数。这是输出：

```
Start Test
End Test
Start Defer
Defer Error : Mimic Error : 2
```

除了一个变化外，输出与第一个测试相同。这次内联 defer 函数写了“Mimic Error：2”。内联 defer 函数似乎引用了 err 变量。因此，如果 err 变量的状态在调用内联 defer 函数之前的任何时间发生变化，您将看到该值。要验证内联 defer 函数是否正在获取对 err 变量的引用，请更改代码以在 Test 函数和内联 defer 函数中写入 err 变量的地址。

```go
func test() {
  fmt.Println("Start Test")

  err := mimicError("1")

  fmt. Println("Err Addr:", &err)

  defer func() {
    fmt.Println("Start Defer")

    if err != nil {
      fmt.Println("Err Addr Defer:", &err)
      fmt.Println("Defer Error:", err)
    }
  }()

  err = mimicError("2")

  fmt.Println("End Test")
}
```

从下面的输出中可以看出，内联 defer 函数对 err 变量具有相同的引用。该地址在 Test 函数内部和 inline defer 函数内部是相同的。

```
Start Test
Err Addr: 0x2101b3200
End Test
Start Defer
Err Addr Defer: 0x2101b3200
Defer Error : Mimic Error : 2
```

只要在 Test 函数终止之前声明了 defer 函数，就会执行 defer 函数。这很棒，但我想要的是始终将 defer 语句放在任何函数的开头的能力。这样就保证每次执行函数时都会调用 defer 函数，我不必过度考虑将 defer 语句放在哪里。在这里适用：  “当你有两种相互竞争的理论做出完全相同的预测时，越简单越好”。我想要的是一个简单的模式，不需要任何思考就可以复制。

唯一的问题是需要先定义 err 变量，然后才能实现 defer 语句。幸运的是，Go 允许返回变量具有名称。

现在将整个程序更改如下：
```go
package main

import (
  "fmt"
)

func main() {
  if err := test(); err != nil {
    fmt.Printf("Main Error: %v\n", err)
  }
}

func mimicError(key string) error {
  return fmt.Errorf("Mimic Error : %s", key)
}

func test() (err error) {
  defer func() {
    fmt.Println("Start Defer")

    if err != nil {
      fmt.Println("Defer Error:", err)
    }
  }()

  fmt.Println("Start Test")

  err = mimicError("1")

  fmt.Println("End Test")
  return err
}
```

Test 函数现在定义了一个名为 err 的返回变量，类型为 error。这很好，因为 err 变量立即存在，您可以将 defer 语句放在函数的最开头。此外，Test 函数现在遵循 Go 约定并将错误类型返回给调用例程。

当你运行程序时，你会得到以下输出：

```
Start Test
End Test
Start Defer
Defer Error : Mimic Error : 1
Main Error: Mimic Error : 1
```

现在是时候谈谈内置函数 panic 了。当任何 Go 函数调用 panic 时，应用程序的正常流程就会停止。调用 panic 的函数会立即结束，并导致调用堆栈上出现恐慌的连锁反应。同一个调用栈中的所有函数都会一个接一个地结束，就像多米诺骨牌倒下一样。最终，恐慌到达调用堆栈的顶部，应用程序崩溃。一件好事是，任何现有的延迟函数都将在此恐慌序列期间执行，并且它们有能力阻止崩溃。

看看这个新的 Test 函数，它调用内置的恐慌函数并从调用中恢复：

```go
func test() error {
  defer func() {
    fmt.Println("Start Panic Defer")

    if r := recover(); r != nil {
      fmt.Println("Defer Panic:", r)
    }
  }()

  fmt.Println("Start Test")
  panic("Mimic Panic")
}
```

仔细看看新的内联 defer 函数：

```go
  defer func() {
    fmt.Println("Start Panic Defer")

    if r := recover(); r != nil {
      fmt.Println("Defer Panic:", r)
    }
  }()
```

内联 defer 函数现在正在调用另一个内置函数recover。恢复函数阻止链式反应在调用堆栈上进一步向上移动。这就像滑动多米诺骨牌，这样就不会再倒下。**`recover`函数只能在`defer`函数内部使用。这是因为在`panic`连锁反应期间，只会执行`defer`函数。**

如果调用了recover 函数并且没有发生panic，则recover 函数将返回nil。如果发生恐慌，那么恐慌就会停止，并返回给恐慌调用的值。

这次代码不是调用 MimicError 函数，而是调用内置的恐慌函数来模拟恐慌。查看运行代码的输出：
```
Start Test
Start Panic Defer
Defer Panic : Mimic Panic
```

内联 defer 函数捕获 panic，将其打印到屏幕上并使其停止运行。另请注意，永远不会显示“结束测试”。该函数在调用 panic 后立即终止。

这很好，但如果有错误，我仍然想显示它。Go 和 defer 关键字的一个很酷的地方是你可以一次声明多个 defer 函数。

更改测试功能如下：
```go
func test() (err error) {
  defer func() {
    fmt.Println("Start Panic Defer")

    if r := recover(); r != nil {
      fmt.Println("Defer Panic :", r)
    }
  }()

  defer func() {
    fmt.Println("Start Defer")

    if err != nil {
      fmt.Println("Defer Error:", err)
    }
  }()

  fmt.Println("Start Test")

  err = mimicError("1")

  panic("Mimic Panic")
}
```


现在两个内联 defer 函数都已合并到 Test 函数的开头。首先是从恐慌中恢复的内联 defer 函数，然后是打印错误的内联 defer 函数。需要注意的一点是，Go 将按照定义的相反顺序（先进后出）执行这些内联延迟函数。

运行代码并查看输出：

```
Start Test
Start Error Defer
Defer Error : Mimic Error : 1
Start Panic Defer
Defer Panic : Mimic Panic
Main Error: Mimic Error : 1
```

Test 函数按预期启动，对 panic 的调用停止了 Test 函数的执行。这会导致首先调用打印错误的内联 defer 函数。由于 Test 函数在恐慌之前调用了 MimicError 函数，因此会打印错误。然后调用从 panic 中恢复的内联 defer 函数并恢复panic。

这段代码有一个问题。主函数不知道避免了恐慌。主函数只知道由于 MimicError 函数调用发生了错误。这是不好的。我希望主函数知道导致panic的错误。这确实是必须报告的错误。

在处理恐慌的内联 defer 函数中，我们需要将导致 panic 的错误分配给 err 变量。

```go
func test() (err error) {
  defer func() {
    fmt.Println("Start Panic Defer")

    if r := recover(); r != nil {
      fmt.Printf("Defer Panic : %v\n", r)

      err = fmt.Errorf("%v", r)
    }
  }()

  defer func() {
    fmt.Println("Start Defer")

    if err != nil {
      fmt.Println("Defer Error:", err)
    }
  }()

  fmt.Println("Start Test")

  err = mimicError("1")

  panic("Mimic Panic")
}
```

现在，当您运行代码时，您会得到以下输出：

```
Start Test
Start Error Defer
Defer Error : Mimic Error : 1
Start Panic Defer
Defer Panic : Mimic Panic
Main Error: Mimic Panic
```



这次 main 函数报告了导致 panic 的错误。

一切看起来都不错，但这段代码并不是真正可扩展的。有两个内联 defer 函数很酷但不实用。我需要的是一个可以处理错误和恐慌的单一函数。

这是完整程序的修订版，带有一个名为 catchPanic 的新函数：
```go
package main

import (
  "fmt"
)

func main() {
  if err := test(); err != nil {
    fmt.Println("Main Error:", err)
  }
}

func catchPanic(err error, functionName string) {
  if r := recover(); r != nil {
    fmt.Printf("%s : PANIC Defered : %v\n", functionName, r)

    if err != nil {
      err = fmt.Errorf("%v", r)
    }
  } else if err != nil {
    fmt.Printf("%s : ERROR : %v\n", functionName, err)
  }
}

func mimicError(key string) error {
  return fmt.Errorf("Mimic Error : %s", key)
}

func test() (err error) {
  defer catchPanic(err, "Test")
  fmt.Println("Start Test")

  err = mimicError("1")

  fmt.Println("End Test")
  return err
}
```

新函数 catchPanic 结合了紧急恢复和错误处理。这次不是定义内联 defer 函数，而是代码为 defer 语句使用了一个外部函数。

在使用新的 catchPanic defer 函数的第一次测试中，我们需要确保我们没有破坏我们的错误处理。

运行代码并查看输出：

```go
Start Test
End Test
Main Error: Mimic Error : 1
```

一切看起来都不错。现在我们需要测试一个 panic 。

```go
func test() (err error) {
  defer catchPanic(err, "Test")
  fmt.Println("Start Test")

  err = mimicError("1")
  panic("Mimic Panic")
}
```

运行代码并查看输出：

```
Start Test
Test5 : PANIC Defered : Mimic Panic
Main Error: Mimic Error : 1
```

这里有个问题。Main 提供了来自 MimicError 函数调用的错误，而不是来 panic 。什么地方出了错？

因为 defer 现在正在调用一个外部函数，所以代码失去了内联函数和闭包带来的所有优点。

更改代码以从 Test 函数和 CatchPanic defer 函数内部打印 err 变量的地址。

```go
func _CatchPanic(err error, functionName string) {
  if r := recover(); r != nil {
    fmt.Printf("%s : PANIC Defered : %v\n", functionName, r)

    fmt.Println("Err Addr Defer:", &err)
 
    if err != nil {
      err = fmt.Errorf("%v", r)
    }
  } else if err != nil {
    fmt.Printf("%s : ERROR : %v\n", functionName, err)
  }
}

func test() (err error) {
  fmt.Println("Err Addr:", &err)
  defer _CatchPanic(err, "Test7")
  fmt.Printf("Start Test\n")

  err = mimicError("1")

  panic("Mimic Panic")
}
```

当您运行代码时，您可以看到为什么 main 没有从 panic 中得到错误。

```
Err Addr: 0x2101b31f0
Start Test
Test5 : PANIC Defered : Mimic Panic
Err Addr Defer: 0x2101b3270
Main Error: Mimic Error : 1
```

当 Test 函数将 err 变量传递给 catchPanic defer 函数时，它是按值传递变量。在 Go 中，所有参数都是按值传递的。所以 catchPanic defer 函数有它自己的 err 变量副本。对 catchPanic 副本的任何更改都保留在 catchPanic 中。

为了解决传值问题，代码需要通过引用传递 err 变量。

```go
package main

import (
  "fmt"
)

func main() {
  if err := testFinal(); err != nil {
    fmt.Println("Main Error:", err)
  }
}

func catchPanic(err *error, functionName string) {
  if r := recover(); r != nil {
    fmt.Printf("%s : PANIC Defered : %v\n", functionName, r)

    if err != nil {
      *err = fmt.Errorf("%v", r)
    }
  } else if err != nil && *err != nil {
    fmt.Printf("%s : ERROR : %v\n", functionName, *err)
  }
}

func mimicError(key string) error {
  return fmt.Errorf("Mimic Error : %s", key)
}

func testFinal() (err error) {
  defer catchPanic(&err, "TestFinal")
  fmt.Println("Start Test")

  err = mimicError("1")

  panic("Mimic Panic")
}
```

现在运行代码并查看输出：

```
Start Test
Test6 : PANIC Defered : Mimic Panic
Main Error: Mimic Panic
```

主函数现在报告由于恐慌而发生的错误。

如果您还想捕获堆栈跟踪，只需将此更改为 catchPanic。请记住导入“运行时”。

```go
func catchPanic(err *error, functionName string) {
  if r := recover(); r != nil {
    fmt.Printf("%s : PANIC Defered : %v\n", functionName, r)

    // Capture the stack trace
    buf := make([]byte, 10000)
    runtime.Stack(buf, false)

    fmt.Printf("%s : Stack Trace : %s", functionName, string(buf))

    if err != nil {
      *err = fmt.Errorf("%v", r)
    }
  } else if err != nil && *err != nil {
    fmt.Printf("%s : ERROR : %v\n", functionName, *err)
  }
}
```

使用此模式，您可以实现可以处理错误和捕获恐慌情况的 Go 例程。在许多情况下，这些情况只需要记录或报告到调用堆栈即可正常处理。有一个地方来实现这种类型的代码，并用一种简单的方法将它集成到每个函数中，将减少错误并保持代码干净。

但是我了解到最好只使用这种模式来捕捉 panic 。将错误记录留给应用程序逻辑。如果没有，那么您可能会记录错误两次。

```go
func catchPanic(err *error, functionName string) {
  if r := recover(); r != nil {
    fmt.Printf("%s : PANIC Defered : %v\n", functionName, r)

    // Capture the stack trace
    buf := make([]byte, 10000)
    runtime.Stack(buf, false)

    fmt.Printf("%s : Stack Trace : %s", functionName, string(buf))

    if err != nil {
      *err = fmt.Errorf("%v", r)
    }
  }
}
```

一如既往，我希望这可以帮助您进行 Go 编程。