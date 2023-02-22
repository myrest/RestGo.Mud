package ExitsPosition

import "rest.com.tw/tinymud/src/RestGo.MUD.Core.Command/CommonServices/StructCollection"

type Exit struct {
	RoomID    int        //房間代碼
	KeyNumber int        //鑰匙代碼
	CanPick   bool       //可否被技能Pick開鎖
	Status    DoorStatus //門的狀態
}

func (e *Exit) Lock() StructCollection.ActionResult {
	switch e.Status {
	case Locked:
		return StructCollection.CompletedWithDoNothing
	case Open:
		return StructCollection.Failed
	default:
		e.Status = Locked
		return StructCollection.CompletedWithSuccess
	}
}
