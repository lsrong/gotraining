# Object Oriented Programming in Go
William KennedyJuly 10, 2013  

Someone asked a question on the forum today on how to gain the benefits of inheritance without embedding. It is really important for everyone to think in terms of Go and not the languages they are leaving behind. I can’t tell you much code I removed from my early Go implementations because it wasn’t necessary. The language designers have years of experience and knowledge. Hindsight is helping to create a language that is fast, lean and really fun to code in.  
今天论坛上有人问了一个问题，如何在不嵌入的情况下获得继承的好处。对于每个人来说，考虑 Go 而不是他们留下的语言，这真的很重要。我不能告诉你我从早期的 Go 实现中删除了多少代码，因为它不是必需的。语言设计者拥有多年的经验和知识。 Hindsight 正在帮助创建一种快速、精简且真正有趣的编程语言。

I consider Go to be a light object oriented programming language. Yes it does have encapsulation and type member functions but it lacks inheritance and therefore traditional polymorphism. For me, inheritance is useless unless you want to implement polymorphism. With the way interfaces are implemented in Go, there is no need for inheritance. Go took the best parts of OOP, left out the rest and gave us a better way to write polymorphic code.  
我认为 Go 是一种轻量级的面向对象的编程语言。是的，它确实有封装和类型成员函数，但它缺乏继承，因此缺乏传统的多态性。对我来说，除非你想实现多态，否则继承是没有用的。使用 Go 中接口的实现方式，不需要继承。 Go 继承了 OOP 最好的部分，忽略了其余部分，并为我们提供了一种编写多态代码的更好方法。

Here is a quick view of OOP in Go. Start with these three structs:  
这是 Go 中 OOP 的快速视图。从这三个结构开始：

```go
type Animal struct {
    Name string
    mean bool
}

type Cat struct {
    Basics Animal
    MeowStrength int
}

type Dog struct {
    Animal
    BarkStrength int
}
```

Here are three structs you would probably see in any OOP example. We have a base struct and two other structs that are specific to the base. The Animal structure contains attributes that all animals share and the other two structs are specific to cats and dogs.  
以下是您可能会在任何 OOP 示例中看到的三个结构。我们有一个基础结构和另外两个特定于基础的结构。 Animal 结构包含所有动物共享的属性，另外两个结构是特定于猫和狗的。

All of the member properties (fields) are public except for mean. The mean field in the Animal struct starts with a lowercase letter. In Go, the case of the first letter for variables, structs, fields, functions, etc. determine the access specification. Use a capital letter and it’s public, use a lowercase letter and it’s private.  
除均值外，所有成员属性（字段）都是公开的。 Animal 结构中的 mean 字段以小写字母开头。在 Go 中，变量、结构、字段、函数等的首字母大小写决定了访问规范。使用大写字母是公开的，使用小写字母是私人的。

*Note: The concept of public and private access in Go is not exactly true:*
*Note: Go 中公共和私有访问的概念并不完全正确*
*https://www.ardanlabs.com/blog/2014/03/exportedunexported-identifiers-in-go.html*

Since there is no inheritance in Go, composition is your only choice. The Cat struct has a field called Basics which is of type Animal. The Dog struct is using an un-named struct (embedding) for the Animal type. It’s up to you to decide which is better for you and I will show you both implementations.  
由于 Go 中没有继承，组合是你唯一的选择。 Cat 结构体有一个名为 Basics 的字段，其类型为 Animal。 Dog 结构体对 Animal 类型使用了一个未命名的结构体（嵌入）。由您决定哪个更适合您，我将向您展示这两种实现方式。

I want to thank John McLaughlin for his comment about un-named structs!!  
我要感谢 John McLaughlin 对未命名结构的评论！！

To create a member function (method) for both Cat and Dog, the syntax is as follows:  
为Cat和Dog创建成员函数（方法），语法如下：

