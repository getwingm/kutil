package kutil

import (
	"fmt"
	"reflect"
)

type KValue interface {
	FunCompare(o interface{}) int64 //0:equal,-1:small,1:big
}

type wrap_value interface {
	FunCompare(o wrap_value) int64
	Value() interface{}
}

type KRbtree struct {
	left  *KRbtree
	value wrap_value
	right *KRbtree
	isset bool
}

func wrap(val interface{}) wrap_value {
	if v, ok := val.(KValue); ok {
		return &wrap_kvalue{v}
	}
	if v, ok := val.(int); ok {
		return &wrap_int{v}
	}
	if v, ok := val.(int8); ok {
		return &wrap_int8{v}
	}
	if v, ok := val.(int16); ok {
		return &wrap_int16{v}
	}
	if v, ok := val.(int32); ok {
		return &wrap_int32{v}
	}
	if v, ok := val.(int64); ok {
		return &wrap_int64{v}
	}
	if v, ok := val.(uint); ok {
		return &wrap_uint{v}
	}
	if v, ok := val.(uint8); ok {
		return &wrap_uint8{v}
	}
	if v, ok := val.(uint16); ok {
		return &wrap_uint16{v}
	}
	if v, ok := val.(uint32); ok {
		return &wrap_uint32{v}
	}
	if v, ok := val.(uint64); ok {
		return &wrap_uint64{v}
	}
	if v, ok := val.(float32); ok {
		return &wrap_float32{v}
	}
	if v, ok := val.(float64); ok {
		return &wrap_float64{v}
	}
	if v, ok := val.(string); ok {
		return &wrap_string{v}
	}
	panic(fmt.Sprintln("Unknown datatype ", reflect.TypeOf(val), " in ", val))
	return nil
}

func (k *KRbtree) Add(val interface{}) bool {
	return k.add(wrap(val))
}

func (k *KRbtree) add(val wrap_value) bool {
	if k.isset {
		v := k.value.FunCompare(val)
		if v == 0 {
			return false
		}
		if v > 0 {
			if k.left == nil {
				k.left = new(KRbtree)
			}
			return k.left.add(val)
		}

		if k.right == nil {
			k.right = new(KRbtree)
		}
		return k.right.add(val)
	}
	k.isset = true
	k.value = val
	return true
}

func (k *KRbtree) removeFirst() wrap_value {
	if k.left != nil {
		value := k.left.removeFirst()
		if !k.left.isset {
			k.left = nil
		}
		return value
	}
	value := k.value
	k.remove(value)
	return value
}

func (k *KRbtree) replaceWith(o *KRbtree) {
	k.value = o.value
	k.left = o.left
	k.right = o.right
}

func (k *KRbtree) Remove(val interface{}) bool {
	return k.remove(wrap(val))
}

func (k *KRbtree) remove(val wrap_value) bool {
	if !k.isset {
		return false
	}
	v := k.value.FunCompare(val)
	if v == 0 {
		if k.right != nil {
			if k.left == nil {
				k.replaceWith(k.right)
			} else {
				k.value = k.right.removeFirst()
			}
		} else if k.left != nil {
			k.replaceWith(k.left)
		} else {
			k.isset = false
		}
	}
	if v > 0 {
		if k.left != nil {
			defer func() {
				if !k.left.isset {
					k.left = nil
				}
			}()
			return k.left.remove(val)
		}
		return false
	}

	if k.right != nil {
		defer func() {
			if !k.right.isset {
				k.right = nil
			}
		}()
		return k.right.remove(val)
	}
	return false
}

func (k *KRbtree) Contains(val interface{}) bool {
	return k.contains(wrap(val))
}

func (k *KRbtree) contains(val wrap_value) bool {
	if !k.isset {
		return false
	}
	v := k.value.FunCompare(val)
	if v == 0 {
		return true
	}
	if v > 0 {
		if k.left == nil {
			return false
		}
		return k.left.contains(val)
	}
	if k.right == nil {
		return false
	}
	return k.right.contains(val)
}

func (k *KRbtree) Find(val interface{}) interface{} {
	if !k.isset {
		return nil
	}
	return k.find(wrap(val))
}

func (k *KRbtree) find(val wrap_value) interface{} {
	v := k.value.FunCompare(val)
	if v == 0 {
		return k.value
	}
	if v > 0 {
		if k.left == nil {
			return nil
		}
		return k.left.find(val)
	}

	if k.right == nil {
		return nil
	}
	return k.right.find(val)
}

func (k *KRbtree) First() interface{} {
	if !k.isset {
		return nil
	}
	if k.left != nil {
		return k.left.First()
	}
	return k.value.Value()
}

func (k *KRbtree) Last() interface{} {
	if !k.isset {
		return nil
	}
	if k.right != nil {
		return k.right.Last()
	}
	return k.value.Value()
}

func (k *KRbtree) Walk(fn func(interface{}) bool) {
	if k.isset {
		k.walk(fn)
	}
}

func (k *KRbtree) walk(fn func(interface{}) bool) bool {
	if k.left != nil {
		if !k.left.walk(fn) {
			return false
		}
	}
	if !fn(k.value.Value()) {
		return false
	}
	if k.right != nil {
		if !k.right.walk(fn) {
			return false
		}
	}
	return true
}

func (k *KRbtree) Length() int {
	count := 0
	k.Walk(func(v interface{}) bool {
		count++
		return true
	})
	return count
}

