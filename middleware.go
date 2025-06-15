package tracer

import (
	"net/http"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

func AddTracingMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			span := trace.SpanFromContext(r.Context())
			span.AddEvent("got request", trace.WithAttributes(attribute.KeyValue{
				Key:   attribute.Key("url"),
				Value: attribute.StringValue(r.URL.Path),
			}))
			defer span.End()
			handler.ServeHTTP(w, r)
		},
	)
}
