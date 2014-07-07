package goenet

/*
#cgo CFLAGS: -I/usr/local/include/enet
#cgo LDFLAGS: -L/usr/local/lib -lenet

#include "enet.h"
*/
import "C"

import (
	"unsafe"
)

type ENetPeer C.ENetPeer

// Private

func (p *ENetPeer) ConnectID() int {
	return int(p.connectID)
}

// SetData can set a reference to any arbitrary Go data.
func (p *ENetPeer) SetData(data interface{}) {
	p.data = unsafe.Pointer(&data)
}

// Data returns referenced Go data.  Must be dereferenced as a pointer.
//
//   peer.Data().(*MyType)
func (p *ENetPeer) Data() interface{} {
	return unsafe.Pointer(p.data)
}

// Public

func (p *ENetPeer) Send(channelID int, packet *ENetPacket) int {
	return int(C.enet_peer_send((*C.ENetPeer)(p), C.enet_uint8(channelID), (*C.ENetPacket)(packet)))
}

func (p *ENetPeer) Receive(channelID *C.enet_uint8) *ENetPacket {
	return (*ENetPacket)(C.enet_peer_receive((*C.ENetPeer)(p), channelID))
}

func (p *ENetPeer) Ping() {
	C.enet_peer_ping((*C.ENetPeer)(p))
}

func (p *ENetPeer) PingInterval(pingInterval int) {
	C.enet_peer_ping_interval((*C.ENetPeer)(p), C.enet_uint32(pingInterval))
}

func (p *ENetPeer) Timeout(timeoutLimit, timeoutMinimum, timeoutMaximum int) {
	C.enet_peer_timeout((*C.ENetPeer)(p), C.enet_uint32(timeoutLimit), C.enet_uint32(timeoutMinimum), C.enet_uint32(timeoutMaximum))
}

func (p *ENetPeer) Reset() {
	C.enet_peer_reset((*C.ENetPeer)(p))
}

func (p *ENetPeer) Disconnect(data int) {
	C.enet_peer_disconnect((*C.ENetPeer)(p), C.enet_uint32(data))
}

func (p *ENetPeer) DisconnectNow(data int) {
	C.enet_peer_disconnect_now((*C.ENetPeer)(p), C.enet_uint32(data))
}

func (p *ENetPeer) DisconnectLater(data int) {
	C.enet_peer_disconnect_later((*C.ENetPeer)(p), C.enet_uint32(data))
}

func (p *ENetPeer) ThrottleConfigure(interval, acceleration, deceleration int) {
	C.enet_peer_throttle_configure((*C.ENetPeer)(p), C.enet_uint32(interval), C.enet_uint32(acceleration), C.enet_uint32(deceleration))
}
