package main

// Sample program demonstrating when implicit interface conversions are provided by the compiler.
// 演示何时由编译器提供隐式接口转换的示例程序。
import "fmt"

type Mover interface {
	Move()
}

type Locker interface {
	Lock()
	Unlock()
}

type MoveLocker interface {
	Mover
	Locker
}

type bike struct{}

func (bike) Move() {
	fmt.Println("Moving bike")
}

func (bike) Lock() {
	fmt.Println("Locking bike")
}

func (bike) Unlock() {
	fmt.Println("Unlocking bike")
}

func main() {
	var ml MoveLocker
	var m Mover

	ml = bike{}

	// An interface value of type MoveLocker can be implicitly converted into
	// a value of type Mover. They both declare a method named move.
	// MoveLocker 类型的接口值可以隐式转换为 Mover 类型的值。
	// 它们都声明了一个名为 move 的方法。
	// ml（MoveLocker） 包含了m （Mover）的接口方法，会自动隐式转换成 Mover类型
	m = ml

	//	ml = m
	// ./convertions.go:40:5: cannot use m (type Mover) as type MoveLocker in assignment:
	//	Mover does not implement MoveLocker (missing Lock method)

	// Interface type Mover does not declare methods named lock and unlock.
	// Therefore, the compiler can't perform an implicit conversion to assign
	// a value of interface type Mover to an interface value of type MoveLocker.
	// It is irrelevant that the concrete type value of type bike that is stored
	// inside of the Mover interface value implements the MoveLocker interface.

	// 接口类型 Mover 没有声明名为 lock 和 unlock 的方法。
	// 因此，编译器无法执行隐式转换以将接口类型 Mover 的值分配给 MoveLocker 类型的接口值。
	// 存储在 Mover 接口值中的类型自行车的具体类型值实现了 MoveLocker 接口是无关紧要的。

	// We can perform a type assertion at runtime to support the assignment.

	// Perform a type assertion against the Mover interface value to access
	// a COPY of the concrete type value of type bike that was stored inside
	// of it. Then assign the COPY of the concrete type to the MoveLocker
	// interface.
	// 对 Mover 接口值执行类型断言，以访问存储在其中的自行车类型的具体类型值的副本。
	// 然后将具体类型的COPY赋值给MoveLocker接口。
	// 将m(Mover)转换成具体的实现类型bike, 在将bike复制给MoveLocker类型变量.
	b := m.(bike)
	ml = b

	// It's important to note that the type assertion syntax provides a way
	// to state what type of value is stored inside the interface. This is
	// more powerful from a language and readability standpoint, than using
	// a casting syntax, like in other languages.
	// 需要注意的是，类型断言语法提供了一种方法来说明接口中存储的值的类型。
	// 从语言和可读性的角度来看，这比在其他语言中使用强制转换语法更强大。
}
