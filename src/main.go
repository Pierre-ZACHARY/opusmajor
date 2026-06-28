package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/grafana/pyroscope-go/godeltaprof/http/pprof"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	httpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests handled by this service.",
		},
		[]string{"method", "path", "status"},
	)

	httpRequestDurationSeconds = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "HTTP request duration in seconds.",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "path", "status"},
	)

	playerDataRequestsTotal = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "player_data_requests_total",
			Help: "Total number of calls to /player-data.",
		},
	)
)

func init() {
	prometheus.MustRegister(httpRequestsTotal, httpRequestDurationSeconds, playerDataRequestsTotal)
}

func main() {
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(requestMetricsAndLogs())

	// Expose standard Go pprof endpoints in pull mode under /debug/pprof/*.
	router.Any("/debug/pprof/*any", gin.WrapH(http.DefaultServeMux))
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	router.GET("/health", func(c *gin.Context) {
		log.Printf("health check hit")
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	router.GET("/player-data", func(c *gin.Context) {
		playerDataRequestsTotal.Inc()
		log.Printf("player data requested")
		c.JSON(http.StatusOK, gin.H{
			"player": "demo-player",
			"level":  7,
			"status": "active",
		})
	})

	// Gin listens on 0.0.0.0:8080 by default.
	_ = router.Run()
}

func requestMetricsAndLogs() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.FullPath()
		if path == "" {
			path = c.Request.URL.Path
		}

		c.Next()

		statusCode := c.Writer.Status()
		status := http.StatusText(statusCode)
		if status == "" {
			status = "unknown"
		}

		httpRequestsTotal.WithLabelValues(c.Request.Method, path, status).Inc()
		httpRequestDurationSeconds.WithLabelValues(c.Request.Method, path, status).Observe(time.Since(start).Seconds())

		log.Printf("method=%s path=%s status=%d duration_ms=%d client_ip=%s",
			c.Request.Method,
			path,
			statusCode,
			time.Since(start).Milliseconds(),
			c.ClientIP(),
		)
	}
}
