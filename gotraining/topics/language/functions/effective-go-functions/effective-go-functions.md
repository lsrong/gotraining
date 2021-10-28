# Functions 函数
## Multiple return values 多返回值
Go 不寻常的特性之一是函数和方法可以返回多个值。这种形式可用于改进 C 程序中的几个笨拙的习惯用法：带内错误返回，例如-1 for EOF 和修改按地址传递的参数。

在 C 中，写入错误由负计数表示，错误代码隐藏在易失性位置。在 Go 中，Write 可以返回一个计数和一个错误：“是的，你写了一些字节，但不是全部，因为你填满了设备”。`os`包中写入文件的方法 `Wirite` 的定义：
```go
func (file *File) Write(b []byte) (n int, err error)
```
正如文档所说，当 n != len(b)时候,它返回写入的字节数和non-nil error。这是一种常见的风格；有关更多示例，请参阅错误处理部分。

类似的方法不需要传递指向返回值的指针来模拟引用参数。这是一个简单的函数，从字节切片中的某个位置获取一个数字，返回该数字和下一个位置。
```go
func nextInt(b []byte, i int)(int, int){
	for ;i< len(b) && isDigit(b[i]); i++{
    }
    x := 0
    for ;i<len(b)&&isDigit(b[i]); i++{
        x = x*10 + int(b[i]) - '0'	
    }
    return x, i
}
```

## Named result parameters 命名结果参数
Go 函数的返回或结果“参数”可以命名并用作常规变量，就像传入参数一样。当命名时，它们在函数开始时被初始化为它们的类型的零值；如果函数执行return不带参数的语句，则结果参数的当前值将用作返回值。

名称不是强制性的，但它们可以使代码更短、更清晰：它们是文档。如果我们命名nextInt它的结果就很明显返回的int 是哪个
```go
func nextInt(b []byte, pos int)(value, nextPost int)
```
因为命名结果被初始化并绑定到一个简单的返回，它们可以简化和澄清。这是一个io.ReadFull很好地使用它们的版本：
```go
func ReadFull(r Reader, buf []byte) (n int, err error) {
    for len(buf) > 0 && err == nil {
        var nr int
        nr, err = r.Read(buf)
        n += nr
        buf = buf[nr:]
    }
    return
}
```

## Defer 
Go 的defer语句安排一个函数调用（ 延迟函数）在函数执行defer返回之前立即运行。这是一种不寻常但有效的方法来处理诸如必须释放资源的情况，而不管函数采用哪条路径返回。规范示例是解锁互斥锁或关闭文件。

```go
// Contents returns the file's contents as a string.
func Contents(filename string) (string, error) {
    f, err := os.Open(filename)
    if err != nil {
        return "", err
    }
    defer f.Close()  // f.Close will run when we're finished.

    var result []byte
    buf := make([]byte, 100)
    for {
        n, err := f.Read(buf[0:])
        result = append(result, buf[0:n]...) // append is discussed later.
        if err != nil {
            if err == io.EOF {
                break
            }
            return "", err  // f will be closed if we return here.
        }
    }
    return string(result), nil // f will be closed if we return here.
}
```

推迟对诸如此类的函数的调用Close有两个优点。首先，它保证您永远不会忘记关闭文件，如果您稍后编辑该函数以添加新的返回路径，则很容易犯这个错误。其次，这意味着关闭在于开启资源的附近，这比将其放在函数的末尾要清晰得多。

延迟函数的参数（如果函数是方法，则包括接收者）在延迟 执行时计算，而不是在调用执行时计算。除了避免担心在函数执行时变量会改变值，这意味着单个延迟调用站点可以延迟多个函数执行。这是一个愚蠢的例子。
```go
// defer要避免计算业务逻辑
for i := 0; i < 5; i++ {
    defer fmt.Printf("%d ", i)
}
```
延迟函数以 LIFO (后进先出 last in first out)顺序执行，因此该代码将导致 4 3 2 1 0在函数返回时打印。一个更合理的例子是通过程序跟踪函数执行的简单方法。我们可以编写几个简单的跟踪例程，如下所示：

```go
package main

import "fmt"

func trace(s string) string {
	fmt.Println("entering:",s)
	return s
}

func un(s string){
	fmt.Println("leaving:", s)
}

func a(){
	defer un(trace("a"))
	fmt.Println("in a")
}

func b(){
	defer un(trace("b"))
	fmt.Println("in b")
	a()
}

func main(){
	b()
}
```

输出:
```
entering: b
in b
entering: a
in a
leaving: a
leaving: b
```

对于习惯于其他语言的块级资源管理的程序员来说，这defer可能看起来很奇怪，但它最有趣和最强大的应用程序恰恰来自于它不是基于块而是基于函数的事实。在 `panic` 和 `recover` 的部分中,我们将看到其可能性的另一个示例。


