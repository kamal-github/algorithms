package main

import (
	"fmt"
	"sync"
	"time"
)

type anyFn *func() interface{}

type Throttler struct {
	m *sync.Map
}

func NewThrottler() *Throttler {
	return &Throttler{
		m: &sync.Map{},
	}
}

func (t *Throttler) throttle(fn anyFn, waitSec int64, retCh chan interface{}) {
	lastExecTimeI, loaded := t.m.LoadOrStore(fn, time.Now())
	if !loaded {
		go func() {
			retCh <- (*fn)()
		}()
		return
	}

	lastExecTime := lastExecTimeI.(time.Time)
	now := time.Now()

	if waitSec > (now.Unix() - lastExecTime.Unix()) {
		time.Sleep(time.Duration(waitSec-(now.Unix()-lastExecTime.Unix())) * time.Second)
	}

	go func() {
		retCh <- (*fn)()
		t.m.Store(fn, time.Now())
		return
	}()
}

func foo() interface{} {
	return 0
}

func main() {
	myFoo := foo
	ch := make(chan interface{})
	defer close(ch)

	t := NewThrottler()
	go t.throttle(&myFoo, 10, ch)
	go t.throttle(&myFoo, 20, ch)
	go t.throttle(&myFoo, 30, ch)

	for i := 0; i < 3; i++ {
		fmt.Println("got it", <-ch)
	}
}
