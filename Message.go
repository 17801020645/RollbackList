package main

type PolicyReq struct {
	header *Header
	body   *ReqBody
	status int
	seq    int
}

type PolicyRsp struct {
	header *Header
	body   *RspBody
	status int
	seq    int
}

type HeaderInf interface{}

type Header struct {
	a int
	b int
	c int
}

type BodyInf interface{}

type ReqBody struct {
	ruleList []*Rule
	//sessionRule *SessionRule
	d int
	e int
	f int
}
