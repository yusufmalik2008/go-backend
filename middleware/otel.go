// middleware/otel.go
func OTELTracer() fiber.Handler {
	tracer := otel.Tracer("taskflow")
	return func(c *fiber.Ctx) error {
		ctx, span := tracer.Start(c.Context(), c.Path())
		defer span.End()
		c.Locals("trace", span)
		return c.Next()
	}
}

// metrics
var (
	httpRequestsTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{Name: "http_requests_total"},
		[]string{"path", "method", "status"},
	)
)

// di handler
httpRequestsTotal.WithLabelValues(c.Path(), c.Method(), "200").Inc()
