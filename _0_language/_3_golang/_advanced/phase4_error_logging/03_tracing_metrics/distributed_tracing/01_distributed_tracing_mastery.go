// 🔍 DISTRIBUTED TRACING MASTERY
// Advanced distributed tracing with OpenTelemetry concepts
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"sync"
	"time"
)

// ============================================================================
// TRACE CONCEPTS
// ============================================================================

type TraceID string
type SpanID string

type TraceContext struct {
	TraceID      TraceID `json:"trace_id"`
	SpanID       SpanID  `json:"span_id"`
	ParentSpanID SpanID  `json:"parent_span_id,omitempty"`
	Sampled      bool    `json:"sampled"`
	Flags        int     `json:"flags"`
}

type Span struct {
	TraceID      TraceID                 `json:"trace_id"`
	SpanID       SpanID                  `json:"span_id"`
	ParentSpanID SpanID                  `json:"parent_span_id,omitempty"`
	Name         string                  `json:"name"`
	StartTime    time.Time               `json:"start_time"`
	EndTime      time.Time               `json:"end_time"`
	Duration     time.Duration           `json:"duration"`
	Status       SpanStatus              `json:"status"`
	Attributes   map[string]interface{}  `json:"attributes"`
	Events       []SpanEvent             `json:"events"`
	Links        []SpanLink              `json:"links"`
	Kind         SpanKind                `json:"kind"`
}

type SpanStatus struct {
	Code    StatusCode `json:"code"`
	Message string     `json:"message,omitempty"`
}

type SpanEvent struct {
	Name       string                 `json:"name"`
	Timestamp  time.Time              `json:"timestamp"`
	Attributes map[string]interface{} `json:"attributes"`
}

type SpanLink struct {
	TraceID TraceID                `json:"trace_id"`
	SpanID  SpanID                 `json:"span_id"`
	Attributes map[string]interface{} `json:"attributes"`
}

type SpanKind int

const (
	SpanKindUnspecified SpanKind = iota
	SpanKindInternal
	SpanKindServer
	SpanKindClient
	SpanKindProducer
	SpanKindConsumer
)

type StatusCode int

const (
	StatusCodeUnset StatusCode = iota
	StatusCodeOK
	StatusCodeError
)

// ============================================================================
// TRACER IMPLEMENTATION
// ============================================================================

type Tracer struct {
	name        string
	version     string
	spans       []Span
	mu          sync.RWMutex
	exporters   []SpanExporter
	sampler     Sampler
}

type SpanExporter interface {
	Export(spans []Span) error
	Shutdown() error
}

type Sampler interface {
	ShouldSample(traceID TraceID, spanID SpanID, name string) bool
}

type AlwaysOnSampler struct{}

func (s AlwaysOnSampler) ShouldSample(traceID TraceID, spanID SpanID, name string) bool {
	return true
}

type ProbabilisticSampler struct {
	probability float64
}

func NewProbabilisticSampler(probability float64) *ProbabilisticSampler {
	return &ProbabilisticSampler{probability: probability}
}

func (s ProbabilisticSampler) ShouldSample(traceID TraceID, spanID SpanID, name string) bool {
	return rand.Float64() < s.probability
}

func NewTracer(name, version string) *Tracer {
	return &Tracer{
		name:      name,
		version:   version,
		spans:     make([]Span, 0),
		exporters: make([]SpanExporter, 0),
		sampler:   AlwaysOnSampler{},
	}
}

func (t *Tracer) AddExporter(exporter SpanExporter) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.exporters = append(t.exporters, exporter)
}

func (t *Tracer) SetSampler(sampler Sampler) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.sampler = sampler
}

func (t *Tracer) StartSpan(ctx context.Context, name string, opts ...SpanOption) (context.Context, *Span) {
	// Extract trace context from parent context
	traceCtx := extractTraceContext(ctx)
	
	// Generate new span ID
	spanID := generateSpanID()
	
	// Check if we should sample this span
	if !t.sampler.ShouldSample(traceCtx.TraceID, spanID, name) {
		return ctx, nil
	}
	
	// Create new span
	span := &Span{
		TraceID:      traceCtx.TraceID,
		SpanID:       spanID,
		ParentSpanID: traceCtx.SpanID,
		Name:         name,
		StartTime:    time.Now(),
		Status:       SpanStatus{Code: StatusCodeUnset},
		Attributes:   make(map[string]interface{}),
		Events:       make([]SpanEvent, 0),
		Links:        make([]SpanLink, 0),
		Kind:         SpanKindInternal,
	}
	
	// Apply span options
	for _, opt := range opts {
		opt(span)
	}
	
	// Create new context with this span
	newCtx := context.WithValue(ctx, "span", span)
	newCtx = context.WithValue(newCtx, "trace_context", TraceContext{
		TraceID:      span.TraceID,
		SpanID:       span.SpanID,
		ParentSpanID: span.ParentSpanID,
		Sampled:      true,
	})
	
	return newCtx, span
}

