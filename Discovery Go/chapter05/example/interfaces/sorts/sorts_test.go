package sorts

import (
	"container/heap"
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"testing"

	"github.com/glowingedge/discoveryGo/chapter5/interfaces/subTask"
)

type CaseInsensitive []string

func (c CaseInsensitive) Len() int {
	return len(c)
}

func (c CaseInsensitive) Less(i, j int) bool {
	return strings.ToLower(c[i]) < strings.ToLower(c[j]) ||
		(strings.ToLower(c[i]) == strings.ToLower(c[j])) && c[i] < c[j]
}

func (c CaseInsensitive) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c *CaseInsensitive) Push(x interface{}) {
	*c = append(*c, x.(string))
}

func (c *CaseInsensitive) Pop() interface{} {
	len := c.Len()
	last := (*c)[len-1]
	*c = (*c)[:len-1]
	return last
}

func Join(sep string, a ...interface{}) string {
	if len(a) == 0 {
		return ""
	}
	t := make([]string, len(a))
	for i := range a {
		switch x := a[i].(type) {
		case string:
			t[i] = x
		case int:
			t[i] = strconv.Itoa(x)
		case fmt.Stringer:
			t[i] = x.String()
		}
	}
	return strings.Join(t, sep)
}

func ExampleCaseInsensitive_sort() {
	apple := CaseInsensitive([]string{
		"iPhone", "iPad", "MacBook", "AppStore",
	})
	sort.Sort(apple)
	fmt.Println(apple)
	// Output:
	// [AppStore iPad iPhone MacBook]
}

func ExampleCaseInsensitive_heap() {
	apple := CaseInsensitive([]string{
		"iPhone", "iPad", "MacBook", "AppStore",
	})
	heap.Init(&apple)
	for apple.Len() > 0 {
		popped := heap.Pop(&apple)
		s := popped.(string)
		fmt.Println(s)
	}
	// Output:
	// AppStore
	// iPad
	// iPhone
	// MacBook
}

func ExampleJoin() {
	t := subTask.Task{
		Title:    "abcd",
		Status:   subTask.DONE,
		Deadline: nil,
	}
	fmt.Println(Join("---", 1, "TWO", 3, t))
	// Output:
	// 1---TWO---3---[V] abcd <nil>
}

func TestCaseInsensitive_sort_tableTest(t *testing.T) {
	var cases = []struct {
		in, want CaseInsensitive
	}{
		{CaseInsensitive([]string{"iPhone", "iPad", "MacBook", "AppStore"}), CaseInsensitive([]string{"AppStore", "iPad", "iPhone", "MacBook"})},
	}

	for i, c := range cases {
		sort.Sort(c.in)
		if !reflect.DeepEqual(c.in, c.want) {
			t.Errorf("Case %d. Sort(%s) == %s, want %s", i, c.in, c.in, c.want)
		}
	}
}
