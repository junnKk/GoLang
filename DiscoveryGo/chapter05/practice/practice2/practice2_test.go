package sorts

import (
	"reflect"
	"sort"
	"strings"
	"testing"

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
