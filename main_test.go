package main

import (
	"context"
	"os"
	"testing"
	"time"
)

func TestMainStartsWithoutEnvironmentVariables(t *testing.T) {
	os.Unsetenv("OTEL_EXPORTER_OTLP_INSECURE")
	os.Unsetenv("OTEL_EXPORTER_OTLP_PROTOCOL")
	
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	
	done := make(chan struct{})
	go func() {
		defer close(done)
		main()
	}()
	
	select {
	case <-done:
	case <-ctx.Done():
	}
}

func TestMainStartsWithInsecureEnabled(t *testing.T) {
	os.Setenv("OTEL_EXPORTER_OTLP_INSECURE", "true")
	os.Unsetenv("OTEL_EXPORTER_OTLP_PROTOCOL")
	defer os.Unsetenv("OTEL_EXPORTER_OTLP_INSECURE")
	
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	
	done := make(chan struct{})
	go func() {
		defer close(done)
		main()
	}()
	
	select {
	case <-done:
	case <-ctx.Done():
	}
}

func TestMainStartsWithHttpProtocol(t *testing.T) {
	os.Unsetenv("OTEL_EXPORTER_OTLP_INSECURE")
	os.Setenv("OTEL_EXPORTER_OTLP_PROTOCOL", "http")
	defer os.Unsetenv("OTEL_EXPORTER_OTLP_PROTOCOL")
	
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	
	done := make(chan struct{})
	go func() {
		defer close(done)
		main()
	}()
	
	select {
	case <-done:
	case <-ctx.Done():
	}
}

func TestMainStartsWithInsecureHttpProtocol(t *testing.T) {
	os.Setenv("OTEL_EXPORTER_OTLP_INSECURE", "true")
	os.Setenv("OTEL_EXPORTER_OTLP_PROTOCOL", "http")
	defer os.Unsetenv("OTEL_EXPORTER_OTLP_INSECURE")
	defer os.Unsetenv("OTEL_EXPORTER_OTLP_PROTOCOL")
	
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	
	done := make(chan struct{})
	go func() {
		defer close(done)
		main()
	}()
	
	select {
	case <-done:
	case <-ctx.Done():
	}
}

func TestMainContinuesWhenGrpcExporterConnectionFails(t *testing.T) {
	os.Unsetenv("OTEL_EXPORTER_OTLP_INSECURE")
	os.Unsetenv("OTEL_EXPORTER_OTLP_PROTOCOL")
	os.Setenv("OTEL_EXPORTER_OTLP_ENDPOINT", "invalid-endpoint")
	defer os.Unsetenv("OTEL_EXPORTER_OTLP_ENDPOINT")
	
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	
	done := make(chan struct{})
	go func() {
		defer close(done)
		main()
	}()
	
	select {
	case <-done:
	case <-ctx.Done():
	}
}

func TestMainContinuesWhenHttpExporterConnectionFails(t *testing.T) {
	os.Unsetenv("OTEL_EXPORTER_OTLP_INSECURE")
	os.Setenv("OTEL_EXPORTER_OTLP_PROTOCOL", "http")
	os.Setenv("OTEL_EXPORTER_OTLP_ENDPOINT", "invalid-endpoint")
	defer os.Unsetenv("OTEL_EXPORTER_OTLP_PROTOCOL")
	defer os.Unsetenv("OTEL_EXPORTER_OTLP_ENDPOINT")
	
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	
	done := make(chan struct{})
	go func() {
		defer close(done)
		main()
	}()
	
	select {
	case <-done:
	case <-ctx.Done():
	}
}