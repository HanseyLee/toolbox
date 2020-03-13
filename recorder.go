package aop

import "time"

// aop, aspect oriented programming
// non-blocking, is sutiable for authuentication, logging, security, check point recording...

var (
	// RecorderChanCap default capacity of RecorderChan channel
	RecorderChanCap       = 512
	DefaultHandlerTimeOut = time.Hour * 12
)

// Recorder receive check point status and record
type Recorder struct {
	// RecorderChan store msgs, such as check point status
	recorderChan chan *Message
	// Handler handle msg
	handler *Handler
	TimeOut time.Duration
}

// Message in recorderChan
type Message struct {
	args      []interface{}
	receivers []interface{}
	exception interface{}
}

// Args sets the args of Message
func (m *Message) Args(args ...interface{}) *Message {
	m.args = nil
	m.args = append(m.args, args...)
	return m
}

// Receivers sets the receivers of return values
func (m *Message) Receivers(receivers ...interface{}) *Message {
	m.receivers = nil
	m.receivers = append(m.receivers, receivers...)
	return m
}

// Exception sets the exception of Message
func (m *Message) Exception(ex interface{}) *Message {
	m.exception = ex
	return m
}

// NewRecorder new a recorder with default capacity
func NewRecorder() *Recorder {
	r := new(Recorder)
	r.recorderChan = make(chan *Message, RecorderChanCap)
	r.TimeOut = DefaultHandlerTimeOut
	return r
}

// RegisterHandlerFunc regitster msg handler
func (r *Recorder) RegisterHandlerFunc(f interface{}) *Handler {
	h := NewHandler()
	h.SetFunc(f)
	r.handler = h
	return h
}

// Start a gorountine, blocking receiving and handling one by one
func (r *Recorder) Start() {
	go func() {
		for {
			msg := <-r.recorderChan
			r.handler.SetArgs(msg.args...).
				SetReceivers(msg.receivers...).
				SetException(msg.exception)
			r.handler.DoWithTimeOut(r.TimeOut)
		}
	}()
}

// CollectMsg collect msgs published by other funcs
func (r *Recorder) CollectMsg(msg *Message) {
	r.recorderChan <- msg
}

// CollectMsgArgs collect msg args published by other funcs
func (r *Recorder) CollectMsgArgs(args ...interface{}) {
	msg := new(Message)
	msg.args = append(msg.args, args...)
	r.recorderChan <- msg
}

// GetRecorder new recorder and start message channel
func GetRecorder(handlerFunc interface{}) *Recorder {
	r := NewRecorder()
	r.RegisterHandlerFunc(handlerFunc)
	r.Start()
	return r
}
