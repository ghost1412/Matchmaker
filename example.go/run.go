package main

import (
	matchmaker "matchmaker"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	partySize := 5
	parties := make([]*matchmaker.Party, 0)
	tickerChannel := make(chan *matchmaker.Player, 100)
	connectionsMap := map[string]*websocket.Conn{}
	connections := map[*websocket.Conn]bool{}
	upgrader := websocket.Upgrader{}
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	go matchmaker.MatchmakerJob(tickerChannel, &parties, partySize)

	httpServer := matchmaker.MatchmakingServer{tickerChannel, connectionsMap, connections, upgrader}
	httpServer.Start()
}