```go
func (dog *Dog) MakeNoise(){
    barkStrength := dog.BarkStrength
    
    if dog.mean == true {
    barkStrength = barkStrength * 5
    }
    
    for bark := 0; bark < barkStrength; bark++ {
    fmt.Printf("BARK ")
}

func (cat *Cat) MakeNoise() {
    meowStrength := **cat**.MeowStrength

    if cat.Basics.mean == true {
        meowStrength = meowStrength * 5
    }

    for meow := 0; meow < meowStrength; meow++ {
        fmt.Printf("MEOW ")
    }

    fmt.Println("")
}
```

Before the name of the method we specify a receiver which is a pointer to each type. Now both Cat and Dog have methods called MakeNoise.  
在方法名称之前，我们指定了一个接收器，它是指向每种类型的指针。现在 Cat 和 Dog 都有名为 MakeNoise 的方法。

Both these methods do the same thing. Each animal speaks in their native tongue based on their bark or meow strength and if they are mean. Silly, but it shows you how to access the referenced object (value).  
这两种方法都做同样的事情。每只动物都会根据它们的吠声或喵叫声的强度以及它们是否刻薄，用它们的母语说话。这是一个简单的示例，但它向您展示了如何访问引用的对象（值）。

When using the Dog reference we access the Animal fields directly. With the Cat reference we use the named field called Basics.  
使用 Dog 引用时，我们直接访问 Animal 字段。对于 Cat 引用，我们使用名为 Basics 的命名字段。

So far we have covered encapsulation, composition, access specifications and member functions. All that is left is how to create polymorphic behavior.  
到目前为止，我们已经介绍了封装、组合、访问规范和成员函数。剩下的就是如何创建多态行为。

We use interfaces to create polymorphic behavior:  
我们使用接口来创建多态行为：
```go
type AnimalSounder interface {
    MakeNoise()
}

func MakeSomeNoise(animalSounder AnimalSounder) {
    animalSounder.MakeNoise()
}
```
Here we add an interface and a public function that takes a value of the interface type. Actually the function will take a reference to a value of a type that implements this interface. An interface is not a type that can be instantiated. An interface is a declaration of behavior that is implemented by other types.   
这里我们添加了一个接口和一个接受接口类型值的公共函数。实际上，该函数将引用实现此接口的类型的值。接口不是可以实例化的类型。接口是由其他类型实现的行为声明。


There is a Go convention of naming interfaces with the "er" suffix when the interface only contains one method.  
当接口只包含一个方法时，有一种用“er”后缀命名接口的 Go 约定。

In Go, any type that implements an interface, via methods, then represents the interface type. In our case both Cat and Dog have implemented the AnimalSounder interface with pointer receivers and therefore are considered to be of type AnimalSounder.   
在 Go 中，任何通过方法实现接口的类型都代表接口类型。在我们的例子中，Cat 和 Dog 都实现了带有指针接收器的 AnimalSounder 接口，因此被认为是 AnimalSounder 类型。

This means that pointers of both Cat and Dog can be passed as parameters to the MakeSomeNoise function. The MakeSomeNoise function implements polymorphic behavior through the AnimalSounder interface.   
这意味着 Cat 和 Dog 的指针都可以作为参数传递给 MakeSomeNoise 函数。 MakeSomeNoise 函数通过 AnimalSounder 接口实现多态行为。

If you wanted to reduce the duplication of code in the MakeNoise methods of Cat and Dog, you could create a method for the Animal type to handle it:  
如果您想减少 Cat 和 Dog 的 MakeNoise 方法中的代码重复，您可以为 Animal 类型创建一个方法来处理它：

```go
func (animal *Animal) PerformNoise(strength int, sound string) {
    if animal.mean == true {
        strength = strength * 5
    }

    for voice := 0; voice < strength; voice++ {
        fmt.Printf("%s ", sound)
    }

    fmt.Println("")
}

func (dog *Dog) MakeNoise() {
    dog.PerformNoise(dog.BarkStrength, "BARK")
}

func (cat *Cat) MakeNoise() {
    cat.Basics.PerformNoise(cat.MeowStrength, "MEOW")
}
```

