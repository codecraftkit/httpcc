package middlewares

import "net/http"

func Apply(mux http.Handler, middlewares []func(next http.Handler) http.Handler) http.Handler {
	for i := len(middlewares) - 1; i >= 0; i-- {
		mux = middlewares[i](mux)
	}
	return mux
}
