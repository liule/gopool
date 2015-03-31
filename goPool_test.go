package gopool

import (
	"strconv"
	"testing"
	"time"
)

func TestGoPool(t *testing.T) {
	if err := GoPoolInit(5, 10); err != nil {
		panic(err)
	}
	callback := func(task interface{}) {
		param, ok := task.(string)
		if !ok {
			println("param error")
			return
		}
		println("param", param)
	}
	GoPoolSetCallback(callback)
	GoPoolStart()
	for i := 0; i < 2000; i++ {
		GoPoolTaskAdd(strconv.Itoa(i))
		time.Sleep(1e5)
		if i%1000 == 0 {
			GoPoolAddProcess(5)
		}
	}
	GoPoolTotal()
}
