package goenet

/*
#cgo CFLAGS: -I/usr/local/include/enet
#cgo LDFLAGS: -L/usr/local/lib -lenet

#include "enet.h"
*/
import "C"

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
