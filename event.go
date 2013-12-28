package goenet

/*
#cgo CFLAGS: -I/usr/local/include/enet
#cgo LDFLAGS: -L/usr/local/lib -lenet

#include "enet.h"

// tip from https://groups.google.com/d/msg/golang-nuts/WnA9F1cNDy0/-Rsd-hrjhOAJ
static ENetEventType _go_event_get_type(ENetEvent *e) {
	return e->type;
}
*/
import "C"

type ENetEvent C.ENetEvent

// Since "type" is a reserved Go word, and e._type does not work in this case, it gets wrapped.
func (e *ENetEvent) EventType() ENetEventType {
	return ENetEventType(C._go_event_get_type((*C.ENetEvent)(e)))
}

func (e *ENetEvent) Packet() *ENetPacket {
	return (*ENetPacket)(e.packet)
}

func (e *ENetEvent) ChannelID() int {
	return int(e.channelID)
}

func (e *ENetEvent) DataLength() int {
	return int(e.packet.dataLength)
}

func (e *ENetEvent) Peer() *ENetPeer {
	return (*ENetPeer)(e.peer)
}
