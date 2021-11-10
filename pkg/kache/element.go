package kache

import (
	"time"
)

type element struct {
	hits    int
	value   []byte
	lastUse time.Time
}

func (v *element) read() []byte {
	v.hits++
	v.lastUse = time.Now()
	return v.value
}

func (v *element) set(newV []byte) {
	v.hits++
	v.value = newV
	v.lastUse = time.Now()
}
