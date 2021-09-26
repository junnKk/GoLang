package multiset

import "fmt"

func ExampleMultiSet() {
	m := MultiSet{}
	m.Insert("a")
	m.Insert("a")
	m.Insert("b")
	m.Erase("a")
	fmt.Println(m)
	fmt.Println(m.Count("a"))
	// Output:
	// { a b }
	// 1
}
