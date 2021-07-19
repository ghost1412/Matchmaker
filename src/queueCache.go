package matchmaker

type ActivePlayers struct {
	ActivePlayers map[string]bool
}

func NewActivePlayers() *ActivePlayers {
	return &ActivePlayers{make(map[string]bool)}
}

func (ap *ActivePlayers) IsActive(player string) bool {
	return ap.ActivePlayers[player]
}

func (ap *ActivePlayers) Add(player string) {
	ap.ActivePlayers[player] = true
}

func (ap *ActivePlayers) Remove(player string) {
	delete(ap.ActivePlayers, player)
}
