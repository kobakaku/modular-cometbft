package types

// BaseHeader contains the most basic data of a header/
type BaseHeader struct {
	Height uint64
}

// Header defined the structure of block header/
type Header struct {
	BaseHeader

	// TODO: cometbft上の任意のheaderを記述する
}

func (h *Header) Height() uint64 {
	return h.Height()
}
