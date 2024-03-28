package main

import "errors"

type PolicyControl struct {
	session   *Session
	emitChain *EmitChain
}

func (p *PolicyControl) Pre(req *PolicyReq) error {
	if req == nil {
		return errors.New("req is nil")
	}

	if req.status != 1 && req.seq == 10086 {
		//对消息状态 & 消息编号进行判断
		println("judge req status & seq")
	}

	if req.header != nil {
		//消息头进行判断
		println("judge req header")
	}

	if req.body != nil {
		println("judge req body")
		err := p.Handle(req.body)
		if err != nil {
			println("handle req body error 1000")
		}

	}

	return nil
}

func (p *PolicyControl) Handle(body *ReqBody) error {
	println("handle req body")
	// constructBodyToPcc
	_, err := p.emitChain.BuildBodyToPcc(func() (BodyInf, error) {
		b := new(RspBody)
		//b.e, b.f, b.d = 1, 2, 3
		b.ruleList = make([]*Rule, 5, 5)

		return b, nil
	}).BuildHead(func() (HeaderInf, error) {
		h := new(Header)
		h.a, h.b, h.c = 1, 2, 3
		return h, nil
	}).Retry(3).Timeout(10).Emit()

	// rsp to PolicyRsp

	err = p.HandlePolicyRsp(new(PolicyRsp))
	if err != nil {
		return errors.New("")
	}
	return nil
}

func (p *PolicyControl) HandlePolicyRsp(rsp *PolicyRsp) error {
	println("handle req body")
	//消息检查

	//err := p.HandlePolicyRspBody(rsp.body)
	//if err != nil {
	//	return err
	//}
	return nil
}

//func (p *PolicyControl) HandlePolicyRspBody(body *RspBody) error {
//	println("handle req body")
//
//	err := p.session.handleRspBody(body)
//	if err != nil {
//		return err
//	}
//	return nil
//}
