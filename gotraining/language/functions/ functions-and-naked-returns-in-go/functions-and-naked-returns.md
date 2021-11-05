# Functions and Naked Returns In Go -- William Kennedy

在 Go 中，从函数返回的值是按值传递的。当涉及到从函数返回值时，Go 为您提供了一些很好的灵活性。

这是从函数返回两个值的简单示例：

```go
package main

import (
	"fmt"
)

func main() {
	id, err := ReturnId()

	if err != nil {
		fmt.Printf("ERROR: %s", err)
		return
	}

	fmt.Printf("Id: %d\n", id)
}

func ReturnId() (int, error) {
	id := 10
	return id, nil
}
```

函数 ReturnId 返回一个整数类型和错误类型的值。这是在 Go 中非常常见的事情。错误处理是通过从您的函数和调用函数在继续之前评估该值返回一个错误类型的值来执行的。

如果您在函数调用返回后出于某种原因不关心错误，您可以执行以下操作：

```go
id, _ := ReturnId()
```

这次我用下划线表示第二个返回参数的返回值，也就是错误。这真的很好，因为我不需要声明一个变量来保存传入的值，我可以简单地忽略它。

您还可以选择命名返回参数：

```go
func ReturnId() (id int, err error) {
    id = 10
    return id, err
}
```

如果您命名返回参数，则就像使用函数参数一样创建局部变量。这次设置 id 变量时，我从短变量声明中删除了冒号 (:) 并将其转换为赋值操作。然后在返回中我正常指定返回变量。

命名返回参数是记录返回内容的好方法。您还可以使用命名参数做其他事情，或者不做：

```go
func ReturnId() (id int, err error) {
    id = 10
    return
}
```

这就是所谓的 naked return。我已经从 return 语句中删除了参数。Go 编译器会自动在返回参数局部变量中返回当前值。虽然这真的很酷，但您需要注意阴影：

```go
func ReturnId() (id int, err error) {
   id = 10

   if id == 10 {
   	  // 变量重命名冲突 
      err := fmt.Errorf("Invalid Id\n")
      return
   }

   return
}
```

如果您尝试编译它，您将收到以下编译器错误：

```
err is shadowed during return
```

要了解为什么会出现此错误，您需要了解大括号在函数内部的作用。每组大括号定义了一个新的范围级别。以这段代码为例：

```go
func main() {
   id := 10
   id := 20

   fmt.Printf("Id: %d\n", id)
}
```

如果您尝试编译此代码，则会出现以下错误：

```
no new variables on left side of :=
```

这是有道理的，因为您试图两次声明相同的变量名。如果我们将代码更改为如下所示，错误就会消失：

```go
func main() {
   id := 10

   {
       id := 20
       fmt.Printf("Id: %d\n", id)
   }

   fmt.Printf("Id: %d\n", id)
}
```

大括号定义了一个新的堆栈框架，因此定义了一个新的范围。变量名称可以在新的一组大括号内重复使用。当代码到达结束大括号时，堆栈的一小部分被弹出。

再次查看导致阴影错误的代码：

```go
func ReturnId() (id int, err error) {
   id = 10

   if id == 10 {
   	  // 没使用返回参数列表定义的err 变量
      err := fmt.Errorf("Invalid Id\n")
      return
   }

   return
}
```

在 if 语句中，我们创建了一个名为 err 的新变量。我们没有使用声明为函数返回参数的 err 变量。编译器识别出这一点并产生错误。如果编译器没有报告此错误，您将永远不会看到 if 语句中发生的错误。return err 变量是默认传递的

变量 在使用 defer 语句时，命名返回参数会非常方便：

```go
func ReturnId() (id int, err error) {
   defer func() {
      if id == 10 {
         err = fmt.Errorf("Invalid Id\n")
      }
   }()

   id = 10

   return
}
```

由于返回参数已命名，因此您可以在 defer 函数中引用它们。您甚至可以在 defer 调用中更改返回参数的值，调用函数将看到新值。此版本将显示错误消息。

您需要注意 defer 语句是与其余代码内联计算的：
```go
func ReturnId() (id int, err error) {
   defer func(id int) {
      if id == 10 {
         err = fmt.Errorf("Invalid Id\n")
      }
   }(id)

   id = 10

   return
}
```

此版本不显示错误消息。id 的值直到 defer 语句被评估后才为 10。

有时使用命名返回参数是有意义的，例如在函数顶部使用 defer 语句时。如果您从函数中传递原始值，那么这样的事情就没有意义：
```go

package main

import (
   "fmt"
)

func main() {
   ans := AddNumbers(10, 12)
   fmt.Printf("Answer: %d\n", ans)
}

func AddNumbers(a int, b int) (result int) {
   return a + b
}
```

返回参数是为 AddNumbers 函数命名的，但从未使用过。相反，我们直接从 return 中返回操作的答案。这向您展示了即使您命名了返回参数，您仍然可以返回任何您想要的值。