Now the Animal type has a method with the business logic for making noise. The business logic stays within the scope of the type it belongs to. The other cool benefit is we don’t need to pass the mean field in as a parameter because it already belongs to the Animal type.  
现在，Animal 类型有一个带有制造噪音的业务逻辑的方法。业务逻辑保持在它所属类型的范围内。另一个很酷的好处是我们不需要将 mean 字段作为参数传入，因为它已经属于 Animal 类型。

Here is the complete working sample program:  
这是完整的工作示例程序：

```go
package main

import "fmt"

type Animal struct {
	Name string
	mean bool
}

type AnimalSounder interface {
	MakeNoise()
}

type Dog struct {
	Animal
	BarkStrength int
}

type Cat struct {
	Basics       Animal
	MeowStrength int
}

func (a *Animal) PerformNoise(strength int, sound string) {
	if a.mean {
		strength *= 5
	}
	for i := 0; i < strength; i++ {
		fmt.Printf("%s ", sound)
	}

	fmt.Println()
}

func (d *Dog) MakeNoise() {
	d.PerformNoise(d.BarkStrength, "BARK")
}

func (c *Cat) MakeNoise() {
	c.Basics.PerformNoise(c.MeowStrength, "MEOW")
}

func MakeSomeNoise(animalSounder AnimalSounder) {
	animalSounder.MakeNoise()
}

func main() {
	myDog := &Dog{
		Animal{
			"Rover",
			false,
		},
		2,
	}
	myCat := &Cat{
		Basics: Animal{
			Name: "Julius",
			mean: true,
		},
		MeowStrength: 3,
	}
	MakeSomeNoise(myDog)
	MakeSomeNoise(myCat)
}
```

Here is the output:
```shell
BARK BARK
MEOW MEOW MEOW MEOW MEOW MEOW MEOW MEOW MEOW MEOW MEOW MEOW MEOW MEOW MEOW
```

Someone posted an example on the board about embedding an interface inside of a struct. Here is an example:  
有人发布了一个关于在结构中嵌入接口的示例。下面是一个例子：

```go
package main

import (
"fmt"
)

type HornSounder interface {
    SoundHorn()
}

type Vehicle struct {
    List [2]HornSounder
}

type Car struct {
    Sound string
}

type Bike struct {
    Sound string
}

func main() {
    vehicle := new(Vehicle)
    vehicle.List[0] = &Car{"BEEP"}
    vehicle.List[1] = &Bike{"RING"}

    for _, hornSounder := range vehicle.List {
        hornSounder.SoundHorn()
    }
}

func (car *Car) SoundHorn() {
    fmt.Println(car.Sound)
}

func (bike *Bike) SoundHorn() {
    fmt.Println(bike.Sound)
}

func PressHorn(hornSounder HornSounder) {
    hornSounder.SoundHorn()
}
```



In this example the Vehicle struct maintains a list of values that implement the HornSounder interface. In main we create a new vehicle and assign a Car and Bike pointer to the list. This assignment is possible because Car and Bike both implement the interface. Then using a simple loop, we use the interface to sound the horn.  
在这个例子中，Vehicle 结构维护了一个实现 HornSounder 接口的值列表。在 main 中，我们创建了一个新车辆并将 Car 和 Bike 指针分配给列表。这种分配是可能的，因为 Car 和 Bike 都实现了接口。然后使用一个简单的循环，我们使用界面来使喇叭发声。

Everything you need to implement OOP in your application is there in Go. As I said before, Go took the best parts of OOP, left out the rest and gave us a better way to write polymorphic code.  
在应用程序中实现 OOP 所需的一切都在 Go 中。正如我之前所说的，Go 保留了 OOP 最好的部分，省略了其余部分，并为我们提供了一种编写多态代码的更好方法。

To learn more on related topics check out these posts:
要了解有关相关主题的更多信息，请查看以下帖子：

* [methods-interfaces-and-embedded-types](https://www.ardanlabs.com/blog/2014/05/methods-interfaces-and-embedded-types.html)
* [how-packages-work-in-go-language](https://www.ardanlabs.com/blog/2013/07/how-packages-work-in-go-language.html)
* [singleton-design-pattern-in-go](https://www.ardanlabs.com/blog/2013/07/singleton-design-pattern-in-go.html)