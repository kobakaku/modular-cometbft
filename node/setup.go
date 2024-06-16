package node

import (
	cmcfg "github.com/cometbft/cometbft/config"
	"github.com/cometbft/cometbft/proxy"
)

type MetricsProvider func(chainID string) *proxy.Metrics

func DefaultMetricsProvider(cmcfg *cmcfg.InstrumentationConfig) MetricsProvider {
	return func(chainID string) *proxy.Metrics {
		if cmcfg.Prometheus {
			return proxy.PrometheusMetrics(cmcfg.Namespace, "chain_id", chainID)
		}
		return proxy.NopMetrics()
	}
}
