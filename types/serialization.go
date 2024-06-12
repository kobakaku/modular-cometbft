package types

import (
	"bytes"
	"encoding/gob"
)

// MarshalBinary encodes Block into binary format.
func (b *Block) MarshalBinary(block *Block) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	err := gob.NewEncoder(buf).Encode(block)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// UnmarshalBinary decodes binary format of Block into Block.
func (b *Block) UnmarshalBinary(data []byte) (*Block, error) {
	var block *Block
	buf := bytes.NewBuffer(data)
	err := gob.NewDecoder(buf).Decode(&block)
	if err != nil {
		return block, err
	}
	return block, nil
}
