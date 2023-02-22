package WordInformation

import (
	"time"

	"rest.com.tw/tinymud/src/RestGo.MUD.Core/World/Timer/VirtualClock"
)

var StartTime = time.Now().UTC().Add(8 * time.Hour).Format("2006-01-02 15:04:05")
var StartVirtualTime = VirtualClock.GetDateString()

func init() {

}
