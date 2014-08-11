package kutil

import (
	"fmt"
	"testing"
)

func TestToBool(t *testing.T) {
	b, ok := ToBool("F")
	fmt.Println(b, ok)
	b, ok = ToBool("AF")
	fmt.Println(b, ok)

	b, ok = ToBool("T")
	fmt.Println(b, ok)

	b, ok = ToBool("1")
	fmt.Println(b, ok)

	b, ok = ToBool("0")
	fmt.Println(b, ok)

	b, ok = ToBool(1)
	fmt.Println(b, ok)

	b, ok = ToBool(0)
	fmt.Println(b, ok)

	b, ok = ToBool(0.0)
	fmt.Println(b, ok)

	b, ok = ToBool(9.3)
	fmt.Println(b, ok)

}

func TestToInt8(t *testing.T) {
	b, ok := ToInt8("F")
	fmt.Println(b, ok)
	b, ok = ToInt8("AF")
	fmt.Println(b, ok)

	b, ok = ToInt8("T")
	fmt.Println(b, ok)

	b, ok = ToInt8("1")
	fmt.Println(b, ok)

	b, ok = ToInt8("0")
	fmt.Println(b, ok)

	b, ok = ToInt8(1)
	fmt.Println(b, ok)

	b, ok = ToInt8(0)
	fmt.Println(b, ok)

	b, ok = ToInt8(0.001)
	fmt.Println(b, ok)

	b, ok = ToInt8(9.3)
	fmt.Println(b, ok)

}
