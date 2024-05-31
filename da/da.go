package da

import (
	goDA "github.com/rollkit/go-da"
)

type DAClient struct {
	DA        goDA.DA
	GasPrice  float64
	Namespace goDA.Namespace
}

func NewDAClient(da goDA.DA, gasPrice float64, ns goDA.Namespace) *DAClient {
	return &DAClient{
		DA:        da,
		GasPrice:  gasPrice,
		Namespace: ns,
	}
}
