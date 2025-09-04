package balancer

import (
	ip_hash "github.com/PrathamNabira/Load-balancer-qa/core/ip-hash"
	least_algorithm "github.com/PrathamNabira/Load-balancer-qa/core/least-algorithm"
	random "github.com/PrathamNabira/Load-balancer-qa/core/random"
	round_robin "github.com/PrathamNabira/Load-balancer-qa/core/round-robin"
	"github.com/PrathamNabira/Load-balancer-qa/core/types"
	w_round_robin "github.com/PrathamNabira/Load-balancer-qa/core/w-round-robin"

	"github.com/PrathamNabira/Load-balancer-qa/internal/proxy"

	"github.com/PrathamNabira/Load-balancer-qa/pkg/config"
)

var balancers = map[string]func(config *config.Config, proxyFunc proxy.ProxyFunc) types.IBalancer{
	"round-robin":         round_robin.NewRoundRobin,
	"w-round-robin":       w_round_robin.NewWRoundRobin,
	"ip-hash":             ip_hash.NewIPHash,
	"random":              random.NewRandom,
	"least-connection":    least_algorithm.NewLeastAlgorithm,
	"least-response-time": least_algorithm.NewLeastAlgorithm,
}

func NewBalancer(config *config.Config, proxyFunc proxy.ProxyFunc) types.IBalancer {
	return balancers[config.Type](config, proxyFunc)
}
