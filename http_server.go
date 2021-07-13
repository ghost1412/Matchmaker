package matchmaker

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type RequestTickets struct {
	Name  string
	Skill int
}

type CreateWebsocketConnectionRequest struct {
	Name string
}

type MatchmakingServer struct {
	TicketChan     chan *Player
	ConnectionsMap map[string]*websocket.Conn
	Connections    map[*websocket.Conn]bool
	Upgrader       websocket.Upgrader
}

func (r RequestTickets) Json2Player() Player {
	return NewPlayer(r.Name, r.Skill)
}

func (server MatchmakingServer) handlePlayerTickets(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var req RequestTickets
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&req)
		if err != nil {
			fmt.Println("Error in handling enqueue player ", err)
			return
		}
		player := req.Json2Player()
		server.addPlayerTicket(&player)
	}
}

func (server MatchmakingServer) addPlayerTicket(p *Player) {
	server.TicketChan <- p
}

func (server MatchmakingServer) Start() {
	http.HandleFunc("/search", server.handlePlayerTickets)
	fmt.Printf("Matchmaking server started...\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
