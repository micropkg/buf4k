package buf4k

import "sync"

type Buffer4K []byte

var pool4k sync.Pool = sync.Pool{
	New: func() interface{} {
		return Buffer4K(make([]byte, 4096))
	},
}

func Get4K() Buffer4K {
	return pool4k.Get().(Buffer4K)
}

func Put4K(v []Buffer4K) {
	pool4k.Put(v)
}
