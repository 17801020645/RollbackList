package main

type QosFlow struct {
	name      string
	ruleLists []*Rule //所有topic相同的Rule在一个QosFlow里面
	ruleMap   map[string]*Rule
}

const (
	NEWQOSFLOW Status = iota
	UPDATEQOSFLOW
	DELETEQOSFLOW
)

var HandleQosFLow map[Status]func(rule Rule) error

func initQosFlow() {

	HandleQosFLow[NEWQOSFLOW] = NewQosFlow
	HandleQosFLow[UPDATEQOSFLOW] = UpdateQosFlow
	HandleQosFLow[DELETEQOSFLOW] = DeleteQosFlow

}

func NewQosFlow(rule Rule) error {
	return nil
}
func UpdateQosFlow(rule Rule) error {
	return nil
}
func DeleteQosFlow(rule Rule) error {
	return nil
}

func (q *QosFlow) update(s Status) {
	//f := HandleQosFLow[s]
	////err := f()
	//if err != nil {
	//	return
	//}
}
