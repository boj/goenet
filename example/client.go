package main

import (
	"goenet"
	"log"
)

const (
	HOST string = "localhost"
)

func main() {
	if goenet.Initialize() == 0 {
		defer goenet.Deinitialize()

		address := &goenet.ENetAddress{}
		address.SetHostName(HOST)
		address.SetPort(5555)

		event := &goenet.ENetEvent{}

		client := goenet.NewHost(nil, 1, 2, 5760/8, 1440/8)
		if client == nil {
			panic("Client Initialization Error")
		} else {
			defer client.Destroy()
		}

		peer := client.Connect(address, 2, 0)
		if peer == nil {
			panic("Client could not connect to host")
		} else {
			log.Print("Client started\n")
		}

		if client.Service(event, 1000) > 0 && event.EventType() == goenet.ENET_EVENT_TYPE_CONNECT {
			log.Print("Connected to server\n")
		} else {
			panic("Failed to connect to server")
		}

		for i := 0; i < 10; i++ {
			data := []byte("Hello")
			packet := goenet.NewPacket(data, len(data), goenet.ENET_PACKET_FLAG_RELIABLE)
			peer.Send(0, packet)
		}

		for {
			for client.Service(event, 1000) > 0 {
				switch event.EventType() {
				case goenet.ENET_EVENT_TYPE_RECEIVE:
					length := event.Packet().DataLength()
					packetData := string(event.Packet().Data())
					channel := event.ChannelID()
					log.Printf("packet - length: %d, data: %s, channel: %d\n", length, packetData, channel)
					event.Packet().Destroy() // clean up
					break
				case goenet.ENET_EVENT_TYPE_DISCONNECT:
					log.Printf("Client disconnected by server\n")
					break
				}
			}
		}
	}
}
