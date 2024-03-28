package main

import "container/list"

type Rule struct {
	name  string
	topic string
}

type SessionRule struct {
	name string
}

type Session struct {
	qosFlowList    []*QosFlow //最终的结果
	ruleQosFLowMap map[string]*QosFlow
	ruleMap        map[string]*Rule
	sessionRule    SessionRule
	decisionTree   *DecisionList
}

type RspBody struct {
	ruleList []*Rule
	//sessionRule *SessionRule

}

//收到一个rsp，里面有两个rule，rule1新增->新增QosFlow；rule2 更新；rule3 删除

//func (s *Session) handleRspBody(body *RspBody) error {
//	println("session handle1")
//	// 各种检查略
//
//	s.decisionTree = NewDecisionTreeNode()
//
//	for _, rule := range body.ruleList {
//		if rule != nil && len(rule.topic) != 0 {
//			qosFlow, ok := s.ruleQosFLowMap[rule.topic]
//			if ok {
//				s.decisionTree.Create(rule).Handle(func() error {
//
//				})
//			} else {
//
//			}
//		}
//	}
//
//	return nil
//}
//
//func (s *Session) Decision(rule *Rule) *DecisionList {
//	println("Decision")
//
//	if rule != nil {
//		if oldRule := s.ruleMap[rule.topic]; oldRule == nil {
//			s.decisionTree.Create(rule)
//		} else {
//
//		}
//	}
//
//	return
//}

type QosFlow struct {
	ruleList *list.List
	aa       int
	bb       int
	cc       int
}

type SessionRule struct {
	dd int
	ee int
	ff int
}
