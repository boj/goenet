package main

import (
	"goenet"
	"log"
)

type Player struct {
	Name string
}

func main() {
	player := &Player{
		Name: "Player1",
	}

	if goenet.Initialize() == 0 {
		defer goenet.Deinitialize()

		address := &goenet.ENetAddress{}
		address.SetHost(goenet.ENET_HOST_ANY)
		address.SetPort(5555)

		event := &goenet.ENetEvent{}

		server := goenet.NewHost(address, 32, 2, 0, 0)
		if server == nil {
			panic("Server Initialization Error")
		} else {
			defer server.Destroy()
		}

		log.Print("Server started\n")
		for {
			for server.Service(event, 1000) > 0 {
				switch event.EventType() {
				case goenet.ENET_EVENT_TYPE_CONNECT:
					peer := event.Peer()
					peer.SetData(player)
					log.Printf("Client connected: %d\n", peer.ConnectID())
					break
				case goenet.ENET_EVENT_TYPE_RECEIVE:
					peer := event.Peer()
					player := peer.Data().(*Player)
					length := event.Packet().DataLength()
					packetData := string(event.Packet().Data())
					channel := event.ChannelID()
					log.Printf("packet - length: %d, data: %s, channel: %d, name: %s\n", length, packetData, channel, player.Name)
					event.Packet().Destroy() // clean up
					break
				case goenet.ENET_EVENT_TYPE_DISCONNECT:
					log.Printf("Client disconnected\n")
					break
				}
			}
		}
	}
}
