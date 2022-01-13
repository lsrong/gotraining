package baike

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/learning_golang/topics/packages/context/baike.baidu/userip"
	"net/http"
)

type Result struct {
	Title, Abstract, URL string
	Errno                int
}

// Search 发起一个百度百科的请求,并返回结果
func Search(ctx context.Context, query string) (Result, error) {
	var (
		req    *http.Request
		result Result
		err    error
	)

	// 准备百度百科 api request
	req, err = http.NewRequest("GET", "http://baike.baidu.com/api/openapi/BaikeLemmaCardApi?spm=a2c6h.12873639.0.0.2f852de9arCLOL&scope=103&format=json&appid=379020", nil)
	if err != nil {
		return result, err
	}

	q := req.URL.Query()
	q.Set("bk_key", query)

	// 如果ctx携带用户ip,传递给api请求
	if userIP, ok := userip.FromContext(ctx); ok {
		q.Set("userip", userIP.String())
	}
	req.URL.RawQuery = q.Encode()

	// 处理百度百科接口请求
	err = httpDo(ctx, req, func(resp *http.Response, err error) error {
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
			return err
		}

		return nil
	})
	if result.Errno > 0 {
		return result, fmt.Errorf("baike error %d", result.Errno)
	}
	return result, err
}

// httpDo 发起http请求,并调用处理结果函数, 当http.Do运行过程中ctx.Done 关闭,则返回ctx.Err,
// 否则等待http.Do执行完成返回处理函数的error
func httpDo(ctx context.Context, r *http.Request, f func(r *http.Response, err error) error) error {
	c := make(chan error)
	// 请求增加上下文,可以实现超时,主动取消等机制
	req := r.WithContext(ctx)
	// 异步处理结果
	go func() {
		c <- f(http.DefaultClient.Do(req))
	}()

	// 监听
	select {
	case <-ctx.Done():
		<-c // 等待f函数返回
		return ctx.Err()

	case err := <-c:
		return err
	}
}
