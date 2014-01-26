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

func (a *ENetAddress) SetHost(h uint) {
	a.host = C.enet_uint32(h)
}

func (a *ENetAddress) SetPort(port uint) {
	a.port = C.enet_uint16(port)
}

func (a *ENetAddress) SetHostName(hostName string) int {
	hName := C.CString(hostName)
	defer C.free(unsafe.Pointer(hName))
	return int(C.enet_address_set_host((*C.ENetAddress)(a), hName))
}

func (a *ENetAddress) HostIp(hostName string, nameLength int) int {
	hName := C.CString(hostName)
	defer C.free(unsafe.Pointer(hName))
	return int(C.enet_address_get_host_ip((*C.ENetAddress)(a), hName, C.size_t(nameLength)))
}

func (a *ENetAddress) Host(hostName string, nameLength int) int {
	hName := C.CString(hostName)
	defer C.free(unsafe.Pointer(hName))
	return int(C.enet_address_get_host((*C.ENetAddress)(a), hName, C.size_t(nameLength)))
}
