package config

type NodeConfig struct {
	Light       bool
	DAAddress   string
	DAAuthToken string
}

var DefaultNodeConfig = NodeConfig{
	Light:     false,
	DAAddress: "http://localhost:26658",
}
