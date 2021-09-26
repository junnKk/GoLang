package task

import (
	"fmt"
)

func ExampleInMemoryAccessor() {
	m := NewInMemoryAccessor()
	t1 := Task{
		Title:    "Laundry",
		Status:   TODO,
		Deadline: nil,
		Priority: 2,
	}
	id1, err1 := m.Post(t1)
	fmt.Println(id1, err1)
	fmt.Println(m.Get(id1))

	t2 := Task{
		Title:    "Detergent",
		Status:   TODO,
		Deadline: nil,
		Priority: 1,
	}
	m.Put(id1, t2)
	fmt.Println(m.Get(id1))

	fmt.Println(m.Get(ID(fmt.Sprint(100))))

	m.Delete(id1)
	fmt.Println(m.Get(id1))
	// Output:
	// 1 <nil>
	// [ ] Laundry <nil> <nil>
	// [ ] Detergent <nil> <nil>
	// [ ]  <nil> task does not exist
	// [ ]  <nil> task does not exist
}
