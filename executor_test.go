package aop

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func TestNewExecutor(t *testing.T) {
	type args struct {
		cap int
	}
	tests := []struct {
		name string
		args args
		want *Executor
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewExecutor(tt.args.cap); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewExecutor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExecutor_Run(t *testing.T) {
	e := NewExecutor(3)
	ret := make([][]interface{}, 10)
	for i := range ret {
		ret[i] = make([]interface{}, 2)
	}
	ex := make([]interface{}, 10)
	t0 := NewTask()
	t1 := NewTask()
	t2 := NewTask()
	t3 := NewTask()
	t4 := NewTask()
	t5 := NewTask()
	t6 := NewTask()
	t7 := NewTask()
	t8 := NewTask()
	t9 := NewTask()
	t0.Func(myJob).Args("t0").Exception(&ex[0])
	t1.Func(myJob).Args("t1").Receivers(&ret[1][0], nil).Exception(&ex[1])
	t2.Func(myJob).Args("t2").Receivers(nil, &ret[2][1]).Exception(&ex[2])
	t3.Func(myJob).Args("t3").Receivers(&ret[3][0], &ret[3][1]).Exception(&ex[3])
	t4.Func(myJob).Args("t4").Receivers(&ret[4][0], &ret[4][1]).Exception(&ex[4])
	t5.Func(myJob).Args("t5").Receivers(&ret[5][0], &ret[5][1]).Exception(&ex[5])
	t6.Func(myJob).Args("t6").Receivers(&ret[6][0], &ret[6][1]).Exception(&ex[6])
	t7.Func(myJob).Args("t7").Receivers(&ret[7][0], &ret[7][1]).Exception(&ex[7])
	t8.Func(myJob).Args("t8").Receivers(&ret[8][0], &ret[8][1]).Exception(&ex[8])
	t9.Func(myJob).Args("t9").Receivers(&ret[9][0], &ret[9][1]).Exception(&ex[9])

	tests := []struct {
		name string
		e    *Executor
		t    *Task
	}{
		// TODO: Add test cases.
		{"0", e, t0},
		{"1", e, t1},
		{"2", e, t2},
		{"3", e, t3},
		{"4", e, t4},
		{"5", e, t5},
		{"6", e, t6},
		{"7", e, t7},
		{"8", e, t8},
		{"9", e, t9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.e.Run(tt.t)
		})
	}
	e.Wait()
}

// use recorder as a fixed-size executor pool
func myJob(p1 interface{}) (string, error) {
	s := fmt.Sprintf("%s", p1)
	// time.Sleep(time.Second * 1)
	// panic(s)
	return "ok: " + s, errors.New("error-here")
}
