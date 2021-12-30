package web

type Middleware func(handler Handler) Handler

func wrapMiddleware(mws []Middleware, handler Handler) Handler {
	for i := len(mws) - 1; i >= 0; i-- {
		if h := mws[i]; h != nil {
			handler = h(handler)
		}
	}
	return handler
}
