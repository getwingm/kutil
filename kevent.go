package kutil

import (
	"time"
)

type KEvent struct {
	ch chan int
}

func (k *KEvent) init() {
	if k.ch == nil {
		k.ch = make(chan int)
	}
}

/*0：表示超时返回，> 0表示事件返回。
 */
func (k *KEvent) Wait(t int64) int {

	k.init()

	for {
		select {
		case <-time.After(time.Duration(t)):
			return 0
		case v := <-k.ch:
			return v
		}
	}
}

/*确保v值必须大于0
 */
func (k *KEvent) Post(v int) {
	k.init()
	if v <= 0 {
		v = 1
	}
	k.ch <- v
}

func (k *KEvent) Close() {
	if k.ch != nil {
		close(k.ch)
	}
}
