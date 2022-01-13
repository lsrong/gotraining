// Package userip provides functions for extracting a user IP address from a
// request and associating it with a Context.
// userip 提供了从请求中提取用户 IP 地址并将其与上下文相关联的功能。
package userip

import (
	"context"
	"fmt"
	"net"
	"net/http"
)

type key int

const userIPKey key = 0

// FromRequest 从请求中解析用户的IP地址
func FromRequest(r *http.Request) (net.IP, error) {
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return nil, fmt.Errorf("userip:%q is not IP:port", r.RemoteAddr)
	}
	userIP := net.ParseIP(ip)
	if userIP == nil {
		return nil, fmt.Errorf("userip:%q is not IP:port", r.RemoteAddr)
	}

	return userIP, nil
}

// NewContext 生成带有用户 IP 的派生上下文
func NewContext(ctx context.Context, userIP net.IP) context.Context {
	return context.WithValue(ctx, userIPKey, userIP)
}

// FromContext 在上下文中解析用户IP
func FromContext(ctx context.Context) (net.IP, bool) {
	userIP, ok := ctx.Value(userIPKey).(net.IP)

	return userIP, ok
}
