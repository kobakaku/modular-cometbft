package config

type NodeConfig struct {
	Light bool
}

var DefaultNodeConfig = NodeConfig{
	Light: false,
}
