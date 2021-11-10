package main

import (
	"fmt"

	"kache/pkg/kache"
)

func main() {

	k := kache.New(3)

	k.Set("one", []byte("adiye"))
	k.Set(23, []byte("adiyea"))
	k.Set("23", []byte("adiyea"))

	fmt.Println(k.Get(23))

}
