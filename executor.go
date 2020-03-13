package aop

import "sync"

import "time"

var (
	DefaultTaskTimeOut = time.Hour * 12
)

// Executor struct
type Executor struct {
	wg *sync.WaitGroup
	// slotPool maintain a fixed size slot pool to execute tasks
	slotPool chan struct{}
}

// Task executor runs, task parallesim is capacity of recorderChan
type Task struct {
	Handler
	TimeOut time.Duration
}

// NewTask new a task with default task timeout
func NewTask() *Task {
	t := new(Task)
	t.TimeOut = DefaultTaskTimeOut
	return t
}

// NewTaskWithTimeOut new a task
func NewTaskWithTimeOut(d time.Duration) *Task {
	t := new(Task)
	t.TimeOut = d
	return t
}

// NewExecutor new a executor
func NewExecutor(cap int) *Executor {
	e := new(Executor)
	e.wg = new(sync.WaitGroup)
	e.slotPool = make(chan struct{}, cap)
	return e
}

// Func sets the args of Task
func (t *Task) Func(f interface{}) *Task {
	t.f = f
	return t
}

// Args sets the args of Task
func (t *Task) Args(args ...interface{}) *Task {
	t.args = nil
	t.args = append(t.args, args...)
	return t
}

// Receivers sets the receivers of return values
func (t *Task) Receivers(receivers ...interface{}) *Task {
	t.receivers = nil
	t.receivers = append(t.receivers, receivers...)
	return t
}

// Exception sets the exception of Message
func (t *Task) Exception(ex interface{}) *Task {
	t.exception = ex
	return t
}

// Run concret task when acquire slot
// isFull(executor.slotPool)? blocking : non-blocking
func (e *Executor) Run(t *Task) {
	e.slotPool <- struct{}{}
	e.wg.Add(1)
	go func(t *Task) {
		defer func() {
			<-e.slotPool
			e.wg.Done()
			err := recover()
			if err != nil {
				panic(err)
			}
		}()
		t.DoWithTimeOut(t.TimeOut)
	}(t)
}

// Wait all goroutine done
func (e *Executor) Wait() {
	e.wg.Wait()
}
