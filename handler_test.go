package aop

import (
	"reflect"
	"testing"
)

func TestNewHandler(t *testing.T) {
	tests := []struct {
		name string
		want *Handler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHandler(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHandler_SetFunc(t *testing.T) {
	type args struct {
		f interface{}
	}
	tests := []struct {
		name string
		h    *Handler
		args args
		want *Handler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.SetFunc(tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Handler.SetFunc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHandler_SetArgs(t *testing.T) {
	type args struct {
		args []interface{}
	}
	tests := []struct {
		name string
		h    *Handler
		args args
		want *Handler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.SetArgs(tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Handler.SetArgs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHandler_SetReceivers(t *testing.T) {
	type args struct {
		receivers []interface{}
	}
	tests := []struct {
		name string
		h    *Handler
		args args
		want *Handler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.SetReceivers(tt.args.receivers...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Handler.SetReceivers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHandler_SetException(t *testing.T) {
	type args struct {
		ex interface{}
	}
	tests := []struct {
		name string
		h    *Handler
		args args
		want *Handler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.SetException(tt.args.ex); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Handler.SetException() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHandler_Do(t *testing.T) {
	tests := []struct {
		name string
		h    *Handler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.Do()
		})
	}
}
