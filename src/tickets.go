package matchmaker

import (
	"time"
)

type RequestTickets struct {
	Name  string
	Skill int
	Ping  int
	ttl   int64
}

func (r RequestTickets) Json2Player() Player {
	return NewPlayer(r.Name, r.Skill, r.Ping, time.Now().Unix()+r.ttl)
}
