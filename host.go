package goenet

/*
#cgo CFLAGS: -I/usr/local/include/enet
#cgo LDFLAGS: -L/usr/local/lib -lenet

#include "enet.h"
*/
import "C"

type ENetHost C.ENetHost

func NewHost(a *ENetAddress, peerCount, channelLimit int, incomingBandwidth, outgoingBandwidth int) *ENetHost {
	return (*ENetHost)(C.enet_host_create((*C.ENetAddress)(a), C.size_t(peerCount), C.size_t(channelLimit), C.enet_uint32(incomingBandwidth), C.enet_uint32(outgoingBandwidth)))
}

func (h *ENetHost) Destroy() {
	C.enet_host_destroy((*C.ENetHost)(h))
}

func (h *ENetHost) Connect(a *ENetAddress, channelCount int, data int) *ENetPeer {
	return (*ENetPeer)(C.enet_host_connect((*C.ENetHost)(h), (*C.ENetAddress)(a), C.size_t(channelCount), C.enet_uint32(data)))
}

func (h *ENetHost) CheckEvents(event *ENetEvent) int {
	return int(C.enet_host_check_events((*C.ENetHost)(h), (*C.ENetEvent)(event)))
}

func (h *ENetHost) Service(event *ENetEvent, timeout int) int {
	return int(C.enet_host_service((*C.ENetHost)(h), (*C.ENetEvent)(event), C.enet_uint32(timeout)))
}

func (h *ENetHost) Flush() {
	C.enet_host_flush((*C.ENetHost)(h))
}

func (h *ENetHost) Broadcast(channelID int, packet *ENetPacket) {
	C.enet_host_broadcast((*C.ENetHost)(h), C.enet_uint8(channelID), (*C.ENetPacket)(packet))
}

func (h *ENetHost) Compress(compressor *ENetCompressor) {
	C.enet_host_compress((*C.ENetHost)(h), (*C.ENetCompressor)(compressor))
}

func (h *ENetHost) CompressWithRangeCoder() int {
	return int(C.enet_host_compress_with_range_coder((*C.ENetHost)(h)))
}

func (h *ENetHost) ChannelLimit(channelLimit int) {
	C.enet_host_channel_limit((*C.ENetHost)(h), C.size_t(channelLimit))
}

func (h *ENetHost) BandwidthLimit(incomingBandwidth, outgoingBandwidth int) {
	C.enet_host_bandwidth_limit((*C.ENetHost)(h), C.enet_uint32(incomingBandwidth), C.enet_uint32(incomingBandwidth))
}
