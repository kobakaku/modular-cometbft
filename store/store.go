package store

type Store struct {
	// TODO: heightが非同期に処理できていないから、毎回初期化されてしまう。
	height uint64
}

func New() *Store {
	return &Store{
		height: 0,
	}
}

func (s *Store) SetHeight(height uint64) {
	s.height = height
}

func (s *Store) Height() uint64 {
	return s.height
}
