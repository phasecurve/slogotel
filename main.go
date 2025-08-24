package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.opentelemetry.io/contrib/bridges/otelslog"
	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploggrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploghttp"
	sdklog "go.opentelemetry.io/otel/sdk/log"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	insecure := os.Getenv("OTEL_EXPORTER_OTLP_INSECURE") == "true"
	useHTTP := os.Getenv("OTEL_EXPORTER_OTLP_PROTOCOL") == "http"

	var processor *sdklog.BatchProcessor

	if useHTTP {
		opts := []otlploghttp.Option{}
		if insecure {
			opts = append(opts, otlploghttp.WithInsecure())
		}
		exp, err := otlploghttp.New(ctx, opts...)
		if err != nil {
			slog.Error("failed to create OTLP HTTP log exporter", "err", err)
			os.Exit(1)
		}
		processor = sdklog.NewBatchProcessor(exp)
	} else {
		opts := []otlploggrpc.Option{}
		if insecure {
			opts = append(opts, otlploggrpc.WithInsecure())
		}
		exp, err := otlploggrpc.New(ctx, opts...)
		if err != nil {
			slog.Error("failed to create OTLP gRPC log exporter", "err", err)
			os.Exit(1)
		}
		processor = sdklog.NewBatchProcessor(exp)
	}

	provider := sdklog.NewLoggerProvider(sdklog.WithProcessor(processor))
	defer func() {
		shCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		_ = provider.Shutdown(shCtx)
	}()

	handler := otelslog.NewHandler("otelslog", otelslog.WithLoggerProvider(provider))
	logger := slog.New(handler)

	logger.Info("logger initialised", "protocol", map[bool]string{true: "http", false: "grpc"}[useHTTP], "insecure", insecure)

	<-ctx.Done()
	logger.Info("shutting down")
}
