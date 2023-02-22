package Player

import (
	"math/rand"
	"time"

	"rest.com.tw/tinymud/src/RestGo.MUD.Core/Creature"
)

type User struct {
	Creature.Creature
	Password string
	Gender   UserGender
	Settings Configuration
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