func (t *Tracer) EndSpan(ctx context.Context, span *Span) {
	if span == nil {
		return
	}
	
	span.EndTime = time.Now()
	span.Duration = span.EndTime.Sub(span.StartTime)
	
	// Add span to collection
	t.mu.Lock()
	t.spans = append(t.spans, *span)
	spansToExport := make([]Span, len(t.spans))
	copy(spansToExport, t.spans)
	t.spans = t.spans[:0] // Clear for next batch
	t.mu.Unlock()
	
	// Export spans
	for _, exporter := range t.exporters {
		exporter.Export(spansToExport)
	}
}

func (t *Tracer) AddEvent(span *Span, name string, attrs map[string]interface{}) {
	if span == nil {
		return
	}
	
	event := SpanEvent{
		Name:       name,
		Timestamp:  time.Now(),
		Attributes: attrs,
	}
	
	span.Events = append(span.Events, event)
}

func (t *Tracer) SetAttributes(span *Span, attrs map[string]interface{}) {
	if span == nil {
		return
	}
	
	for k, v := range attrs {
		span.Attributes[k] = v
	}
}

func (t *Tracer) SetStatus(span *Span, code StatusCode, message string) {
	if span == nil {
		return
	}
	
	span.Status = SpanStatus{
		Code:    code,
		Message: message,
	}
}

// ============================================================================
// SPAN OPTIONS
// ============================================================================

type SpanOption func(*Span)

func WithSpanKind(kind SpanKind) SpanOption {
	return func(s *Span) {
		s.Kind = kind
	}
}

func WithAttributes(attrs map[string]interface{}) SpanOption {
	return func(s *Span) {
		for k, v := range attrs {
			s.Attributes[k] = v
		}
	}
}

func WithLinks(links []SpanLink) SpanOption {
	return func(s *Span) {
		s.Links = links
	}
}

// ============================================================================
// EXPORTERS
// ============================================================================

type ConsoleExporter struct {
	output io.Writer
}

func NewConsoleExporter(output io.Writer) *ConsoleExporter {
	return &ConsoleExporter{output: output}
}

func (ce *ConsoleExporter) Export(spans []Span) error {
	for _, span := range spans {
		jsonData, err := json.MarshalIndent(span, "", "  ")
		if err != nil {
			return err
		}
		fmt.Fprintf(ce.output, "SPAN: %s\n", string(jsonData))
	}
	return nil
}

func (ce *ConsoleExporter) Shutdown() error {
	return nil
}

type FileExporter struct {
	filename string
	file     *os.File
	mu       sync.Mutex
}

func NewFileExporter(filename string) *FileExporter {
	return &FileExporter{filename: filename}
}

func (fe *FileExporter) Export(spans []Span) error {
	fe.mu.Lock()
	defer fe.mu.Unlock()
	
	if fe.file == nil {
		file, err := os.OpenFile(fe.filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			return err
		}
		fe.file = file
	}
	
	for _, span := range spans {
		jsonData, err := json.Marshal(span)
		if err != nil {
			continue
		}
		fe.file.Write(append(jsonData, '\n'))
	}
	
	return nil
}

func (fe *FileExporter) Shutdown() error {
	fe.mu.Lock()
	defer fe.mu.Unlock()
	
	if fe.file != nil {
		return fe.file.Close()
	}
	return nil
}

// ============================================================================
// HTTP TRACING MIDDLEWARE
// ============================================================================

type HTTPTracingMiddleware struct {
	tracer *Tracer
}

func NewHTTPTracingMiddleware(tracer *Tracer) *HTTPTracingMiddleware {
	return &HTTPTracingMiddleware{tracer: tracer}
}

