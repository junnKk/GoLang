package string

import (
	"fmt"
	"strconv"
)

func Example_strCat() {
	s := "abc"
	ps := &s
	s += "def"
	fmt.Println(s)
	fmt.Println(*ps)

	// Output:
	// abcdef
	// abcdef
}

func Example_strConv() {
	var i int
	var k int64
	var k2 int64
	var f float64
	var s string
	var s2 string
	var err error

	i, err = strconv.Atoi("350")                  // i === 350
	k, err = strconv.ParseInt("cc7fdd", 16, 32)   // k == 13402077
	k2, err = strconv.ParseInt("0xcc7fdd", 0, 32) // k2 == 13402077
	f, err = strconv.ParseFloat("3.14", 64)       // f == 3.14
	s = strconv.Itoa(340)                         // s == "340"
	s2 = strconv.FormatInt(13402077, 16)          // s2 == "cc8fdd"

	fmt.Println(i)
	fmt.Println(k)
	fmt.Println(k2)
	fmt.Println(f)
	fmt.Println(s)
	fmt.Println(s2)
	fmt.Println(err == nil)

	// Output:
	// 350
	// 13402077
	// 13402077
	// 3.14
	// 340
	// cc7fdd
	// true
}
