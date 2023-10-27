package scheduling

import (
	"sort"
	"time"
)

const DeadFormat = "02/01/2006 15:04"

type Task struct {
	Title string
	Description string
	Deadline time.Time
	Duration time.Duration
}

type By func(t1, t2 *Task) bool
type taskSorter struct{
	tasks []Task
	by func(t1, t2 *Task) bool
}

func (by By) Sort (tasks []Task) {
	ps := &taskSorter {
		tasks: tasks,
		by: by,
	}
	sort.Sort(ps)
}

func (t *taskSorter) Len() int {
	return len(t.tasks)
}

func (t *taskSorter) Swap(i, j int) {
	t.tasks[i], t.tasks[j] = t.tasks[j], t.tasks[i]	
}

func (t *taskSorter) Less (i, j int) bool {
	return t.by(&t.tasks[i], &t.tasks[j])
}

func (by By) sortTask(ts []Task) {
	deadLine := func(t1, t2 *Task) bool {
		return t1.Deadline.Compare(t2.Deadline) < 0
	}
	By(deadLine).Sort(ts)
}

type ScheTask struct {
	T Task
	Start time.Time
	End time.Time;
}

func Scheduling(ts []Task) []ScheTask{
	deadLine := func(t1, t2 *Task) bool {
		return t1.Deadline.Compare(t2.Deadline) < 0
	}	
	By(deadLine).Sort(ts)
	sche := make([]ScheTask, 0, len(ts))
	tm := time.Now()
	for _, t := range ts {
		st := ScheTask{}
		st.T = t
		st.Start = tm
		st.End = tm.Add(t.Duration)
		tm = tm.Add(t.Duration)
		sche = append(sche, st)	
	}
	return sche
}
