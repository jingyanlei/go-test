package wechat

import "go-test/src/alarm"

type Wechat struct {

}

func (e *Wechat)Update(a alarm.Alarm) {
	println("wechat send ok ", a.(*alarm.ConcreteAlarm).GetValue())
}