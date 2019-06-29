package alarm_test

import (
	"testing"
	"go-test/src/alarm"
	"go-test/src/alarm/mobile"
	"go-test/src/alarm/wechat"
	"go-test/src/alarm/email"
)

func TestSend(t *testing.T)  {
	//t.Parallel()
	concreteAlarm := alarm.NewConcreteAlarm()

	email := new(email.Email)
	concreteAlarm.Attach(email)

	mobile := new(mobile.Mobile)
	concreteAlarm.Attach(mobile)

	wechat := new(wechat.Wechat)
	concreteAlarm.Attach(wechat)

	//concreteAlarm.Detach(email)

	concreteAlarm.SetValue(999)
}