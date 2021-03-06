package main

import (
	"fmt"

	go_shark "./lib"
	"github.com/google/gopacket/pcap"
)

func main() {
	fmt.Println("Welcome to go-shark")

	ifaces := go_shark.CaptureInterfaces()

	currentIface := new(pcap.Interface)

	for _, iface := range ifaces {
		if go_shark.HasIp(&iface, "192.168.1.101") {
			currentIface = &iface
			break
		}
	}

	if current_iface != nil {
		fmt.Println("Interface found")

		go_shark.PrintInterface(*current_iface)

		stream := go_shark.GetPacketsStream(current_iface)

		for packet := range stream {
			fmt.Print(packet)
		}

	}
}
