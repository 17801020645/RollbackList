package main

//
//import (
//	"errors"
//)
//
//type Status int
//
//const (
//	PLUS   Status = iota //加
//	REDUCE               //减
//	RIDE                 //乘
//	EXCEPT               //除
//)
//
//type DecisionList struct {
//	a         interface{}
//	b         interface{}
//	handle    func(a, b interface{}) error
//	rollback  func(a, b interface{}) error
//	t         Status
//	observers []*QosFlow
//	next      *DecisionList
//	pre       *DecisionList
//}
//
//func (d *DecisionList) RegisterObserver(observer Status, rule *Rule) *DecisionList {
//	//d.observers = append(d.observers, observer)
//	return d
//}
//
//func (d *DecisionList) notifyObserver() *DecisionList {
//	//for _, observer := range d.observers {
//	//	//observer
//	//}
//	return d
//}
//
//func NewDecisionTreeNode() *DecisionList {
//	return &DecisionList{}
//}
//
//func (d *DecisionList) New(s Status, a, b interface{}) *DecisionList {
//	//println("Create")
//	d.a = a
//	d.b = b
//	d.t = s
//	return d
//}
//
//func (d *DecisionList) Handle(f func(interface{}, interface{}) error) *DecisionList {
//	//println("Handle")
//	//d.t = t
//	d.handle = f
//	return d
//}
//
//func (d *DecisionList) Rollback(f func(a, b interface{}) error) *DecisionList {
//	//println("Rollback")
//	d.rollback = f
//	return d
//}
//
//func (d *DecisionList) Commit() {
//	//defer func() {
//	//	println("defer")
//	//}()
//	println("Commit")
//	err := d.handle(d.a, d.b)
//	if err != nil {
//		println("error")
//		err = d.rollback(d.a, d.b)
//		if err != nil {
//			return
//		}
//	}
//	println("success")
//
//}
//
//func Iteration(head *DecisionList) error {
//	n := head
//	var err error
//	for n != nil {
//		err = n.handle(n.a, n.b)
//		if err != nil {
//			break
//		}
//		n = n.next
//	}
//
//	if err != nil {
//		for n != nil {
//			err = n.rollback(n.a, n.b)
//			if err != nil {
//				break
//			}
//			n = n.pre
//		}
//	}
//
//	return nil
//}
//
//func main1() {
//
//	sum := int(0)
//	node1 := NewDecisionTreeNode()
//	node1.New(PLUS, 1, 2).Handle(func(a, b interface{}) error {
//		c, d := a.(int), b.(int)
//		sum += c + d
//		return nil
//	}).Rollback(func(a, b interface{}) error {
//		c, d := a.(int), b.(int)
//		sum -= c + d
//		return nil
//	})
//
//	node11 := NewDecisionTreeNode()
//	node11.New(PLUS, 3, 4).Handle(func(a, b interface{}) error {
//		c, d := a.(int), b.(int)
//		sum += c + d
//		return errors.New("error")
//	}).Rollback(func(a, b interface{}) error {
//		c, d := a.(int), b.(int)
//		sum -= c + d
//		return nil
//	})
//
//	node111 := NewDecisionTreeNode()
//	node111.New(PLUS, 3, 4).Handle(func(a, b interface{}) error {
//		c, d := a.(int), b.(int)
//		sum += c + d
//		return nil
//	}).Rollback(func(a, b interface{}) error {
//		c, d := a.(int), b.(int)
//		sum -= c + d
//		return nil
//	})
//
//	node1.next = node11
//	node11.pre = node1
//
//	node11.next = node111
//	node111.pre = node11
//
//	Iteration(node1)
//	println(sum)
//}
