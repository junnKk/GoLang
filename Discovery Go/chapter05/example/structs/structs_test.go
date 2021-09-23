package structs

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/glowingedge/discoveryGo/seq"
)

func TestHowToWriteTC(t *testing.T) {
	var cases = []struct {
		in, want int
	}{
		{0, 0},
		{5, 5},
		{6, 8},
	}

	for i, c := range cases {
		got := seq.Fib(c.in)
		if got != c.want {
			t.Errorf("Case %d. Fib(%d) == %d, want %d", i, c.in, got, c.want)
		}
	}
}

func ExampleDeadline_OverDue() {
	d1 := NewDeadline(time.Now().Add(-4 * time.Hour))
	d2 := NewDeadline(time.Now().Add(4 * time.Hour))
	fmt.Println(d1.OverDue())
	fmt.Println(d2.OverDue())
	// Output:
	// true
	// false
}

func Example_taskTestAll() {
	d1 := NewDeadline(time.Now().Add(-4 * time.Hour))
	d2 := NewDeadline(time.Now().Add(4 * time.Hour))
	t1 := DueTask{"4h ago", TODO, d1}
	t2 := DueTask{"4h later", TODO, d2}
	t3 := DueTask{"no due", TODO, nil}
	fmt.Println(t1.OverDue())
	fmt.Println(t2.OverDue())
	fmt.Println(t3.OverDue())
	// Output:
	// true
	// false
	// false
}

func Example_marshalJSON() {
	t := DueTask{
		"Laundry",
		DONE,
		NewDeadline(time.Date(2015, time.August, 16, 15, 43, 0, 0, time.UTC)),
	}
	b, err := json.Marshal(t)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(b))

	t2 := &DueTask{
		"Laundry",
		DONE,
		NewDeadline(time.Date(2015, time.August, 16, 15, 43, 0, 0, time.UTC)),
	}
	b2, _ := json.Marshal(struct {
		*DueTask
		Title      string `json:"title,omitempty"`
		Additional string
	}{DueTask: t2, Additional: "abcd"})
	fmt.Println(string(b2))

	// Output:
	// {"Title":"Laundry","Deadline":"2015-08-16T15:43:00Z"}
	// {"Deadline":"2015-08-16T15:43:00Z","Additional":"abcd"}
}

func Example_unmarshalJSON() {
	b := []byte(`{"Title":"Laundry","Status":2,"Deadline":"2015-08-16T15:43:00Z"}`)
	t := DueTask{}
	err := json.Unmarshal(b, &t)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(t.Title)
	fmt.Println(t.Status)
	fmt.Println(t.Deadline.UTC())
	// Output:
	// Laundry
	// 0
	// 2015-08-16 15:43:00 +0000 UTC
}
