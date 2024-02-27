package sms

type Driver interface {
	// 发送短信
	Send(pthone string, message Message, config map[string]string) bool
}
