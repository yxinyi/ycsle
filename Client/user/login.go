package user

import (
	"csmsg"
)

const (
	LOGIN_NONE = iota
	LOGIN_START
	LOGIN_SUCCESS
)

var (
	funcList = make([]func(), 0)
)

func init() {
	RegLoginFunc(func() {
		Obj.WriteMsgWithMarshal(&csmsg.C2S_Login{})
		Obj.LoginState = LOGIN_START
	})
	EventRegister()
}

func RegLoginFunc(f func()) {
	funcList = append(funcList, f)
}

func LoginFunc(u *MainUser) {
	Obj = u
	for _, funcIt := range funcList {
		funcIt()
	}
}

func EventRegister() {
	csmsg.Processor.SetHandler(&csmsg.S2C_Login{}, func(list []interface{}) {
		Obj.LoginState = LOGIN_SUCCESS
	})
}
