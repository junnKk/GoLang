package method

import "fmt"

type VertexID int

func (id VertexID) String() string {
	return fmt.Sprintf("VertexID(%d)", id)
}

func (id VertexID) add(n int) {
	id += VertexID(n)
	fmt.Println("dd ", id)
}
