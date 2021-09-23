package structs

import "time"

var task = struct {
	title string
	done  bool
	due   *time.Time
}{"laudry", false, nil}

type Task struct {
	title string
	done  bool
	due   *time.Time
}

type BetterTask struct {
	title  string
	status status
	due    *time.Time
}

type status int

const (
	UNKNOWN status = iota
	TODO
	DONE
)

type ByteSize int

const (
	KB ByteSize = 1 << (10 * (1 + iota))
	MB
	GB
	TB
	PB
)

type Deadline struct {
	time.Time
}

func NewDeadline(t time.Time) *Deadline {
	return &Deadline{t}
}

// OverDue returns true if the deadline is before the current time.
func (d *Deadline) OverDue() bool {
	return d != nil && time.Time((*d).Time).Before(time.Now())
}

// OverDue returns true if the deadline is before the current time.
func (t *DueTask) OverDue() bool {
	return t.Deadline.OverDue()
}

type DueTask struct {
	Title    string `json:"title"`
	Status   status `json:"-"`
	Deadline *Deadline
}
