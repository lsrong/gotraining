# 理解指针和内存分配 - William Kennedy

在 Go 语言团队提供的文档中，您会找到关于指针和内存分配的重要信息。这是该文档的链接：

[http://golang.org/doc/faq#Pointers](http://golang.org/doc/faq#Pointers) 

我们需要从了解所有变量都包含一个值开始。基于变量所代表的类型将决定我们如何使用它来操作它所包含的内存。阅读这篇文章以了解更多信息：[了解 Go 中的类型](https://www.ardanlabs.com/blog/2013/07/understanding-type-in-go.html)

在 Go 中，我们可以创建包含值本身的“值”或值的地址的变量。当变量的“值”是地址时，变量被视为指针。

在下图中，我们有一个名为 myVariable 的变量。myVariable 的“值”是分配给相同类型的值的地址。myVariable 被认为是一个指针变量。

![截屏](https://www.ardanlabs.com/images/goinggo/Screen+Shot+2013-07-27+at+2.57.16+PM.png)

![截屏](https://www.ardanlabs.com/images/goinggo/Screen+Shot+2013-07-27+at+3.01.52+PM.png)

在下图中，myVariable 的“值”是值本身，而不是对值的引用。

要访问值的属性，我们使用选择器运算符。选择器运算符允许我们访问值中的特定字段。语法始终为 Value.FieldName，其中句点 (.) 是选择器运算符。

在 C 编程语言中，我们需要根据使用的变量类型使用不同的选择器运算符。如果变量的“值”是值，我们使用句点 (.)。如果变量的“值”是地址，我们使用箭头 (->)。

Go 的一个真正好处是您无需担心要使用哪种类型的选择器运算符。在 Go 中，无论变量是值还是指针，我们只使用句点 (.)。编译器负责访问值的底层细节。

那么为什么这一切都很重要？当我们开始使用函数来抽象和分解逻辑时，这变得很重要。最终您需要将变量传递给这些函数，并且您需要知道您传递的是什么。

在 Go 中，变量按值传递给函数。这意味着指定的每个变量的“值”都被复制到堆栈中以供该函数访问。在这个例子中，我们调用了一个函数，该函数应该改变在 main 中分配的值的状态。

```go
package main

import (
	"fmt"
	"unsafe"
)

type MyType struct {
	Value1 int
	Value2 string
}

func main() {
	// 初始化'值'类型
	myValue := MyType{10, "Bill"}
	pointer := unsafe.Pointer(&myValue)
	fmt.Printf("1. Addr: %v Value1 : %d Value2: %s\n",
		pointer,
		myValue.Value1,
		myValue.Value2)

	// 传递值副本给函数
	changeMyValue(myValue)

	// 调用完毕堆栈会弹出myValue副本.
	fmt.Printf("3. Addr: %v Value1 : %d Value2: %s\n",
		pointer,
		myValue.Value1,
		myValue.Value2)
}

// 	myValue实际为参数副本, 堆栈中会开辟新内存空间拷贝myValue,作为参数传递给函数调用
func changeMyValue(myValue MyType) {
	// 改变myValue结构体值,
	myValue.Value1 = 20
	myValue.Value2 = "Jill"

	// Create a pointer to the memory for myValue
	pointer := unsafe.Pointer(&myValue)

	// Display the address and values
	fmt.Printf("2. Addr: %v Value1 : %d Value2: %s\n",
		pointer,
		myValue.Value1,
		myValue.Value2)
}
```

这是程序的输出： 
1. Addr: 0xc00000c030 Value1 : 10 Value2: Bill
2. **Addr: 0xc00000c048 Value1 : 20 Value2: Jill**
3. Addr: 0xc00000c030 Value1 : 10 Value2: Bill

那么出了什么问题呢？函数对 main 的 myValue 所做的更改在函数调用后并未保持更改。main 中 myValue 变量的“值”不包含对该值的引用，它不是一个指针。main 中 myValue 变量的“值”就是值。**当我们将 main 中 myValue 变量的“值”传递给函数时，该值的副本被放置在堆栈中。该函数正在改变它自己的值版本。一旦函数终止，堆栈被弹出，并且副本在技术上消失**了。从未触及 main 中 myValue 变量的“值”。

为了解决这个问题，我们可以以某种方式分配内存以获取引用。然后 main 中 myValue 变量的“值”将是新值的地址，一个指针变量。然后我们可以更改函数以接受地址的“值”到值。

```go
package main

import (
	"fmt"
	"unsafe"
)

type MyType struct {
	Value1 int
	Value2 string
}

func main() {
	// 初始化引用类型
	myValue := &MyType{10, "Bill"}
	pointer := unsafe.Pointer(myValue)
	fmt.Printf("1. Addr: %v Value1 : %d Value2: %s\n",
		pointer,
		myValue.Value1,
		myValue.Value2)

	// 传递指针给函数
	changeMyValue(myValue)

	// 函数changMyValue修改的数据都是同一个数据块
	fmt.Printf("3. Addr: %v Value1 : %d Value2: %s\n",
		pointer,
		myValue.Value1,
		myValue.Value2)
}

// 	myValue参数的'值'为MyType指针,堆栈中开辟保存的为指针,此时myValue可以表示为指针变量,`值为指针的变量`
func changeMyValue(myValue *MyType) {
	// 改变myValue结构体值,
	myValue.Value1 = 20
	myValue.Value2 = "Jill"

	// Create a pointer to the memory for myValue
	pointer := unsafe.Pointer(myValue)

	// Display the address and values
	fmt.Printf("2. Addr: %v Value1 : %d Value2: %s\n",
		pointer,
		myValue.Value1,
		myValue.Value2)
}

```

当我们使用与号 (&) 运算符来分配值时，会返回一个引用。这意味着 main 中 myValue 变量的“值”现在是一个指针变量，其值是新分配值的地址。当我们将 main 中 myValue 变量的“值”传递给函数时，函数 myValue 变量现在包含值的地址，而不是副本。我们现在有两个指向同一个值的指针。main 中的 myValue 变量和函数中的 myValue 变量。

如果我们再次运行该程序，该函数现在会按照我们希望的方式工作。它改变了在 main 中分配的值的状态。

1. Addr: 0xc0000ae018 Value1 : 10 Value2: Bill
2. **Addr: 0xc0000ae018 Value1 : 20 Value2: Jill**
3. Addr: 0xc0000ae018 Value1 : 20 Value2: Jill


在函数调用期间，不再将值复制到堆栈上，而是复制值的地址。该函数现在通过局部指针变量引用相同的值，并更改这些值。

题为“Effective Go”的 Go 文档有一个关于内存分配的重要部分，其中包括数组、切片和映射的工作原理：

[http://golang.org/doc/effective_go.html#allocation_new](http://golang.org/doc/effective_go.html#allocation_new)

## 关键字 new 和 make。

new 关键字用于在内存中分配指定类型的值。内存分配被清零。无法在调用 new 时进一步初始化内存。换句话说，在使用 new 时，您不能为指定类型的属性指定特定值。

如果要在分配值时指定值，请使用复合文字。它们有两种风格，可以指定或不指定字段名称。
```
    // Allocate an value of type MyType
    // Values must be in the correct order
    myValue := MyType{10, "Bill"}

 
    // Allocate a value of type MyType
    // Use labeling to specify the values
    myValue := MyType{
        Value1: 10,
        Value2: "Bill",
    }
```

make 关键字仅用于分配和初始化切片、映射和通道。Make 不返回引用，它返回一个数据结构的“值”，该数据结构被创建和初始化以操作新的切片、映射或通道。此数据结构包含对用于操作切片、映射或通道的其他数据结构的引用。

将映射按值传递给函数是如何工作的。看看这个示例代码：
```go
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	myMap := make(map[string]string)
	myMap["Bill"] = "Jill"

	pointer := unsafe.Pointer(&myMap)
	fmt.Printf("1. Addr: %v Value : %s\n", pointer, myMap["Bill"])

	changeMyMap(myMap)
	fmt.Printf("3. Addr: %v Value : %s\n", pointer, myMap["Bill"])

	changeMyMapAddr(&myMap)
	fmt.Printf("5. Addr: %v Value : %s\n", pointer, myMap["Bill"])

}

// 传递为map数据结构,由于map是引用类型,只会复制数据结构,而不会复制底层数据,因此myMap参数具有与上次相同映射的变量,栈地址不一样而已
func changeMyMap(myMap map[string]string) {
	myMap["Bill"] = "Joan"
	pointer := unsafe.Pointer(&myMap)

	fmt.Printf("2. Addr: %v Value : %s\n", pointer, myMap["Bill"])
}

// 传递为map的引用地址,具有数据结构不会复制,而是传入指针地址
func changeMyMapAddr(myMapPointer *map[string]string) {
	(*myMapPointer)["Bill"] = "Jenny"
	pointer := unsafe.Pointer(myMapPointer)

	fmt.Printf("4. Addr: %v Value : %s\n", pointer, (*myMapPointer)["Bill"])
}

```

这是程序的输出：

1. Addr: 0xc00000e028 Value : Jill
2. **Addr: 0xc00000e038 Value : Joan**
3. Addr: 0xc00000e028 Value : Joan
4. Addr: 0xc00000e028 Value : Jenny
5. Addr: 0xc00000e028 Value : Jenny

我们初始化了一个map并添加了一个名为“Bill”的键来分配“Jill”的值。然后我们将 map 变量的值传递给 ChangeMyMap 函数。请记住，myMap 变量不是指针，因此在函数调用期间，myMap（一种数据结构）的“值”被复制到堆栈中。因为 myMap 的“值”是一个数据结构，它包含对映射内部的引用，函数可以使用它的数据结构副本对映射进行更改，函数调用后 main 将看到这些更改。

如果您查看输出，您会发现当我们按值传递map时，该函数拥有自己的map数据结构副本。您可以看到对map所做的更改在函数调用后得到反映。在 main 中，我们显示map键“Bill”的值

这是不必要的，但 ChangeMyMapAddr 函数显示了如何在 main 中传递和使用对 myMap 变量的引用。Go 团队再次确保传递map变量的“值”可以毫无问题地执行。请注意，当我们想要访问map时，我们需要如何取消引用 myMapPointer 变量。这是因为 Go 编译器不允许我们直接通过指针变量访问映射。取消引用一个指针变量相当于拥有一个其值为该值的变量。

我花时间写了这篇文章，因为有时可能会混淆变量的“值”包含什么。如果您的变量的“值”是一个大值，并且您将该变量的“值”传递给一个函数，那么您将在堆栈上制作该变量的一个大副本。您要确保将地址传递给您的函数，除非您有一个非常特殊的用例。

map、切片和通道是不同的。您可以毫无顾虑地按值传递这些变量。当我们将映射变量传递给函数时，我们复制的是数据结构而不是整个映射。

- 引用类型(map,slice,channel)可以按值传递给函数
- 内置类型除非有特殊用途,一般按值传递.
- 结构体类型一般使用指针传递,除非实现为基础类型(属性不需要变化)的话可以按照值传递
