package aop

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func TestNewRecorder(t *testing.T) {
	tests := []struct {
		name string
		want *Recorder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRecorder(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRecorder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRecorder_RegisterHandler(t *testing.T) {
	type args struct {
		f interface{}
	}
	tests := []struct {
		name string
		r    *Recorder
		args args
		want *Handler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.RegisterHandlerFunc(tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Recorder.RegisterHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRecorder_Start(t *testing.T) {
	tests := []struct {
		name string
		r    *Recorder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.r.Start()
		})
	}
}

func TestRecorder_CollectMsg(t *testing.T) {
	r := GetRecorder(myHandler)
	ret := make([][]interface{}, 4)
	for i := range ret {
		ret[i] = make([]interface{}, 2)
	}
	ex := make([]interface{}, 4)
	tests := []struct {
		name string
		r    *Recorder
		msg  *Message
	}{
		// TODO: Add test cases.
		{"1", r, &Message{[]interface{}{"m1", "m11"}, []interface{}{nil, nil}, &ex[0]}},
		{"2", r, &Message{[]interface{}{"m2", "m22"}, []interface{}{&ret[1][0], nil}, &ex[1]}},
		{"3", r, &Message{[]interface{}{"m3", "m33"}, []interface{}{&ret[2][0], &ret[2][1]}, &ex[2]}},
		{"4", r, &Message{[]interface{}{"m4", "m44"}, []interface{}{&ret[3][0], &ret[3][1]}, &ex[3]}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.r.CollectMsg(tt.msg)
		})
	}
}
func TestRecorder_CollectMsg2(t *testing.T) {
	r := GetRecorder(myHandler)
	ret := make([][]interface{}, 4)
	for i := range ret {
		ret[i] = make([]interface{}, 2)
	}
	ex := make([]interface{}, 4)
	m1 := new(Message)
	m2 := new(Message)
	m3 := new(Message)
	m4 := new(Message)
	// same as m1.Args("m1", "m11").Receivers(nil, nil).Exception(&ex[0])
	m1.Args("m1", "m11").Exception(&ex[0])
	m2.Args("m2", "m22").Receivers(&ret[1][0], nil).Exception(&ex[1])
	m3.Args("m3", "m33").Receivers(&ret[2][0], &ret[2][1]).Exception(&ex[2])
	m4.Args("m4", "m44").Receivers(&ret[3][0], &ret[3][1]).Exception(&ex[3])
	tests := []struct {
		name string
		r    *Recorder
		msg  *Message
	}{
		// TODO: Add test cases.
		{"1", r, m1},
		{"2", r, m2},
		{"3", r, m3},
		{"4", r, m4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.r.CollectMsg(tt.msg)
		})
	}
}

func TestRecorder_CollectMsgArgs(t *testing.T) {
	r := GetRecorder(myHandler)
	tests := []struct {
		name string
		r    *Recorder
		msg  []interface{}
	}{
		// TODO: Add test cases.
		{"1", r, []interface{}{"m1", "m11"}},
		{"2", r, []interface{}{"m2", "m22"}},
		{"3", r, []interface{}{"m3", "m33"}},
		{"4", r, []interface{}{"m4", "m44"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.r.CollectMsgArgs(tt.msg...)
		})
	}
}

func TestGetRecorder(t *testing.T) {
	type args struct {
		handlerFunc interface{}
	}
	tests := []struct {
		name string
		args args
		want *Recorder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetRecorder(tt.args.handlerFunc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRecorder() = %v, want %v", got, tt.want)
			}
		})
	}
}

// recorder handler, method on stuct is also ok
func myHandler(p1, p2 interface{}) (string, error) {
	s := fmt.Sprintf("%s %s", p1, p2)
	// panic(s)
	return "ok: " + s, errors.New("errorhere")
}
