package metrics

import (
	"net/http"
	"strconv"

	"github.com/micro-in-cn/starter-kit/pkg/plugin/utils/request"
	"github.com/micro-in-cn/starter-kit/pkg/plugin/utils/response"
	"github.com/micro/micro/v3/plugin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

//DefObjectives of prometheus
var (
	DefObjectives = map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001}
)

//Prometheus for metircs
type Prometheus struct {
	options *Options
}

// Size returns the size of request object.
func Size(r *http.Request) float64 {
	size := 0
	if r.URL != nil {
		size = len(r.URL.String())
	}

	size += len(r.Method)
	size += len(r.Proto)

	for name, values := range r.Header {
		size += len(name)
		for _, value := range values {
			size += len(value)
		}
	}
	size += len(r.Host)

	// r.Form and r.MultipartForm are assumed to be included in r.URL.
	if r.ContentLength != -1 {
		size += int(r.ContentLength)
	}
	return float64(size)
}

func (p *Prometheus) handler(h http.Handler) http.Handler {
	opts := p.options
	md := make(map[string]string)

	labels := []string{"host"}
	reqTotalCounter := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: opts.namespace,
			Subsystem: opts.subsystem,
			Name:      "request_total",
			Help:      "Total request count.",
		},
		[]string{"host", "status"},
	)

	reqDurSummary := prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Namespace:  opts.namespace,
			Subsystem:  opts.subsystem,
			Name:       "request_latency_seconds",
			Help:       "Request latencies in seconds.",
			Objectives: DefObjectives,
		},
		labels,
	)

	reqDurHistogram := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: opts.namespace,
			Subsystem: opts.subsystem,
			Name:      "request_duration_seconds",
			Help:      "Request time in seconds.",
		},
		labels,
	)

	reqSizeSummary := prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Namespace:  opts.namespace,
			Subsystem:  opts.subsystem,
			Name:       "request_size_bytes",
			Help:       "Request size in bytes.",
			Objectives: DefObjectives,
		},
		labels,
	)

	respSizeSummary := prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Namespace:  opts.namespace,
			Subsystem:  opts.subsystem,
			Name:       "response_size_bytes",
			Help:       "Response size in bytes.",
			Objectives: DefObjectives,
		},
		labels,
	)

	reg := prometheus.NewRegistry()
	wrapreg := prometheus.WrapRegistererWith(md, reg)
	wrapreg.MustRegister(
		prometheus.NewProcessCollector(prometheus.ProcessCollectorOpts{}),
		prometheus.NewGoCollector(),
		reqTotalCounter,
		reqDurSummary,
		reqDurHistogram,
		reqSizeSummary,
		respSizeSummary,
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
		if p.options.skipperFunc(r) {
			h.ServeHTTP(w, r)
			return
		}

		values := []string{r.Host}
		timer := prometheus.NewTimer(prometheus.ObserverFunc(func(v float64) {
			reqDurSummary.WithLabelValues(values...).Observe(v)
			reqDurHistogram.WithLabelValues(values...).Observe(v)
		}))
		defer timer.ObserveDuration()

		ww := response.WrapWriter{ResponseWriter: w}
		h.ServeHTTP(&ww, r)

		reqSizeSummary.WithLabelValues(values...).Observe(float64(request.Size(r)))
		respSizeSummary.WithLabelValues(values...).Observe(float64(ww.Size))
		reqTotalCounter.WithLabelValues(r.Host, strconv.Itoa(ww.StatusCode)).Inc()
	})
}

func newPrometheus(opts ...Option) plugin.Plugin {
	options := newOptions(opts...)
	p := Prometheus{options: &options}

	return plugin.NewPlugin(
		plugin.WithName("metrics"),
		plugin.WithHandler(p.handler),
	)
}
