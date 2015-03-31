package gopool

import (
	"log"
	"runtime"
	"sync"
)

var sumMap = make(map[int]int)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	log.SetFlags(log.Ltime | log.Lmicroseconds)
}

type GoPoolStruct struct {
	goMin    int
	goMax    int
	taskChan chan interface{}
	callback func(task interface{})
	sync.RWMutex
}

var goPoolStruct = &GoPoolStruct{}

// gorutine最大总数，0默认最大数
func GoPoolInit(goMin, goMax int) error {
	goPoolStruct.Lock()
	defer goPoolStruct.Unlock()
	goPoolStruct.goMin = goMin
	goPoolStruct.goMax = goMax
	goPoolStruct.taskChan = make(chan interface{}, goMax*2)
	return nil
}

func GoPoolStart() {
	for i := 0; i < goPoolStruct.goMin; i++ {
		go goPoolProcess(i)
	}
}

func GoPoolSetMax(max int) {
	goPoolStruct.Lock()
	defer goPoolStruct.Unlock()
	goPoolStruct.goMax = max
}

func GoPoolAddProcess(sum int) {
	for i := goPoolStruct.goMax; i < sum; i++ {
		go goPoolProcess(i)
	}
}

func GoPoolSetCallback(callback func(task interface{})) {
	goPoolStruct.Lock()
	defer goPoolStruct.Unlock()
	goPoolStruct.callback = callback
}

func goPoolProcess(i int) {
	for {
		task := <-goPoolStruct.taskChan
		println("goPoolProcess", i)
		goPoolStruct.callback(task)
		sumMap[i]++
	}
}

func GoPoolTaskAdd(task interface{}) {
	select {
	case goPoolStruct.taskChan <- task:
	default:
		log.Fatal("gorutine not enough")
	}
}

func GoPoolTotal() {
	for k, v := range sumMap {
		println("total:", k, v)
	}
}
