package config

type NodeConfig struct {
	DBPath string

	Light       bool
	DAAddress   string
	DAAuthToken string
	DAGasPrice  float64
	DANamespace string
}

var DefaultNodeConfig = NodeConfig{
	DBPath: "data",

	Light:       false,
	DAAddress:   "http://localhost:26658",
	DAGasPrice:  0,
	DANamespace: "DA Layer",
}
