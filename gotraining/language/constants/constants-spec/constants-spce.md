# Constants 常量

常量有 *布尔(bool)常量* 、  *符文(rune)常量* 、  *整数(integer)常量* 、  *浮点(float)常量* 、*复数(complex)常量*和 *字符串(string)常量* 。符文、整数、浮点数和复数常量统称为 *数字常量* 。

常量值由 [符文(Rune)](https://golang.org/ref/spec#Rune_literals) 、 [整数(Integer)](https://golang.org/ref/spec#Integer_literals) 、 [浮点数(float)](https://golang.org/ref/spec#Floating-point_literals) 、 [虚数(Imaginary)](https://golang.org/ref/spec#Imaginary_literals) 或 [字符串(string)](https://golang.org/ref/spec#String_literals) 文字、表示常量的标识符、[常量表达式](https://golang.org/ref/spec#Constant_expressions) 、结果为常量的[转换](https://golang.org/ref/spec#Conversions) 或某些内置函数的结果值表示功能，例如 `unsafe.Sizeof`应用到任何值， `cap`或`len`施加到 [一些表达式](https://golang.org/ref/spec#Length_and_capacity) ， `real`并且`imag`施加到一个复常数和`complex`施加到数字常数。布尔真值由预先声明的常量 `true`和 表示`false`。预先声明的标识符 [iota](https://golang.org/ref/spec#Iota) 表示整数常量。

通常，复数常量是[常量表达式的](https://golang.org/ref/spec#Constant_expressions) 一种形式， 将在该部分中进行讨论。

数字常量代表任意精度的精确值并且不会溢出。因此，没有表示 IEEE-754 负零、无穷大和非数字值的常数。

常量可以是有[类型的](https://golang.org/ref/spec#Types) 或 *无类型的* 。文字常量、`true`、`false`、`iota`和某些 仅包含无类型常量操作数的[常量表达式](https://golang.org/ref/spec#Constant_expressions) 是无类型的。

常量可以通过[常量声明](https://golang.org/ref/spec#Constant_declarations) 或[转换](https://golang.org/ref/spec#Conversions) 显式指定类型，或者在[变量声明](https://golang.org/ref/spec#Variable_declarations) 或 [赋值中使用时](https://golang.org/ref/spec#Assignments) 或作为[表达式中](https://golang.org/ref/spec#Expressions) 的操作数时隐式 指定。如果常量值不能[表示](https://golang.org/ref/spec#Representability) 为相应类型的值，则会出错。

非类型化常数具有一个 *默认类型* ，其是其中恒定隐式转换在需要类型化值，例如上下文，在该类型[短变量声明](https://golang.org/ref/spec#Short_variable_declarations) 如`i := 0`在没有明确的类型。一个无类型恒定的默认类型是`bool`，`rune`， `int`，`float64`，`complex128`或`string` 分别，这取决于它是否是一个布尔值，符，整数，浮点，复杂，或字符串常量。

实现限制：尽管数字常量在语言中具有任意精度，但编译器可以使用精度有限的内部表示来实现它们。也就是说，每个实现都必须：

* 表示至少 256 位的整数常量。
* 表示浮点常量，包括复数常量的部分，尾数至少为 256 位，带符号的二进制指数至少为 16 位。
* 如果无法精确表示整数常量，则给出错误。
* 如果由于溢出而无法表示浮点或复数常量，则给出错误。
* 如果由于精度限制而无法表示浮点或复数常量，则舍入到最接近的可表示常量。

这些要求既适用于文字常量，也适用于计算[常量表达式](https://golang.org/ref/spec#Constant_expressions) 的结果。
