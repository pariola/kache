package kache

type kache struct {
	size  int
	items map[interface{}]*element
}

// New
func New(size int) *kache {
	return &kache{
		size:  size,
		items: make(map[interface{}]*element),
	}
}

// evict
func (k *kache) evict() {
	panic("evict functionality not implemented")
}

// Set
func (k *kache) Set(key interface{}, value []byte) bool {

	// already exists
	if elem, ok := k.items[key]; ok {
		elem.set(value)
		return false
	}

	evict := len(k.items)+1 > k.size

	if evict {
		k.evict()
	}

	elem := new(element)
	elem.set(value)

	k.items[key] = elem

	return evict
}

// Get
func (k *kache) Get(key interface{}) ([]byte, bool) {

	v, ok := k.items[key]

	if !ok {
		return nil, ok
	}

	return v.read(), ok
}
