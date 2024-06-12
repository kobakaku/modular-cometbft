package types

type Block struct {
	Header Header
	Txs    []Tx
}

// TODO: []byteで表現する
// TODO: pendingしているTXはmempoolの中にいれる
type Tx struct {
	isPending bool
}

func New() *Block {
	return new(Block)
}

func (b *Block) Height() uint64 {
	return b.Header.BaseHeader.Height
}
