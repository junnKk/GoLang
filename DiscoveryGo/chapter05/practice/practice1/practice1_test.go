package subTask

import "fmt"

func ExampleMarkDone() {
	t := Task{
		Title:    "Laundry",
		Status:   TODO,
		Deadline: nil,
		Priority: 2,
		SubTasks: []Task{{
			Title:    "Laundry",
			Status:   TODO,
			Deadline: nil,
			Priority: 2,
			SubTasks: []Task{
				{"Put", DONE, nil, 2, nil},
				{"Detergent", TODO, nil, 2, nil},
			},
		}, {
			Title:    "Dry",
			Status:   TODO,
			Deadline: nil,
			Priority: 2,
			SubTasks: nil,
		}, {
			Title:    "Fold",
			Status:   TODO,
			Deadline: nil,
			Priority: 2,
			SubTasks: nil,
		}},
	}
	t.MarkDone()
	fmt.Println(IncludeSubTasks(t))
	// Output:
	// [V] Laundry <nil>
	//  [V] Laundry <nil>
	//   [V] Put <nil>
	//   [V] Detergent <nil>
	//  [V] Dry <nil>
	//  [V] Fold <nil>
}
