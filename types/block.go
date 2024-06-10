package types

type Block struct {
	Txs []Tx
}

// TODO: []byteで表現する
// TODO: pendingしているTXはmempoolの中にいれる
type Tx struct {
	isPending bool
}

func New() *Block {
	return &Block{}
}
