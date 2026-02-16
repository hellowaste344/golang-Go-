package main

import (
	"fmt"
	"sync"
)

var m = map[int]int{}

var mu = &sync.RWMutex{}

func writeLoop(m map[int]int, mu *sync.RWMutex){
	for{
		for i:= 0; i<100; i++{
			mu.Lock()
			m[i] = 1
			mu.Unlock()
		}
	}
}


func readLoop(m map[int]int, mu *sync.RWMutex){
	for{
		mu.RLock()
		for k, v:= range m{
			fmt.Println(k,"-",v)
		}
		mu.RUnlock()
	}
}