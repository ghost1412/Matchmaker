package matchmaker

import (
	"errors"
	"math"
	"sync"
	"time"
)

type Player struct {
	mux        sync.Mutex
	name       string
	skill      int
	timestamp  int64
	foundParty bool
	delta      int
	party      *Party
	inParty    bool
	inProcess  bool
}

func NewPlayer(name string, skill int) Player {
	return Player{name: name, skill: skill, timestamp: time.Now().Unix(), foundParty: false, delta: 2, party: nil, inParty: false, inProcess: false}
}

func (player *Player) lock() {
	player.mux.Lock()
}

func (player *Player) unlock() {
	player.mux.Unlock()
}

func (player *Player) findParty(parties []*Party) (*Party, error) {
	if player.inProcess {
		return nil, errors.New("player already queued...")
	}
	player.inProcess = true
	goodParties := player.getGoodParties(parties)
	return findOptimalParty(goodParties), nil
}

func (player *Player) getGoodParties(parties []*Party) []*Party {
	var goodParties []*Party
	for _, party := range parties {
		if inSkillGroup(player, party) {
			goodParties = append(goodParties, party)
		}
	}
	return goodParties
}

func findOptimalParty(parties []*Party) *Party {
	if len(parties) == 0 {
		return nil
	}

	bestParty := parties[0]
	maxLen := len(bestParty.players)
	for _, p := range parties {
		if len(p.players) > maxLen {
			maxLen = len(p.players)
			bestParty = p
		}
	}
	var minCrAt int64 = math.MaxInt64
	for _, p := range parties {
		if len(p.players) == maxLen {
			if p.createdAt < minCrAt {
				minCrAt = p.createdAt
				bestParty = p
			}
		}
	}
	return bestParty
}

func inSkillGroup(player *Player, party *Party) bool {
	ps := player.skill
	as := party.avgSkill
	d := player.delta

	return ps >= (as-d) && ps <= (as+d)
}
