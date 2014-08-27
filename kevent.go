package kutil

import (
	"time"
)

type KEvent struct {
	ch chan int
}

func (k *KEvent) init() {
	if k.ch == nil {
		k.ch = make(chan int, 1)
	}
}

/*0：表示超时返回，> 0表示事件返回。
 */
func (k *KEvent) Wait(ms int) int {

	k.init()

	for {
		select {
		case <-time.After(time.Millisecond * time.Duration(ms)):
			return 0
		case v := <-k.ch:
			return v
		}
	}
}

/*确保v值必须大于0，t：为写入超时返回
 */
func (k *KEvent) Post(v int, ms int) bool {
	k.init()
	if v <= 0 {
		v = 1
	}
	for {
		select {
		case <-time.After(time.Millisecond * time.Duration(ms)):
			return false
		case k.ch <- v:
			return true
		}
	}
}

func (k *KEvent) Close() {
	if k.ch != nil {
		close(k.ch)
		k.ch = nil
	}
}
