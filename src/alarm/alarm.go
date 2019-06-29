package alarm

import (
	"container/list"
)

type alarmChan chan string

var alarmChans = make(alarmChan, 10)

type Alarm interface {
	Attach(Observer) //注册观察者
	Detach(Observer) //释放观察者
	Notify()         //通知所有注册的观察者
}

type Observer interface {
	Update(Alarm) 	 //观察者进行更新状态
}

//implements Subject
type ConcreteAlarm struct {
	observers *list.List
	value     int
}

//
func NewConcreteAlarm() *ConcreteAlarm {
	concreteAlarm := new(ConcreteAlarm)
	concreteAlarm.observers = list.New()
	return concreteAlarm
}

func (s *ConcreteAlarm)Attach(observer Observer) {
	s.observers.PushBack(observer)
}

func (s *ConcreteAlarm)Detach(observer Observer) {
	for ob := s.observers.Front(); ob != nil; ob = ob.Next() {
		if ob.Value.(Observer) == observer {
			s.observers.Remove(ob)
			break
		}
	}
}

func (s *ConcreteAlarm)Notify() {
	for ob := s.observers.Front(); ob != nil; ob = ob.Next() {
		ob.Value.(Observer).Update(s)
	}
}

func (s *ConcreteAlarm)SetValue(value int)  {
	s.value = value
	s.Notify()
}

func (s *ConcreteAlarm)GetValue() int {
	return s.value
}