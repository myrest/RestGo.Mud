package StructCollection

import (
	"fmt"
	"net"

	"rest.com.tw/tinymud/src/RestGo.MUD.Core/Creature/Player"
)

type MudClient struct {
	Conn         net.Conn
	ConnectionID string
	User         Player.User
}

func (s *MudClient) SendMessage(msg ...interface{}) {
	SendMessage(s, msg...)
}

func (s *MudClient) SendFMessage(msg string, args ...interface{}) {
	SendFMessage(s, msg, args...)
}

var SendFMessage = func(s *MudClient, msg string, args ...interface{}) {
	fmt.Fprintf(s.Conn, msg+"\n", args...)
}

var SendMessage = func(s *MudClient, msg ...interface{}) {
	fmt.Fprintln(s.Conn, msg...)
}

type WorldGeography struct { //正常情況下都不會顯示這些資訊
	World  string //世界 : 亞爾斯隆世界 (區域最大)
	Region string //區域 : 新手區 (區域次中)
	Area   string //地區 : 裝備儲藏室 (區域最小)
	RoomID int    //房間號碼
}

type ActionResult int

const (
	CompletedWithSuccess ActionResult = iota
	CompletedWithDoNothing
	Failed
)
