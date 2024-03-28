package main

import (
	"errors"
)

type DecisionList struct {
	a        interface{}
	b        interface{}
	handle   func(a, b interface{}) error
	rollback func(a, b interface{}) error
	next     *DecisionList
	pre      *DecisionList
}

func NewDecisionTreeNode() *DecisionList {
	return &DecisionList{}
}

func (d *DecisionList) New(a, b interface{}) *DecisionList {
	d.a = a
	d.b = b
	return d
}

func (d *DecisionList) Handle(f func(interface{}, interface{}) error) *DecisionList {
	d.handle = f
	return d
}

func (d *DecisionList) Rollback(f func(a, b interface{}) error) *DecisionList {
	d.rollback = f
	return d
}

func Iteration(head *DecisionList) error {
	n := head
	var err error
	for n != nil {
		err = n.handle(n.a, n.b)
		if err != nil {
			break
		}
		n = n.next
	}

	if err != nil {
		for n != nil {
			err = n.rollback(n.a, n.b)
			if err != nil {
				break
			}
			n = n.pre
		}
	}

	return nil
}

func main1() {

	sum := int(0)
	node1 := NewDecisionTreeNode()
	node1.New(1, 2).Handle(func(a, b interface{}) error {
		c, d := a.(int), b.(int)
		sum += c + d
		return nil
	}).Rollback(func(a, b interface{}) error {
		c, d := a.(int), b.(int)
		sum -= c + d
		return nil
	})

	node11 := NewDecisionTreeNode()
	node11.New(3, 4).Handle(func(a, b interface{}) error {
		c, d := a.(int), b.(int)
		sum += c + d
		return errors.New("error")
	}).Rollback(func(a, b interface{}) error {
		c, d := a.(int), b.(int)
		sum -= c + d
		return nil
	})

	node111 := NewDecisionTreeNode()
	node111.New(3, 4).Handle(func(a, b interface{}) error {
		c, d := a.(int), b.(int)
		sum += c + d
		return nil
	}).Rollback(func(a, b interface{}) error {
		c, d := a.(int), b.(int)
		sum -= c + d
		return nil
	})

	node1.next = node11
	node11.pre = node1

	node11.next = node111
	node111.pre = node11

	Iteration(node1)
	println(sum)
}
