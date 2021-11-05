## Pointers vs. Values - effective go
As we saw with ByteSize, methods can be defined for any named type (except a pointer or an interface); the receiver does not have to be a struct.    
正如我们在 ByteSize 中看到的，可以为任何命名类型（指针或接口除外）定义方法；接收者不必是结构体。

In the discussion of slices above, we wrote an Append function. We can define it as a method on slices instead. To do this, we first declare a named type to which we can bind the method, and then make the receiver for the method a value of that type.   
在上面对切片的讨论中，我们编写了一个 Append 函数。我们可以将其定义为切片上的方法。为此，我们首先声明一个可以绑定方法的命名类型，然后使该方法的接收器成为该类型的值。

````go
type ByteSlice []byte

// Append ByteSlice值方法,操作对象为源值的副本,因此需要返回[]byte
func (slice ByteSlice) Append(data []byte)[]byte{
	l := len(slice)
	n := len(data)
	// 重新分配容量
	if l + n > cap(slice){
		new := make([]byte, (l + n) * 2)
		copy(new, slice)
		slice = new
    }
    slice = slice[0:l+n]
    copy(slice[l:], data)
    
    return slice
}

````
This still requires the method to return the updated slice. We can eliminate that clumsiness by redefining the method to take a pointer to a ByteSlice as its receiver, so the method can overwrite the caller's slice.  
这仍然需要返回更新后的切片的方法。我们可以通过重新定义方法以将指向 ByteSlice 的指针作为其接收者来消除这种笨拙，因此该方法可以覆盖调用者的切片。

```go
func (b *ByteSlice)Append(data []byte){
	slice := *b
	l := len(slice)
    n := len(data)
    // 重新分配容量
    if l + n > cap(slice){
        new := make([]byte, (l + n) * 2)
        copy(new, slice)
        slice = new
    }
    slice = slice[0:l+n]
    copy(slice[l:], data)
    
    // 不需要返回
	*p = slice
}
```

In fact, we can do even better. If we modify our function so it looks like a standard Write method, like this,  
事实上，我们还可以做得更好。如果我们修改我们的函数，让它看起来像一个标准的 Write 方法，像这样，
```go
func (p *ByteSlice) Write(data []byte) (n int, err error) {
    slice := *p
    // 省略细节,与上面实现一样
    *p = slice
    return len(data), nil
}
```

then the type *ByteSlice satisfies the standard interface io.Writer, which is handy. For instance, we can print into one.  
那么ByteSlice类型满足标准接口io.Writer，方便。例如，我们可以打印成一个。
```go
    var b ByteSlice
    fmt.Fprintf(&b, "This hour has %d days\n", 7)
```

We pass the address of a ByteSlice because only *ByteSlice satisfies io.Writer. The rule about pointers vs. values for receivers is that value methods can be invoked on pointers and values, but pointer methods can only be invoked on pointers.  
我们传递一个 ByteSlice 的地址，因为只有 *ByteSlice(指针) 满足 io.Writer。关于接收者的指针与值的规则是值方法可以在指针和值上调用，但指针方法只能在指针上调用。

This rule arises because pointer methods can modify the receiver; invoking them on a value would cause the method to receive a copy of the value, so any modifications would be discarded. The language therefore disallows this mistake. There is a handy exception, though. When the value is addressable, the language takes care of the common case of invoking a pointer method on a value by inserting the address operator automatically. In our example, the variable b is addressable, so we can call its Write method with just b.Write. The compiler will rewrite that to (&b).Write for us.  
**出现这个规则是因为指针方法可以修改接收者；对一个值调用它们会导致该方法接收该值的副本，因此任何修改都将被丢弃**。因此，该语言不允许这种错误。但是，有一个方便的例外。当值可寻址时，该语言会通过自动插入地址运算符来处理对值调用指针方法的常见情况。在我们的例子中，变量 b 是可寻址的，所以我们可以只用 b.Write 调用它的 Write 方法。编译器会将其重写为 (&b).Write 。

By the way, the idea of using Write on a slice of bytes is central to the implementation of bytes.Buffer.  
顺便说一下，在字节切片上使用 Write 的想法是 bytes.Buffer 实现的核心。