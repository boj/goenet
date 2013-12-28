package goenet

/*
#cgo CFLAGS: -I/usr/local/include/enet
#cgo LDFLAGS: -L/usr/local/lib -lenet

#include "enet.h"
*/
import "C"

import (
	"reflect"
	"unsafe"
)

type PeerData struct {
	Value interface{}
}

type ENetPeer C.ENetPeer

// Private

func (peer *ENetPeer) ConnectID() int {
	return int(peer.connectID)
}

// SetData can set a reference to any arbitrary Go data.
func (peer *ENetPeer) SetData(data interface{}) {
	pd := &PeerData{
		Value: data,
	}
	peer.data = unsafe.Pointer(pd)
}

// Data returns referenced Go data.  Must be dereferenced as a pointer.
//
//   peer.Data().(*MyType)
func (peer *ENetPeer) Data() interface{} {
	return reflect.NewAt(reflect.TypeOf(PeerData{}), peer.data).Elem().Interface().(PeerData).Value
}

// Public

func (peer *ENetPeer) Send(channelID int, packet *ENetPacket) int {
	return int(C.enet_peer_send((*C.ENetPeer)(peer), C.enet_uint8(channelID), (*C.ENetPacket)(packet)))
}

func (peer *ENetPeer) Receive(channelID *C.enet_uint8) *ENetPacket {
	return (*ENetPacket)(C.enet_peer_receive((*C.ENetPeer)(peer), channelID))
}

func (peer *ENetPeer) Ping() {
	C.enet_peer_ping((*C.ENetPeer)(peer))
}

func (peer *ENetPeer) PingInterval(pingInterval int) {
	C.enet_peer_ping_interval((*C.ENetPeer)(peer), C.enet_uint32(pingInterval))
}

func (peer *ENetPeer) Timeout(timeoutLimit, timeoutMinimum, timeoutMaximum int) {
	C.enet_peer_timeout((*C.ENetPeer)(peer), C.enet_uint32(timeoutLimit), C.enet_uint32(timeoutMinimum), C.enet_uint32(timeoutMaximum))
}

func (peer *ENetPeer) Reset() {
	C.enet_peer_reset((*C.ENetPeer)(peer))
}

func (peer *ENetPeer) Disconnect(data int) {
	C.enet_peer_disconnect((*C.ENetPeer)(peer), C.enet_uint32(data))
}

func (peer *ENetPeer) DisconnectNow(data int) {
	C.enet_peer_disconnect_now((*C.ENetPeer)(peer), C.enet_uint32(data))
}

func (peer *ENetPeer) DisconnectLater(data int) {
	C.enet_peer_disconnect_later((*C.ENetPeer)(peer), C.enet_uint32(data))
}

func (peer *ENetPeer) ThrottleConfigure(interval, acceleration, deceleration int) {
	C.enet_peer_throttle_configure((*C.ENetPeer)(peer), C.enet_uint32(interval), C.enet_uint32(acceleration), C.enet_uint32(deceleration))
}
