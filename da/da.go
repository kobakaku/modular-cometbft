package da

import (
	goDA "github.com/rollkit/go-da"
)

type DAClient struct {
	DA goDA.DA
}

func NewDAClient(da goDA.DA) *DAClient {
	return &DAClient{
		DA: da,
	}
}
