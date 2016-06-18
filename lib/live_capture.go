package go_shark

import (
	"github.com/google/gopacket/pcap"
	"log"
	"github.com/google/gopacket"
	"time"
)

func GetPacketsStream(iface *pcap.Interface) <-chan gopacket.Packet {
	channel := make(chan gopacket.Packet)

	handle, err := pcap.OpenLive(iface.Name, 1024, false, 30 * time.Second)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
		for packet := range packetSource.Packets() {
			channel <- packet
		}
	}()

	return channel
}
