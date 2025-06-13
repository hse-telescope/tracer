package tracer

import "net/http"

func AddTracingMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			tracer.Start(r.Context(), r.URL.Path)
			handler.ServeHTTP(w, r)
		},
	)
}
