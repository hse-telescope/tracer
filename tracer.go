package tracer

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/resource"
	sdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.opentelemetry.io/otel/trace"
)

var (
	tracer         trace.Tracer = otel.Tracer("tracer")
	tracerProvider *sdk.TracerProvider
)

func SetupTracer(ctx context.Context, serviceName string, otlpTracesURL string) error {
	otlpTracesExporter, err := otlptracegrpc.New(
		ctx,
		otlptracegrpc.WithEndpointURL(otlpTracesURL),
	)
	if err != nil {
		return err
	}

	resource, err := resource.New(
		context.Background(),
		resource.WithSchemaURL(semconv.SchemaURL),
		resource.WithAttributes(semconv.ServiceNameKey.String(serviceName)),
	)
	if err != nil {
		return err
	}

	tracerProvider = sdk.NewTracerProvider(
		sdk.WithBatcher(otlpTracesExporter),
		sdk.WithResource(resource),
	)

	otel.SetTracerProvider(tracerProvider)

	tracer = tracerProvider.Tracer(serviceName)

	return nil
}

func ShutdownTracer(ctx context.Context) error {
	return tracerProvider.Shutdown(ctx)
}

func Start(ctx context.Context, message string, opts ...trace.SpanStartOption) (context.Context, trace.Span) {
	return tracer.Start(ctx, message, opts...)
}

func SpanFromContext(ctx context.Context) trace.Span {
	return trace.SpanFromContext(ctx)
}
