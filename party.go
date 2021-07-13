package matchmaker

import (
	"sync"
	"time"
)

type Party struct {
	mux       sync.Mutex
	id        string
	players   []*Player
	avgSkill  int
	createdAt int64
}

func NewParty() Party {
	return Party{id: randSeq(32), players: []*Player{}, avgSkill: 0, createdAt: time.Now().Unix()}
}

func (party *Party) lock() {
	party.mux.Lock()
}

func (party *Party) unlock() {
	party.mux.Unlock()
}

func (party *Party) addPlayer(player *Player) {
	if party == nil {
		return
	}
	player.lock()
	if player.inParty || player.foundParty {
		player.unlock()
		return
	}
	player.party = party
	player.inParty = true
	player.unlock()
	party.lock()
	defer party.unlock()
	party.players = append(party.players, player)
	party.computeAvgPartySkill()
}

func (party *Party) removePlayer(player *Player) {
	result := make([]*Player, 0)
	party.lock()
	defer party.unlock()
	for _, p := range party.players {
		if p.name != player.name {
			result = append(result, p)
		}
	}
	party.players = result
	party.computeAvgPartySkill()
}

func (party *Party) isEmpty() bool {
	return len(party.players) == 0
}

func (party *Party) computeAvgPartySkill() {
	if party.isEmpty() {
		party.avgSkill = 0
		return
	}
	sum := 0
	for _, p := range party.players {
		sum += p.skill
	}
	party.avgSkill = sum / len(party.players)
}

func (party *Party) markAllFoundParty() {
	party.lock()
	for _, p := range party.players {
		p.foundParty = true
	}
	party.unlock()
}
