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

// tip from https://groups.google.com/d/msg/golang-nuts/WnA9F1cNDy0/-Rsd-hrjhOAJ
static ENetEventType _go_event_get_type(ENetEvent *e) {
	return e->type;
}
*/
import "C"

import (
	"unsafe"
)

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

type ENetAddress C.ENetAddress

func (address *ENetAddress) SetHost(host uint) {
	address.host = C.enet_uint32(host)
}

func (address *ENetAddress) SetPort(port uint) {
	address.port = C.enet_uint16(port)
}

func (address *ENetAddress) SetHostName(hostName string) int {
	hName := C.CString(hostName)
	defer C.free(unsafe.Pointer(hName))
	return int(C.enet_address_set_host((*C.ENetAddress)(address), hName))
}

func (address *ENetAddress) HostIp(hostName string, nameLength int) int {
	hName := C.CString(hostName)
	defer C.free(unsafe.Pointer(hName))
	return int(C.enet_address_get_host_ip((*C.ENetAddress)(address), hName, C.size_t(nameLength)))
}

func (address *ENetAddress) Host(hostName string, nameLength int) int {
	hName := C.CString(hostName)
	defer C.free(unsafe.Pointer(hName))
	return int(C.enet_address_get_host((*C.ENetAddress)(address), hName, C.size_t(nameLength)))
}

type ENetPacket C.ENetPacket

// Currently this method requires that the data be written in byte form.
// Ideally it would be nice to write arbitrary data such as one can do in the C ENet library.
func NewPacket(data []byte, dataLength int, flags ENetPacketFlag) *ENetPacket {
	return (*ENetPacket)(C.enet_packet_create(unsafe.Pointer(&data[0]), C.size_t(dataLength), C.enet_uint32(flags)))
}

func (p *ENetPacket) Data() []byte {
	return (*[1 << 30]byte)(unsafe.Pointer(p.data))[0:p.DataLength()]
}

func (p *ENetPacket) DataLength() int {
	return int(p.dataLength)
}

func (packet *ENetPacket) Destroy() {
	C.enet_packet_destroy((*C.ENetPacket)(packet))
}

func (packet *ENetPacket) Resize(dataLength int) int {
	return int(C.enet_packet_resize((*C.ENetPacket)(packet), C.size_t(dataLength)))
}

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

type ENetHost C.ENetHost

func NewHost(address *ENetAddress, peerCount, channelLimit int, incomingBandwidth, outgoingBandwidth int) *ENetHost {
	return (*ENetHost)(C.enet_host_create((*C.ENetAddress)(address), C.size_t(peerCount), C.size_t(channelLimit), C.enet_uint32(incomingBandwidth), C.enet_uint32(outgoingBandwidth)))
}

func (host *ENetHost) Destroy() {
	C.enet_host_destroy((*C.ENetHost)(host))
}

func (host *ENetHost) Connect(address *ENetAddress, channelCount int, data int) *ENetPeer {
	return (*ENetPeer)(C.enet_host_connect((*C.ENetHost)(host), (*C.ENetAddress)(address), C.size_t(channelCount), C.enet_uint32(data)))
}

func (host *ENetHost) CheckEvents(event *ENetEvent) int {
	return int(C.enet_host_check_events((*C.ENetHost)(host), (*C.ENetEvent)(event)))
}

func (host *ENetHost) Service(event *ENetEvent, timeout int) int {
	return int(C.enet_host_service((*C.ENetHost)(host), (*C.ENetEvent)(event), C.enet_uint32(timeout)))
}

func (host *ENetHost) Flush() {
	C.enet_host_flush((*C.ENetHost)(host))
}

func (host *ENetHost) Broadcast(channelID int, packet *ENetPacket) {
	C.enet_host_broadcast((*C.ENetHost)(host), C.enet_uint8(channelID), (*C.ENetPacket)(packet))
}

func (host *ENetHost) Compress(compressor *ENetCompressor) {
	C.enet_host_compress((*C.ENetHost)(host), (*C.ENetCompressor)(compressor))
}

func (host *ENetHost) CompressWithRangeCoder() int {
	return int(C.enet_host_compress_with_range_coder((*C.ENetHost)(host)))
}

func (host *ENetHost) ChannelLimit(channelLimit int) {
	C.enet_host_channel_limit((*C.ENetHost)(host), C.size_t(channelLimit))
}

func (host *ENetHost) BandwidthLimit(incomingBandwidth, outgoingBandwidth int) {
	C.enet_host_bandwidth_limit((*C.ENetHost)(host), C.enet_uint32(incomingBandwidth), C.enet_uint32(incomingBandwidth))
}

type ENetPeer C.ENetPeer

func (peer *ENetPeer) SetData(data interface{}) {

}

func (peer *ENetPeer) Data() interface{} {
	return "stub"
}

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
