package go_shark


import (
	"github.com/google/gopacket/pcap"
	"log"
)

func CaptureInterfaces()[]pcap.Interface {

	devices, err := pcap.FindAllDevs();

	if err != nil {
		log.Fatal("Cant find Network Devices")
	}

	return devices
}


