package alarm_test

import (
	"testing"
	"go-test/src/alarm"
	"go-test/src/alarm/mobile"
	"go-test/src/alarm/wechat"
	"go-test/src/alarm/email"
	"strconv"
	"time"
)

func TestSend(t *testing.T) {
	//t.Parallel()
	concreteAlarm := alarm.NewConcreteAlarm()

	email := new(email.Email)
	concreteAlarm.Attach(email)

	mobile := new(mobile.Mobile)
	concreteAlarm.Attach(mobile)

	wechat := new(wechat.Wechat)
	concreteAlarm.Attach(wechat)

	concreteAlarm.Detach(email)

	go func() {
		time.Sleep(time.Second * 3)
		for {
			if v, ok := <-alarm.AlarmsChan; ok {
				concreteAlarm.SetValue(v)
			} else {
				println("chan close")
				break
			}
		}
	}()

	for i:=1; i<=15; i++ {
		msg := alarm.AlarmMsg("msg:"+strconv.Itoa(i))
		alarm.AlarmsChan <- msg
	}
	close(alarm.AlarmsChan)
	for{}
}