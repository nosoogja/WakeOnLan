package main

import (
	"encoding/hex"
	"fmt"
	"lib/ut"
)

var p = fmt.Println
var REMOTE_MAC = "34-F2-22-12-12-12"

func main() {
	p(":::::: WakeOnLan ::::::::::")

	packet := MakePacket(REMOTE_MAC)
	p(packet)

	/// todo : send packet.  more later ...........
	// send(REMOTE_MAC, packet)
}

func MakePacket(mac string) []byte {
	packet := []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff}

	mac = ut.Replace(mac, "-", "")

	byMac, err := hex.DecodeString(mac)
	ut.IsNotNil(err, "%v", err)

	for i := 0; i < 16; i++ {
		packet = append(packet, byMac...)
	}

	return packet

}
