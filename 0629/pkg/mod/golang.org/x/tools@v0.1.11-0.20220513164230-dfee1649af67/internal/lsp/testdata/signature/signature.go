// Package signature has tests for signature help.
package signature

import (
	"bytes"
	"encoding/json"
	"math/big"
)

func Foo(a string, b int) (c bool) {
	return
}

func Bar(float64, ...byte) {
}

type myStruct struct{}

func (*myStruct) foo(e *json.Decoder) (*big.Int, error) {
	return nil, nil
}

type MyType struct{}

type MyFunc func(foo int) string

type Alias = int
type OtherAlias = int
type StringAlias = string

func AliasSlice(a []*Alias) (b Alias)                                 { return 0 }
func AliasMap(a map[*Alias]StringAlias) (b, c map[*Alias]StringAlias) { return nil, nil }
func OtherAliasMap(a, b map[Alias]OtherAlias) map[Alias]OtherAlias    { return nil }

func Qux() {
	Foo("foo", 123) //@signature("(", "Foo(a string, b int) (c bool)", 0)
	Foo("foo", 123) //@signature("123", "Foo(a string, b int) (c bool)", 1)
	Foo("foo", 123) //@signature(",", "Foo(a string, b int) (c bool)", 0)
	Foo("foo", 123) //@signature(" 1", "Foo(a string, b int) (c bool)", 1)
	Foo("foo", 123) //@signature(")", "Foo(a string, b int) (c bool)", 1)

	Bar(13.37, 0x13)       //@signature("13.37", "Bar(float64, ...byte)", 0)
	Bar(13.37, 0x37)       //@signature("0x37", "Bar(float64, ...byte)", 1)
	Bar(13.37, 1, 2, 3, 4) //@signature("4", "Bar(float64, ...byte)", 1)

	fn := func(hi, there string) func(i int) rune {
		return func(int) rune { return 0 }
	}

	fn("hi", "there")    //@signature("hi", "", 0)
	fn("hi", "there")    //@signature(",", "fn(hi string, there string) func(i int) rune", 0)
	fn("hi", "there")(1) //@signature("1", "func(i int) rune", 0)

	fnPtr := &fn
	(*fnPtr)("hi", "there") //@signature(",", "func(hi string, there string) func(i int) rune", 0)

	var fnIntf interface{} = Foo
	fnIntf.(func(string, int) bool)("hi", 123) //@signature("123", "func(string, int) bool", 1)

	(&bytes.Buffer{}).Next(2) //@signature("2", "Next(n int) []byte", 0)

	myFunc := MyFunc(func(n int) string { return "" })
	myFunc(123) //@signature("123", "myFunc(foo int) string", 0)

	var ms myStruct
	ms.foo(nil) //@signature("nil", "foo(e *json.Decoder) (*big.Int, error)", 0)

	_ = make([]int, 1, 2) //@signature("2", "make(t Type, size ...int) Type", 1)

	Foo(myFunc(123), 456) //@signature("myFunc", "Foo(a string, b int) (c bool)", 0)
	Foo(myFunc(123), 456) //@signature("123", "myFunc(foo int) string", 0)

	panic("oops!")            //@signature(")", "panic(v interface{})", 0)
	println("hello", "world") //@signature(",", "println(args ...Type)", 0)

	Hello(func() {
		//@signature("//", "", 0)
	})

	AliasSlice()    //@signature(")", "AliasSlice(a []*Alias) (b Alias)", 0)
	AliasMap()      //@signature(")", "AliasMap(a map[*Alias]StringAlias) (b map[*Alias]StringAlias, c map[*Alias]StringAlias)", 0)
	OtherAliasMap() //@signature(")", "OtherAliasMap(a map[Alias]OtherAlias, b map[Alias]OtherAlias) map[Alias]OtherAlias", 0)
}

func Hello(func()) {}
