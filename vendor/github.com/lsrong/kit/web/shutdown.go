package web

// 支持主动关闭服务的设计，应用可控制服务是否关闭。

import "errors"

// shutdownError 服务关闭错误类型。
type shutdownError struct {
	Message string
}

func (se *shutdownError) Error() string {
	return se.Message
}

// NewShutdownError 返回一个服务关闭错误实例。
func NewShutdownError(message string) error {
	return &shutdownError{message}
}

// IsShutdownError 是否为服务关闭错误。
func IsShutdownError(err error) bool {
	var se *shutdownError

	return errors.As(err, &se)
}
