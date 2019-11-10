package metrics

import (
	"net/http"
	"strconv"

	"github.com/micro/go-micro/util/log"
	"github.com/micro/micro/plugin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type wrapWriter struct {
	StatusCode int
	http.ResponseWriter
}

func (ww *wrapWriter) WriteHeader(statusCode int) {
	log.Debugf("statusCode: %d", statusCode)
	ww.StatusCode = statusCode
	ww.ResponseWriter.WriteHeader(statusCode)
}

func newPlugin(opts ...Option) plugin.Plugin {
	options := newOptions(opts...)

	return plugin.NewPlugin(
		plugin.WithName("metrics"),
		plugin.WithHandler(
			func(h http.Handler) http.Handler {
				md := make(map[string]string)

				opsCounter := prometheus.NewCounterVec(
					prometheus.CounterOpts{
						Namespace: "micro",
						Name:      "request_total",
						Help:      "How many go-micro requests processed, partitioned by method and status",
					},
					[]string{"path", "method", "code"},
				)

				timeCounterSummary := prometheus.NewSummaryVec(
					prometheus.SummaryOpts{
						Namespace: "micro",
						Name:      "upstream_latency_microseconds",
						Help:      "Service backend method request latencies in microseconds",
					},
					[]string{"path", "method"},
				)

				timeCounterHistogram := prometheus.NewHistogramVec(
					prometheus.HistogramOpts{
						Namespace: "micro",
						Name:      "request_duration_seconds",
						Help:      "Service method request time in seconds",
					},
					[]string{"path", "method"},
				)

				reg := prometheus.NewRegistry()
				wrapreg := prometheus.WrapRegistererWith(md, reg)
				wrapreg.MustRegister(
					opsCounter,
					timeCounterSummary,
					timeCounterHistogram,
				)

				prometheus.DefaultGatherer = reg
				prometheus.DefaultRegisterer = wrapreg

				return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					// 拦截metrics path，默认"/metrics"
					if r.URL.Path == "/metrics" {
						promhttp.Handler().ServeHTTP(w, r)
						return
					}

					// 静态资源等不需要监控的请求，可以实现SkipperFunc过滤
					if options.skipperFunc(r) {
						h.ServeHTTP(w, r)
						return
					}

					path := r.URL.Path
					method := r.Method
					timer := prometheus.NewTimer(prometheus.ObserverFunc(func(v float64) {
						us := v * 1000000 // make microseconds
						timeCounterSummary.WithLabelValues(path, method).Observe(us)
						timeCounterHistogram.WithLabelValues(path, method).Observe(v)
					}))
					defer timer.ObserveDuration()

					ww := wrapWriter{ResponseWriter: w}
					h.ServeHTTP(&ww, r)
					log.Debugf("statusCode: %d, %s", ww.StatusCode, strconv.Itoa(ww.StatusCode))
					opsCounter.WithLabelValues(path, method, strconv.Itoa(ww.StatusCode)).Inc()
				})
			}),
	)
}

func NewPlugin(opts ...Option) plugin.Plugin {
	return newPlugin(opts...)
}
