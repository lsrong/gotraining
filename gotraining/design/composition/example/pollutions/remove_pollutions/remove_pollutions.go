package main

type Server struct {
	host string
}

func NewServer(host string) *Server {
	return &Server{host: host}
}

func (s *Server) Start() error {
	// implementation
	return nil
}
func (s *Server) Stop() error {
	// implementation
	return nil
}
func (s *Server) Wait() error {
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

// Here are some guidelines around interface pollution:
// * Use an interface:
//      * When users of the API need to provide an implementation detail.
//      * When API’s have multiple implementations that need to be maintained.
//      * When parts of the API that can change have been identified and require decoupling.
// * Question an interface:
//      * When its only purpose is for writing testable API’s (write usable API’s first).
//      * When it’s not providing support for the API to decouple from change.
//      * When it's not clear how the interface makes the code better.

// 以下是有关接口污染的一些指南：
//		使用接口：
//		当 API 的用户需要提供实现细节时。
//		当 API 有多个需要维护的实现时。
//		当 API 中可以更改的部分已被识别并需要解耦时。
//		质疑接口：
//		当它的唯一目的是编写可测试的 API 时（首先编写可用的 API）。
//		当它不为 API 提供支持以与变化解耦时。
//		当不清楚接口如何使代码更好时。
