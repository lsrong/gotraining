package main

// This is an example that creates interface pollution
// by improperly using an interface when one is not needed.

// 这是一个在不需要的时候不正确地使用接口造成接口污染的例子。

type Server interface {
	Start() error
	Stop() error
	Wait() error
}

type server struct {
	host string
}

func NewServer(host string) Server {
	return &server{host: host}
}

func (s *server) Start() error {
	// implementation
	return nil
}
func (s *server) Stop() error {
	// implementation
	return nil
}
func (s *server) Wait() error {
	// implementation
	return nil
}

func main() {
	srv := NewServer("localhost")
	srv.Stop()
	srv.Stop()
	srv.Wait()
}

// =============================================================================

// NOTES:

// Smells:
//  * The package declares an interface that matches the entire API of its own concrete type.
//  * The interface is exported but the concrete type is unexported.
//  * The factory function returns the interface value with the unexported concrete type value inside.
//  * The interface can be removed and nothing changes for the user of the API.
//  * The interface is not decoupling the API from change.

// 该包声明了一个与其自身具体类型的整个 API 相匹配的接口。
// 接口已导出，但具体类型未导出。
// 工厂函数返回带有未导出的具体类型值的接口值。
// 该接口可以被移除，并且对 API 的用户没有任何改变。
// 该接口并未将 API 与更改解耦.
