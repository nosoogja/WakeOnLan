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
}

func MakePacket(mac string) []byte {
	packet := []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff}

	mac = ut.Replace(mac, "-", "")

	byMac := make([]byte, 100)
	nn, err := hex.Decode(byMac, []byte(mac))
	ut.IsNotNil(err, "%v", err)

	for i := 0; i < 16; i++ {
		packet = append(packet, byMac[:nn]...)
	}

	return packet

}
