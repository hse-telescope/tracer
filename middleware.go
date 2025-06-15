package tracer

import (
	"net/http"

	"go.opentelemetry.io/otel/propagation"
)

func AddTracingMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			p := propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{})
			ctx := p.Extract(r.Context(), propagation.HeaderCarrier(r.Header))

			ctx, span := tracer.Start(ctx, r.URL.Path)
			defer span.End()

			r = r.WithContext(ctx)
			handler.ServeHTTP(w, r)
		},
	)
}
