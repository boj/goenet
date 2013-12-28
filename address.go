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
