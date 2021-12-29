package graph

type Uint64Stack []uint64

func (s Uint64Stack) Len() int { return len(s) }

func (s Uint64Stack) Top() uint64 { return s[len(s)-1] }
