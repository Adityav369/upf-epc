// SPDX-License-Identifier: Apache-2.0
// Copyright 2022 Open Networking Foundation

package integration

import "net"

// this file should contain all the struct defs/constants used among different test cases.

const (
	defaultSliceID = 0

	defaultSDFFilter = "permit out udp from any to assigned 80-80"

	ueAddress    = "17.0.0.1"
	upfN3Address = "198.18.0.1"
	nodeBAddress = "198.18.0.10"

	ActionForward uint8 = 0x2
	ActionDrop    uint8 = 0x1
	ActionBuffer  uint8 = 0x4
	ActionNotify  uint8 = 0x8

	srcIfaceAccess = 0x1
	srcIfaceCore   = 0x2

	directionUplink   = 0x1
	directionDownlink = 0x2
)

type pfcpSessionData struct {
	nbAddress    string
	ueAddress    string
	upfN3Address string

	sdfFilter string

	precedence uint32

	ulTEID uint32
	dlTEID uint32

	// QER-related fields
	sessQerID        uint32
	uplinkAppQerID   uint32
	downlinkAppQerID uint32

	// only single QFI is fine, QFI is passed in session QER, but not considered.
	QFI uint8

	sessMBR uint64
	sessGBR uint64

	// uplink/downlink GBR/MBR is always the same
	appMBR uint64
	appGBR uint64
}

type portRange struct {
	low  uint16
	high uint16
}

type appFilter struct {
	proto        uint8
	appIP        net.IP
	appPrefixLen uint32
	appPort      portRange
}

type p4RtValues struct {
	tunnelPeerID uint8
	appID        uint8
	appFilter    appFilter
}

func (af appFilter) isEmpty() bool {
	return af.proto == 0 && len(af.appIP) == 0 &&
		af.appPort.low == 0 && af.appPort.high == 0
}
