#coding: utf-8

import os, time
import socket

p = print

local_ip = "192.168.10.12"
remote_mac = "2C-F0-6D-28-3C-10"


def main():
	soc = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
	addr = ("255.255.255.255", 0)

	soc.bind((local_ip, 0))
	soc.setsockopt(socket.SOL_SOCKET, socket.SO_REUSEADDR, 1)
	soc.setsockopt(socket.SOL_SOCKET, socket.SO_BROADCAST, 1)

	packet = make_packet()
	p(len(packet),  packet.hex())

	while 1:
		ret =soc.sendto(packet, addr)
		p(ret)
		cmd = input("TRY >>")
		if cmd == "x":break

	pass

def make_packet():
	packet = bytearray()
	packet += b'\xff'*6

	mac =  remote_mac
	mac = mac.replace('-', '') * 16

	by = bytes.fromhex(mac)
	packet += by

	return packet

if __name__ == "__main__":
	main()





