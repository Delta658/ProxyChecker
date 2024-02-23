package main

import (
	"sync"

	"github.com/panjf2000/ants/v2"
)

var _Pool *ants.PoolWithFunc
var _wg *sync.WaitGroup

func CreateFool(count int) {
	_wg = &sync.WaitGroup{}
	pool, _ := ants.NewPoolWithFunc(count, func(i interface{}) {
		Work(i)
		_wg.Done()
	})
	_Pool = pool
}

func StartTask(count int) {
	for i := 0; i < count; i++ {
		_wg.Add(1)
		_Pool.Invoke(int(i))
	}
	_wg.Wait()
}