func (k *KRbtree) ToSlice() []interface{} {
	if !k.isset {
		return nil
	}

	slice := make([]interface{}, k.Length())
	i := 0
	k.Walk(func(v interface{}) bool {
		slice[i] = v
		i++
		return true
	})
	return slice
}

func (k *KRbtree) Balance() *KRbtree {
	slice := k.ToSlice()

	balanced := new(KRbtree)

	balanced._balance(slice)

	return balanced
}

func (k *KRbtree) _balance(slice []interface{}) {
	if len(slice) == 0 {
		return
	}
	if len(slice) == 1 {
		k.Add(slice[0])
		return
	}
	mid := len(slice) / 2
	k.Add(slice[mid])
	k._balance(slice[:mid])
	k._balance(slice[mid:])
}

type wrap_int struct {
	val int
}

func (k *wrap_int) FunCompare(o wrap_value) int64 {
	if v, ok := o.(*wrap_int); ok {
		return int64(k.val - v.val)
	}
	return 0
}

func (k *wrap_int) Value() interface{} {
	return k.val
}

type wrap_uint struct {
	val uint
}

func (k *wrap_uint) FunCompare(o wrap_value) int64 {
	if v, ok := o.(*wrap_uint); ok {
		return int64(k.val - v.val)
	}
	return 0
}

func (k *wrap_uint) Value() interface{} {
	return k.val
}

type wrap_int8 struct {
	val int8
}

func (k *wrap_int8) FunCompare(o wrap_value) int64 {
	if v, ok := o.(*wrap_int8); ok {
		return int64(k.val - v.val)
	}
	return 0
}

func (k *wrap_int8) Value() interface{} {
	return k.val
}

type wrap_uint8 struct {
	val uint8
}

func (k *wrap_uint8) FunCompare(o wrap_value) int64 {
	if v, ok := o.(*wrap_uint8); ok {
		return int64(k.val - v.val)
	}
	return 0
}

func (k *wrap_uint8) Value() interface{} {
	return k.val
}

type wrap_int16 struct {
	val int16
}

func (k *wrap_int16) FunCompare(o wrap_value) int64 {
	if v, ok := o.(*wrap_int16); ok {
		return int64(k.val - v.val)
	}
	return 0
}

func (k *wrap_int16) Value() interface{} {
	return k.val
}

type wrap_uint16 struct {
	val uint16
}

func (k *wrap_uint16) FunCompare(o wrap_value) int64 {
	if v, ok := o.(*wrap_uint16); ok {
		return int64(k.val - v.val)
	}
	return 0
}

func (k *wrap_uint16) Value() interface{} {
	return k.val
}

type wrap_int32 struct {
	val int32
}

func (k *wrap_int32) FunCompare(o wrap_value) int64 {
	if v, ok := o.(*wrap_int32); ok {
		return int64(k.val - v.val)
	}
	return 0
}

func (k *wrap_int32) Value() interface{} {
	return k.val
}

type wrap_uint32 struct {
	val uint32
}

func (k *wrap_uint32) FunCompare(o wrap_value) int64 {
	if v, ok := o.(*wrap_uint32); ok {
		return int64(k.val - v.val)
	}
	return 0
}

func (k *wrap_uint32) Value() interface{} {
	return k.val
}

type wrap_int64 struct {
	val int64
}

func (k *wrap_int64) FunCompare(o wrap_value) int64 {
	if v, ok := o.(*wrap_int64); ok {
		return int64(k.val - v.val)
	}
	return 0
}

func (k *wrap_int64) Value() interface{} {
	return k.val
}

type wrap_uint64 struct {
	val uint64
}

func (k *wrap_uint64) FunCompare(o wrap_value) int64 {
	if v, ok := o.(*wrap_uint64); ok {
		return int64(k.val - v.val)
	}
	return 0
}

func (k *wrap_uint64) Value() interface{} {
	return k.val
}

type wrap_float32 struct {
	val float32
}

func (k *wrap_float32) FunCompare(o wrap_value) int64 {
	if v, ok := o.(*wrap_float32); ok {
		if k.val == v.val {
			return 0
		}
		if k.val < v.val {
			return -1
		}
		return 1
	}
	return 0
}

func (k *wrap_float32) Value() interface{} {
	return k.val
}

type wrap_float64 struct {
	val float64
}

func (k *wrap_float64) FunCompare(o wrap_value) int64 {
	if v, ok := o.(*wrap_float64); ok {
		if k.val == v.val {
			return 0
		}
		if k.val < v.val {
			return -1
		}
		return 1
	}
	return 0
}

func (k *wrap_float64) Value() interface{} {
	return k.val
}

type wrap_string struct {
	val string
}

func (k *wrap_string) FunCompare(o wrap_value) int64 {
	if v, ok := o.(*wrap_string); ok {
		if k.val == v.val {
			return 0
		}
		if k.val < v.val {
			return -1
		}
		return 1
	}
	return 0
}

func (k *wrap_string) Value() interface{} {
	return k.val
}

type wrap_kvalue struct {
	val KValue
}

func (k *wrap_kvalue) FunCompare(o wrap_value) int64 {
	if v, ok := o.(*wrap_kvalue); ok {
		return k.val.FunCompare(v.val)
	}
	return 0
}

func (k *wrap_kvalue) Value() interface{} {
	return k.val
}
