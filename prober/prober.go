package prober

import (
	"context"

	"github.com/go-kit/kit/log"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/prometheus/blackbox_exporter/config"
	"net/http"
)

type ProbeFn func(r *http.Request,ctx context.Context, target string, config config.Module, registry *prometheus.Registry, logger log.Logger) bool
