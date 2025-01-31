package router

import (
	"taskmanager/internal/config"
	"taskmanager/internal/controller"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sony/gobreaker"
)

var (
	requestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "taskmanager_requests_total",
			Help: "Total number of requests",
		},
		[]string{"method", "endpoint"},
	)
	requestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "taskmanager_request_duration_seconds",
			Help:    "Histogram of response times for each endpoint",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "endpoint"},
	)
)

func init() {
	prometheus.MustRegister(requestsTotal)
	prometheus.MustRegister(requestDuration)
}

func CircuitBreakerMiddleware(cb *gobreaker.CircuitBreaker) gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := cb.Execute(func() (interface{}, error) {
			c.Next()
			return nil, nil
		})

		if err != nil {
			c.JSON(503, gin.H{"error": "Service unavailable"})
			c.Abort()
			return
		}
	}
}

func MetricsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		timer := prometheus.NewTimer(
			requestDuration.WithLabelValues(c.Request.Method, c.FullPath()),
		)
		defer timer.ObserveDuration()

		requestsTotal.WithLabelValues(c.Request.Method, c.FullPath()).Inc()
		c.Next()
	}
}

func SetupRoutes(r *gin.Engine) {
	cbConfig := config.AppConfig.CircuitBreaker
	cb := gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:        "TaskServiceCircuitBreaker",
		MaxRequests: uint32(cbConfig.MaxRequests),
		Interval:    cbConfig.Interval,
		Timeout:     cbConfig.Timeout,
	})

	r.Use(MetricsMiddleware())
	r.Use(CircuitBreakerMiddleware(cb))

	if config.AppConfig.Metrics.Enabled {
		r.GET(config.AppConfig.Metrics.Endpoint, gin.WrapH(promhttp.Handler()))
	}

	r.GET("/tasks", controller.GetTasks)
	r.POST("/tasks", controller.CreateTask)
	r.GET("/tasks/:id", controller.GetTaskByID)
	r.PUT("/tasks/:id", controller.UpdateTask)
	r.DELETE("/tasks/:id", controller.DeleteTask)
}