func (htm *HTTPTracingMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Extract trace context from headers
		traceCtx := extractTraceContextFromHeaders(r.Header)
		ctx := context.WithValue(r.Context(), "trace_context", traceCtx)
		
		// Start span
		ctx, span := htm.tracer.StartSpan(ctx, r.URL.Path, 
			WithSpanKind(SpanKindServer),
			WithAttributes(map[string]interface{}{
				"http.method":     r.Method,
				"http.url":        r.URL.String(),
				"http.user_agent": r.UserAgent(),
				"http.remote_addr": r.RemoteAddr,
			}),
		)
		
		if span != nil {
			defer htm.tracer.EndSpan(ctx, span)
		}
		
		// Create response writer wrapper
		wrapped := &responseWriter{ResponseWriter: w, statusCode: 200}
		
		// Call next handler
		next.ServeHTTP(wrapped, r.WithContext(ctx))
		
		// Set span attributes
		if span != nil {
			htm.tracer.SetAttributes(span, map[string]interface{}{
				"http.status_code": wrapped.statusCode,
			})
			
			// Set status
			if wrapped.statusCode >= 400 {
				htm.tracer.SetStatus(span, StatusCodeError, fmt.Sprintf("HTTP %d", wrapped.statusCode))
			} else {
				htm.tracer.SetStatus(span, StatusCodeOK, "")
			}
		}
	})
}

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

// ============================================================================
// UTILITY FUNCTIONS
// ============================================================================

func generateTraceID() TraceID {
	return TraceID(fmt.Sprintf("%016x", rand.Uint64()))
}

func generateSpanID() SpanID {
	return SpanID(fmt.Sprintf("%08x", rand.Uint32()))
}

func extractTraceContext(ctx context.Context) TraceContext {
	if traceCtx, ok := ctx.Value("trace_context").(TraceContext); ok {
		return traceCtx
	}
	
	// Generate new trace context
	return TraceContext{
		TraceID: generateTraceID(),
		SpanID:  generateSpanID(),
		Sampled: true,
	}
}

func extractTraceContextFromHeaders(headers http.Header) TraceContext {
	traceID := TraceID(headers.Get("X-Trace-ID"))
	spanID := SpanID(headers.Get("X-Span-ID"))
	parentSpanID := SpanID(headers.Get("X-Parent-Span-ID"))
	
	if traceID == "" {
		traceID = generateTraceID()
	}
	if spanID == "" {
		spanID = generateSpanID()
	}
	
	return TraceContext{
		TraceID:      traceID,
		SpanID:       spanID,
		ParentSpanID: parentSpanID,
		Sampled:      true,
	}
}

// ============================================================================
// DEMONSTRATION FUNCTIONS
// ============================================================================

func demonstrateBasicTracing() {
	fmt.Println("\n=== Basic Tracing ===")
	
	// Create tracer
	tracer := NewTracer("demo-service", "1.0.0")
	
	// Add console exporter
	tracer.AddExporter(NewConsoleExporter(os.Stdout))
	
	// Start root span
	ctx := context.Background()
	ctx, span := tracer.StartSpan(ctx, "root-operation",
		WithSpanKind(SpanKindServer),
		WithAttributes(map[string]interface{}{
			"service.name": "demo-service",
			"operation.type": "user-request",
		}),
	)
	
	if span != nil {
		// Add some events
		tracer.AddEvent(span, "operation.started", map[string]interface{}{
			"user_id": "123",
		})
		
		// Simulate some work
		time.Sleep(10 * time.Millisecond)
		
		// Add more events
		tracer.AddEvent(span, "database.query", map[string]interface{}{
			"query": "SELECT * FROM users",
			"duration_ms": 5,
		})
		
		// Set status
		tracer.SetStatus(span, StatusCodeOK, "")
		
		// End span
		tracer.EndSpan(ctx, span)
	}
}

func demonstrateNestedSpans() {
	fmt.Println("\n=== Nested Spans ===")
	
	// Create tracer
	tracer := NewTracer("demo-service", "1.0.0")
	tracer.AddExporter(NewConsoleExporter(os.Stdout))
	
	// Start root span
	ctx := context.Background()
	ctx, rootSpan := tracer.StartSpan(ctx, "user-request",
		WithSpanKind(SpanKindServer),
		WithAttributes(map[string]interface{}{
			"user.id": "456",
			"request.id": "req-789",
		}),
	)
	
	if rootSpan != nil {
		// Start child span
		ctx, childSpan := tracer.StartSpan(ctx, "database-operation",
			WithSpanKind(SpanKindClient),
			WithAttributes(map[string]interface{}{
				"db.operation": "SELECT",
				"db.table": "users",
			}),
		)
		
		if childSpan != nil {
			// Simulate database work
			time.Sleep(5 * time.Millisecond)
			tracer.SetStatus(childSpan, StatusCodeOK, "")
			tracer.EndSpan(ctx, childSpan)
		}
		
		// Start another child span
		ctx, childSpan2 := tracer.StartSpan(ctx, "external-api-call",
			WithSpanKind(SpanKindClient),
			WithAttributes(map[string]interface{}{
				"http.method": "GET",
				"http.url": "https://api.example.com/users/456",
			}),
		)
		
		if childSpan2 != nil {
			// Simulate API call
			time.Sleep(15 * time.Millisecond)
			tracer.SetStatus(childSpan2, StatusCodeOK, "")
			tracer.EndSpan(ctx, childSpan2)
		}
		
		// End root span
		tracer.SetStatus(rootSpan, StatusCodeOK, "")
		tracer.EndSpan(ctx, rootSpan)
	}
}

