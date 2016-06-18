package main

import (
	"fmt"
	go_shark "./lib"
	"github.com/google/gopacket/pcap"
)

func main() {
	fmt.Println("Welcome to go-shark")

	ifaces := go_shark.CaptureInterfaces();

	 current_iface := new(pcap.Interface);

	for _, iface := range ifaces {
		if(go_shark.HasIp(&iface, "192.168.1.101")) {
			current_iface = &iface;
			break;
		}
	}


	if current_iface != nil {
		fmt.Println("Interface found")

		go_shark.PrintInterface(*current_iface);


	}
}
