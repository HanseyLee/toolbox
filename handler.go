package aop
// refer to https://github.com/buptmiao/parallel/handler.go

import (
	"errors"
	"reflect"
	"time"
)

var (
	ErrArgNotFunction              = errors.New("argument type not function")
	ErrInArgLenNotMatch            = errors.New("input arguments length not match")
	ErrOutArgLenNotMatch           = errors.New("output arguments length not match")
	ErrRecvArgTypeNeitherPtrNorNil = errors.New("receiver argument type is neither pointer nor nil")
)

// Handler instance
type Handler struct {
	// The type of f must be function
	f    interface{}
	args []interface{}
	// The type of every receiver must be ptr, to receive the return value of f call. skip nil
	receivers []interface{}
	// receive panic error when it can not be processed.
	exception interface{}
}

// NewHandler create a new Handler
func NewHandler() *Handler {
	res := new(Handler)
	return res
}

// SetFunc sets the function of Handler
func (h *Handler) SetFunc(f interface{}) *Handler {
	h.f = f
	return h
}

// SetArgs sets the args of Handler
func (h *Handler) SetArgs(args ...interface{}) *Handler {
	h.args = nil
	h.args = append(h.args, args...)
	return h
}

// SetReceivers sets the receivers of return values
func (h *Handler) SetReceivers(receivers ...interface{}) *Handler {
	h.receivers = nil
	h.receivers = append(h.receivers, receivers...)
	return h
}

// SetException sets the exception of Handler
func (h *Handler) SetException(ex interface{}) *Handler {
	h.exception = ex
	return h
}

// Do call the function and return values if exists
func (h *Handler) Do() {
	defer func() {
		err := recover()
		if err != nil {
			if err == ErrArgNotFunction || err == ErrInArgLenNotMatch ||
				err == ErrOutArgLenNotMatch || err == ErrRecvArgTypeNeitherPtrNorNil {
				panic(err)
			}
			v := reflect.ValueOf(h.exception)
			if v.IsValid() {
				v.Elem().Set(reflect.ValueOf(err))
			}
		}
	}()
	f := reflect.ValueOf(h.f)
	typ := f.Type()
	//check if f is a function
	if typ.Kind() != reflect.Func {
		panic(ErrArgNotFunction)
	}
	//check input length, only check '>' is to allow varargs.
	if typ.NumIn() > len(h.args) {
		panic(ErrInArgLenNotMatch)
	}
	//check output
	if h.receivers != nil {
		//check output length
		if typ.NumOut() != len(h.receivers) {
			panic(ErrOutArgLenNotMatch)
		}
		//check if output args is ptr
		for _, v := range h.receivers {
			t := reflect.ValueOf(v)
			if t.IsValid() {
				if t.Type().Kind() != reflect.Ptr {
					panic(ErrRecvArgTypeNeitherPtrNorNil)
				}
			}
		}
	}

	inputs := make([]reflect.Value, len(h.args))
	for i := 0; i < len(h.args); i++ {
		if h.args[i] == nil {
			inputs[i] = reflect.Zero(f.Type().In(i))
		} else {
			inputs[i] = reflect.ValueOf(h.args[i])
		}
	}
	out := f.Call(inputs)
	if h.receivers != nil {
		for i := 0; i < len(h.receivers); i++ {
			v := reflect.ValueOf(h.receivers[i])
			if v.IsValid() {
				v.Elem().Set(out[i])
			}
		}
	}
}

// DoWithTimeOut return when time out
func (h *Handler) DoWithTimeOut(d time.Duration) {
	success := make(chan struct{}, 1)
	go func() {
		h.Do()
		success <- struct{}{}
	}()
	select {
	case <-success:
	case <-time.After(d):
	}
}
