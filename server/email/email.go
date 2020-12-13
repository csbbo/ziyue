package email

import (
	"github.com/jordan-wright/email"
	"github.com/juju/errors"
	"net/smtp"
	"server/loger"
	"time"
)

var myEmail *email.Pool

func Init() {
	p, err := email.NewPool(
		"smtp.qq.com:25",
		4,
		smtp.PlainAuth("", "599883519@qq.com", "jnwflkeolinmbcad", "smtp.qq.com"),
	)
	if err != nil {
		loger.Loger.Error(errors.ErrorStack(errors.Trace(err)))
		panic(err.Error())
	}
	myEmail = p
}

func GetConn() *email.Pool {
	return myEmail
}

func Send(toEmail string, subject string, code string) error {
	e := email.NewEmail()
	e.From = "子曰 <599883519@qq.com>"
	e.To = []string{toEmail}
	e.Subject = subject
	e.HTML = []byte(SendCodeTPl(code))

	p := GetConn()
	err := p.Send(e, 10*time.Second)
	if err != nil {
		return err
	}
	return nil
}