func demonstrateHTTPTracing() {
	fmt.Println("\n=== HTTP Tracing ===")
	
	// Create tracer
	tracer := NewTracer("http-service", "1.0.0")
	tracer.AddExporter(NewConsoleExporter(os.Stdout))
	
	// Create HTTP handler
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get span from context
		if span, ok := r.Context().Value("span").(*Span); ok && span != nil {
			tracer.AddEvent(span, "handler.started", map[string]interface{}{
				"path": r.URL.Path,
			})
		}
		
		// Simulate some work
		time.Sleep(5 * time.Millisecond)
		
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, World!"))
		
		if span, ok := r.Context().Value("span").(*Span); ok && span != nil {
			tracer.AddEvent(span, "handler.completed", map[string]interface{}{
				"response_size": 13,
			})
		}
	})
	
	// Add tracing middleware
	tracingMiddleware := NewHTTPTracingMiddleware(tracer)
	wrappedHandler := tracingMiddleware.Middleware(handler)
	
	// Create test request
	req, _ := http.NewRequest("GET", "/api/users", nil)
	req.Header.Set("X-Trace-ID", "trace-123")
	req.Header.Set("X-Span-ID", "span-456")
	
	// Create response writer
	w := &responseWriter{ResponseWriter: &mockResponseWriter{}, statusCode: 200}
	
	// Serve request
	wrappedHandler.ServeHTTP(w, req)
}

type mockResponseWriter struct{}

func (mrw *mockResponseWriter) Header() http.Header {
	return make(http.Header)
}

func (mrw *mockResponseWriter) Write(data []byte) (int, error) {
	return len(data), nil
}

func (mrw *mockResponseWriter) WriteHeader(code int) {}

func demonstrateSampling() {
	fmt.Println("\n=== Sampling ===")
	
	// Create tracer with probabilistic sampling
	tracer := NewTracer("sampled-service", "1.0.0")
	tracer.SetSampler(NewProbabilisticSampler(0.5)) // 50% sampling
	tracer.AddExporter(NewConsoleExporter(os.Stdout))
	
	// Create multiple spans
	for i := 0; i < 10; i++ {
		ctx := context.Background()
		ctx, span := tracer.StartSpan(ctx, fmt.Sprintf("operation-%d", i))
		
		if span != nil {
			// Simulate work
			time.Sleep(1 * time.Millisecond)
			tracer.SetStatus(span, StatusCodeOK, "")
			tracer.EndSpan(ctx, span)
		}
	}
	
	fmt.Println("   📊 Only ~50% of spans were sampled and exported")
}

func demonstrateFileExport() {
	fmt.Println("\n=== File Export ===")
	
	// Create tracer
	tracer := NewTracer("file-service", "1.0.0")
	
	// Add file exporter
	fileExporter := NewFileExporter("traces.json")
	tracer.AddExporter(fileExporter)
	defer fileExporter.Shutdown()
	
	// Create some spans
	ctx := context.Background()
	ctx, span := tracer.StartSpan(ctx, "file-operation",
		WithAttributes(map[string]interface{}{
			"file.name": "example.txt",
			"file.size": 1024,
		}),
	)
	
	if span != nil {
		time.Sleep(2 * time.Millisecond)
		tracer.SetStatus(span, StatusCodeOK, "")
		tracer.EndSpan(ctx, span)
	}
	
	fmt.Println("   📊 Traces exported to traces.json")
}

func main() {
	fmt.Println("🔍 DISTRIBUTED TRACING MASTERY")
	fmt.Println("==============================")
	
	demonstrateBasicTracing()
	demonstrateNestedSpans()
	demonstrateHTTPTracing()
	demonstrateSampling()
	demonstrateFileExport()
	
	fmt.Println("\n🎉 DISTRIBUTED TRACING MASTERY COMPLETE!")
	fmt.Println("You have successfully demonstrated:")
	fmt.Println("✅ Basic trace and span concepts")
	fmt.Println("✅ Nested span hierarchies")
	fmt.Println("✅ HTTP request tracing")
	fmt.Println("✅ Trace sampling strategies")
	fmt.Println("✅ File-based trace export")
	
	fmt.Println("\n🚀 You are now ready for Prometheus Metrics Mastery!")
}
