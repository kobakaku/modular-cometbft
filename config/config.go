package config

type NodeConfig struct {
	Light       bool
	DAAddress   string
	DAAuthToken string
	DAGasPrice  float64
	DANamespace string
}

var DefaultNodeConfig = NodeConfig{
	Light:       false,
	DAAddress:   "http://localhost:26658",
	DAGasPrice:  0,
	DANamespace: "DA Layer",
}
