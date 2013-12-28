/*
Package goenet is a wrapper for the C based ENet UDP network library.

The majority of the library matches ENet as closely as possible, while adding
idiomatic Go methods to each interface.

Not all things are capable of working equally, so there are a few cases
where this library detours from the underlying ENet library.

The first noteable exception to this is currently the ENetPacket.  This may
change in the future, however, the current method for sending data is to
encode and decode said data using the bytes module.
*/
package goenet

/*
#cgo CFLAGS: -I/usr/local/include/enet
#cgo LDFLAGS: -L/usr/local/lib -lenet

#include "enet.h"
*/
import "C"

const (
	ENET_VERSION_MAJOR int = C.ENET_VERSION_MAJOR
	ENET_VERSION_MINOR int = C.ENET_VERSION_MINOR
	ENET_VERSION_PATCH int = C.ENET_VERSION_PATCH
	ENET_VERSION       int = C.ENET_VERSION
)

type ENetPeerState int
type ENetPacketFlag int
type ENetEventType int

var (
	ENET_HOST_ANY       = uint(C.ENET_HOST_ANY)
	ENET_HOST_BROADCAST = uint(C.ENET_HOST_BROADCAST)
	ENET_PORT_ANY       = uint(C.ENET_PORT_ANY)
)

var (
	ENET_PACKET_FLAG_RELIABLE            = ENetPacketFlag(C.ENET_PACKET_FLAG_RELIABLE)
	ENET_PACKET_FLAG_UNSEQUENCED         = ENetPacketFlag(C.ENET_PACKET_FLAG_UNSEQUENCED)
	ENET_PACKET_FLAG_NO_ALLOCATE         = ENetPacketFlag(C.ENET_PACKET_FLAG_NO_ALLOCATE)
	ENET_PACKET_FLAG_UNRELIABLE_FRAGMENT = ENetPacketFlag(C.ENET_PACKET_FLAG_UNRELIABLE_FRAGMENT)
)

var (
	ENET_PEER_STATE_DISCONNECTED             = ENetPeerState(C.ENET_PEER_STATE_DISCONNECTED)
	ENET_PEER_STATE_CONNECTING               = ENetPeerState(C.ENET_PEER_STATE_CONNECTING)
	ENET_PEER_STATE_ACKNOWLEDGING_CONNECT    = ENetPeerState(C.ENET_PEER_STATE_ACKNOWLEDGING_CONNECT)
	ENET_PEER_STATE_CONNECTION_PENDING       = ENetPeerState(C.ENET_PEER_STATE_CONNECTION_PENDING)
	ENET_PEER_STATE_CONNECTION_SUCCEEDED     = ENetPeerState(C.ENET_PEER_STATE_CONNECTION_SUCCEEDED)
	ENET_PEER_STATE_CONNECTED                = ENetPeerState(C.ENET_PEER_STATE_CONNECTED)
	ENET_PEER_STATE_DISCONNECT_LATER         = ENetPeerState(C.ENET_PEER_STATE_DISCONNECT_LATER)
	ENET_PEER_STATE_DISCONNECTING            = ENetPeerState(C.ENET_PEER_STATE_DISCONNECTING)
	ENET_PEER_STATE_ACKNOWLEDGING_DISCONNECT = ENetPeerState(C.ENET_PEER_STATE_ACKNOWLEDGING_DISCONNECT)
	ENET_PEER_STATE_ZOMBIE                   = ENetPeerState(C.ENET_PEER_STATE_ZOMBIE)
)

var (
	ENET_EVENT_TYPE_NONE       = ENetEventType(C.ENET_EVENT_TYPE_NONE)
	ENET_EVENT_TYPE_CONNECT    = ENetEventType(C.ENET_EVENT_TYPE_CONNECT)
	ENET_EVENT_TYPE_DISCONNECT = ENetEventType(C.ENET_EVENT_TYPE_DISCONNECT)
	ENET_EVENT_TYPE_RECEIVE    = ENetEventType(C.ENET_EVENT_TYPE_RECEIVE)
)

type ENetCallbacks C.ENetCallbacks

type ENetCompressor C.ENetCompressor

func Initialize() int32 {
	return int32(C.enet_initialize())
}

func InitializeWithCallbacks(version int, inits *ENetCallbacks) int {
	return int(C.enet_initialize_with_callbacks(C.ENetVersion(version), (*C.ENetCallbacks)(inits)))
}

func Deinitialize() {
	C.enet_deinitialize()
}
