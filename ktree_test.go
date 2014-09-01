package kutil

import (
	"fmt"
	"testing"
)

type MyT struct {
	text string
}

func (k *MyT) FunCompare(o interface{}) int64 {
	v, ok := o.(*MyT)
	if !ok {
		return 0
	}
	if k.text == v.text {
		return 0
	}
	if k.text < v.text {
		return -1
	}
	return 1
}

func test3(t *testing.T) {
	myt := &MyT{"aaaa"}
	MyWrap(myt)
	fmt.Println(myt.Thankyou(), myt.Willcome())
	if my1, ok1 := interface{}(myt).(MyTOne); ok1 {
		fmt.Println(my1.Thankyou(), ok1)
	}
	if my2, ok2 := interface{}(myt).(MyTTwo); ok2 {
		fmt.Println(my2.Willcome(), ok2)
	}
}

func test5(t *testing.T) {
	tree := &KRbtree{}
	tree.Add("dsf")
	tree.Add("2")
	tree.Add("4.3")
	tree.Walk(func(v interface{}) {
		val, ok := v.(string)
		fmt.Println(val, ok)
	})
}

func test7(t *testing.T) {
	tree := &KRbtree{}
	tree.Add(&MyT{"dsfdsfds"})
	tree.Add(&MyT{"dsfdaasfds"})
	tree.Add(&MyT{"dsccfdsfds"})
	tree.Walk(func(v interface{}) {
		val, ok := v.(*MyT)
		fmt.Println(val.text, ok)
	})
}
