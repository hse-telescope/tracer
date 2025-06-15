package tracer

import (
	"net/http"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

func AddTracingMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			ctx, span := tracer.Start(r.Context(), r.URL.Path)
			span.AddEvent("got request", trace.WithAttributes(attribute.KeyValue{
				Key:   attribute.Key("url"),
				Value: attribute.StringValue(r.URL.Path),
			}))
			defer span.End()
			r = r.WithContext(ctx)
			handler = otelhttp.NewHandler(handler, "smth")
			handler.ServeHTTP(w, r)
		},
	)
}
