package main

import (
	"context"
	"fmt"
)

// Sample program to show how to store and retrieve values from a context.
// 展示如何从上下文中存储和检索值的示例程序。
/**
context.WithValue
func WithValue(parent Context, key, val interface{}) Context

WithValue returns a copy of parent in which the value associated with key is val.
WithValue 返回与 key 关联的值为 val 的 parent 的副本。

Use context Values only for request-scoped data that transits processes and APIs, not for passing optional parameters to functions.
仅将上下文值用于传输流程和 API 的请求范围数据，而不用于将可选参数传递给函数。

The provided key must be comparable and should not be of type string or any other built-in type to avoid collisions between packages using context.
提供的键必须是可比较的，并且不应该是字符串类型或任何其他内置类型，以避免使用上下文的包之间发生冲突。
Users of WithValue should define their own types for keys.
WithValue 的用户应该为键定义自己的类型。
To avoid allocating when assigning to an interface{}, context keys often have concrete type struct{}.
为避免在分配给 interface{} 时进行分配，上下文键通常具有具体类型 struct{}.
Alternatively, exported context key variables' static type should be a pointer or interface.
或者，导出的上下文键变量的静态类型应该是指针或接口。
*/

type TraceID string

type TraceIDKey int

func main() {
	// 定义参数键值
	traceID := TraceID("6ac8a78b-2533-4a31-aa0a-32e4e62f1de3")
	const traceKey TraceIDKey = 0

	// 将 traceID 值存储在上下文中，键类型的值为零。
	ctx := context.WithValue(context.Background(), traceKey, traceID)

	// 从 Context 值包中检索该 traceID 值。
	if uuid, ok := ctx.Value(traceKey).(TraceID); ok {
		fmt.Println("TraceID", uuid)
	}

	// 需要使用同一类定义的类型键才能获取context的值
	if _, ok := ctx.Value(0).(TraceID); !ok {
		fmt.Println("TraceID out found")
	}
}
