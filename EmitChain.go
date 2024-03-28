package main

import (
	"errors"
	"time"
)

type Sent struct {
	emit func(inf interface{}) (interface{}, error)
}

type EmitChain struct {
	timeout   time.Duration `default:"time.Duration(5) * time.Second"` //超时时间
	retry     int           //重试次数
	send      *Sent
	headerInf HeaderInf
	bodyInf   BodyInf
}

func (e *EmitChain) Timeout(t int) *EmitChain {
	println("entry Timeout")
	e.timeout = time.Duration(t) * time.Second
	return e
}

func (e *EmitChain) Retry(r int) *EmitChain {
	println("entry Retry")
	e.retry = r
	return e
}

func (e *EmitChain) BuildBodyToPcc(f func() (BodyInf, error)) *EmitChain {
	println("entry Retry")
	inf, err := f()
	if err != nil {
		errors.New("BuildBodyToPcc error")
	}
	e.bodyInf = inf
	return e
}

func (e *EmitChain) BuildHead(f func() (HeaderInf, error)) *EmitChain {
	println("entry Retry")
	inf, err := f()
	if err != nil {
		errors.New("BuildHead error")
	}
	e.bodyInf = inf
	return e
}

func (e *EmitChain) constructMessage(f func() error) *EmitChain {
	println("entry Retry")
	f()
	return e
}

func (e *EmitChain) Emit() (interface{}, error) {
	return e.emitByParam(e.timeout, e.retry)
}

func (e *EmitChain) emitByParam(timeout time.Duration, retry int) (interface{}, error) {
	println("entry Emit")
	rsp, err := e.send.emit(nil) // todo 这里回头再考虑一下
	if err != nil {
		println("emit fail")
		for i := 0; i < retry; i++ {
			time.Sleep(timeout)
			rsp, err = e.send.emit(nil) // todo 这里回头再考虑一下
			if err == nil {
				break
			}
			// todo 打印失败原因，继续重试
		}
		if err != nil {
			return nil, err
		}
		return rsp, err
	}
	println("emit success")
	return rsp, nil
}
