## Interface and Composition Design 接口和组合设计

Composition goes beyond the mechanics of type embedding and is more than just a paradigm. It is the key for maintaining stability in your software by having the ability to adapt to the data and transformation changes that are coming.  

组合超越了类型嵌入的机制，而不仅仅是一种范式。通过适应即将到来的数据和转换变化的能力，它是保持软件稳定性的关键。

## Notes

* This is much more than the mechanics of type embedding.
* 这不仅仅是类型嵌入的机制。
* 
* Declare types and implement workflows with composition in mind.
* 声明类型并实施工作流时考虑到组合。
* 
* Understand the problem you are trying to solve first. This means understanding the data.
* 首先了解您要解决的问题。这意味着理解数据。
* 
* The goal is to reduce and minimize cascading changes across your software.
* 目标是减少和最小化整个软件的级联更改。
* 
* Interfaces provide the highest form of composition.
* 接口提供了最高形式的组合
* 
* Don't group types by a common DNA but by a common behavior.
* 不要根据共同的 属性(DNA) 来分组类型，而是根据共同的行为。
* 
* Everyone can work together when we focus on what we do and not what we are.
* 当我们专注于我们所做的而不是我们自己时，每个人都可以一起工作。

## Quotes

_"A good API is not just easy to use but also hard to misuse." - JBD_

_"You can always embed, but you cannot decompose big interfaces once they are out there. Keep interfaces small." - JBD_

_"Don't design with interfaces, discover them." - Rob Pike_

_"Duplication is far cheaper than the wrong abstraction." - Sandi Metz_

## Links

[Repeat yourself, do more than one thing, and rewrite everything](https://programmingisterrible.com/post/176657481103/repeat-yourself-do-more-than-one-thing-and) - tef  
[Embedding](https://golang.org/doc/effective_go.html#embedding)   
[Methods, Interfaces and Embedding](https://www.ardanlabs.com/blog/2014/05/methods-interfaces-and-embedded-types.html) - William Kennedy  
[Composition In Go](https://www.ardanlabs.com/blog/2015/09/composition-with-go.html) - William Kennedy  
[Reducing Type Hierarchies](https://www.ardanlabs.com/blog/2016/10/reducing-type-hierarchies.html) - William Kennedy  
[Avoid Interface Pollution](https://www.ardanlabs.com/blog/2016/10/avoid-interface-pollution.html) - William Kennedy

## Code Review

#### Grouping Types

[Grouping By State](example/grouping/grouping_state/grouping_state.go) ([Go Playground](https://play.golang.org/p/Dh_cCEz3o0N))  
[Grouping By Behavior](example/grouping/grouping_behavior/group_behavior.go) ([Go Playground](https://play.golang.org/p/wRpHBoPu79K))

#### Decoupling

[Struct Composition](example/decoupling/struct_composition/struct_composition.go) ([Go Playground](https://play.golang.org/p/7nsTaKtlMWc))  
[Decoupling With Interface](example/decoupling/decoupling_interface/decouping_interface.go) ([Go Playground](https://play.golang.org/p/HjP8V6ScpKi))  
[Interface Composition](example/decoupling/interface_composition/interface_composition.go) ([Go Playground](https://play.golang.org/p/EDbjyjjpxAi))  
[Decoupling With Interface Composition](example/decoupling/decoupling_interface_composition/decoupling_interface_composition.go) ([Go Playground](https://play.golang.org/p/zeO2cJLP46B))  
[Remove Interface Pollution](example/decoupling/interface_pollution/interface_pollution.go) ([Go Playground](https://play.golang.org/p/Kg4JKGwJGGy))  
[More Precise API](example/decoupling/precise_api/precise_api.go) ([Go Playground](https://play.golang.org/p/cdvbrsgclGX))

#### Conversion and Assertions

[Interface Conversions](example/assertions/conversions/conversions.go) ([Go Playground](https://play.golang.org/p/sNP3bMR1kc-))  
[Runtime Type Assertions](example/assertions/assertions/assertions.go) ([Go Playground](https://play.golang.org/p/PtdQOc9xZ7S))  
[Behavior Changes](example/assertions/behavior/behavior.go) ([Go Playground](https://play.golang.org/p/AYhH8yXDcuy))

#### Interface Pollution

[Create Interface Pollution](example/pollutions/create_pollutions/create_pollutions.go) ([Go Playground](https://play.golang.org/p/DCqTbY14loz))  
[Remove Interface Pollution](example/pollutions/remove_pollutions/remove_pollutions.go) ([Go Playground](https://play.golang.org/p/K3w2eX7V1j2))

## Exercises

### Exercise 1

Using the template, declare a set of concrete types that implement the set of predefined interface types. Then create values of these types and use them to complete a set of predefined tasks.

[Template](exercise/template1/template1.go) ([Go Playground](https://play.golang.org/p/6Kp_E3Wim0G)) |
[Answer](exercise/exercise1/exercise1.go) ([Go Playground](https://play.golang.org/p/hwUADsRfnax))
