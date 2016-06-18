package go_shark

import (
	"github.com/google/gopacket/pcap"
	"net"
	"fmt"
)

func HasIp(iface *pcap.Interface, ip string) bool {

	for _, addr := range iface.Addresses {
		if(net.IP.String(addr.IP) == ip) {
			return true;
		}
	}

	return  false;
}

func PrintInterface(iface pcap.Interface) {
	fmt.Println("\nName: ", iface.Name)
	fmt.Println("Description: ", iface.Description)
	fmt.Println("Devices addresses: ", iface.Description)
	for _, address := range iface.Addresses {
		fmt.Println("- IP address: ", address.IP)
		fmt.Println("- Subnet mask: ", address.Netmask)
	}
}
