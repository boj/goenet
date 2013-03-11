package main

import (
	"goenet"
	"log"
)

func main() {
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
					log.Printf("Client connected\n")
					break
				case goenet.ENET_EVENT_TYPE_RECEIVE:
					length := event.Packet().DataLength()
					packetData := string(event.Packet().Data())
					channel := event.ChannelID()
					log.Printf("packet - length: %d, data: %s, channel: %d\n", length, packetData, channel)
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
