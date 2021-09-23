package method

import "fmt"

func ExampleVertexID_method() {
	i := VertexID(100)
	fmt.Println(i)
	fmt.Println(i.String())
	i.add(10)
	fmt.Println(i)
	// Output:
	// VertexID(100)
	// VertexID(100)
}
