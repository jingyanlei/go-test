package email

import "go-test/src/alarm"

type Email struct {

}

func (e *Email)Update(a alarm.Alarm) {
	println("email send ok ", a.(*alarm.ConcreteAlarm).GetValue())
}