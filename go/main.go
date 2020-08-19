package main

import (
	. "stakeSrm/ftx"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	client := New(API, SECRET, SUBACCOUNT)
	wg.Add(1)
	go auto_staking(client, &wg)
	wg.Wait()
}
