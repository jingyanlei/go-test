package mobile

import "go-test/src/alarm"

type Mobile struct {

}

func (e *Mobile)Update(a alarm.Alarm) {
	println("mobile send ok ", a.(*alarm.ConcreteAlarm).GetValue())
}