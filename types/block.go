package types

type Block struct {
	Header Header
	Txs    []Tx
}

type Tx []byte

func New() *Block {
	return new(Block)
}

func (b *Block) Height() uint64 {
	return b.Header.BaseHeader.Height
}
