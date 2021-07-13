package matchmaker

import (
	"fmt"
	"math/rand"
	"time"
)

func MatchmakerJob(jobs chan *Player, parties *[]*Party, partySize int) {
	for player := range jobs {
		if player.foundParty || player.inProcess {
			continue
		}
		party, err := player.findParty(*parties)
		if err != nil {
			fmt.Println("Player already queued...")
			continue
		}
		if party == nil {
			p := NewParty()
			party = &p
			addParty(parties, party)
		}
		if player.party != party {
			party.addPlayer(player)
		}
		if len(party.players) >= partySize {
			party.markAllFoundParty()
			player_names := ""
			for player := range party.players {
				player_names += party.players[player].name
				if player != len(party.players)-1 {
					player_names += ", "
				}
			}
			fmt.Println("Formed the following party with players : [", player_names, "]", "Avg Skill Set:", party.avgSkill)
			removeParty(parties, party)
		} else {
			extendTicket(player, jobs)
		}
	}
}

func extendTicket(p *Player, jobs chan *Player) {
	timer := time.NewTimer(3 * time.Second)
	go func() {
		<-timer.C
		p.lock()
		if !p.foundParty {
			p.party.removePlayer(p)
			p.delta = p.delta * 2
			p.party = nil
			p.inParty = false
			p.inProcess = false
			p.unlock()
			jobs <- p
		} else {
			p.unlock()
		}
	}()
}

func addParty(parties *[]*Party, party *Party) {
	*parties = append(*parties, party)
}

func removeParty(parties *[]*Party, party *Party) {
	result := make([]*Party, 0)
	for _, p := range *parties {
		if p != party {
			result = append(result, p)
		}
	}
	*parties = result
}

func randSeq(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
