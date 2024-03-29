package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	// 获取系统的单播接口地址列表
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
	}

	var currentIP, currentNetworkHardwareName string

	// 遍历所有地址
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		// = GET LOCAL IP ADDRESS
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				fmt.Println("Current IP address : ", ipnet.IP.String())
				currentIP = ipnet.IP.String()
			}
		}
	}

	fmt.Println("------------------------------")
	fmt.Println("We want the interface name that has the current IP address")
	fmt.Println("MUST NOT be binded to 127.0.0.1 ")
	fmt.Println("------------------------------")

	// get all the system's network interfaces
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println(err)
		return
	}

	// 遍历所有接口
	for _, inter := range interfaces {
		// 找到对应的接口
		if (inter.Flags & net.FlagUp) != 0 {
			fmt.Printf("Found, Name: %s, MAC address: %s\n", inter.Name, inter.HardwareAddr)
		}
	}

	fmt.Println("------------------------------")

	for _, interf := range interfaces {
		if addrs, err := interf.Addrs(); err == nil {
			for index, addr := range addrs {
				fmt.Println("[", index, "]", interf.Name, ">", addr)

				// only interested in the name with current IP address
				if strings.Contains(addr.String(), currentIP) {
					fmt.Println("Use name : ", interf.Name)
					currentNetworkHardwareName = interf.Name
				}
			}
		}
	}

	fmt.Println("------------------------------")

	// extract the hardware information base on the interface name
	// capture above
	netInterface, err := net.InterfaceByName(currentNetworkHardwareName)

	if err != nil {
		fmt.Println(err)
	}

	name := netInterface.Name
	macAddress := netInterface.HardwareAddr

	fmt.Println("Hardware name : ", name)
	fmt.Println("MAC address : ", macAddress)

	// verify if the MAC address can be parsed properly
	hwAddr, err := net.ParseMAC(macAddress.String())

	if err != nil {
		fmt.Println("No able to parse MAC address : ", err)
		os.Exit(-1)
	}

	fmt.Printf("Physical hardware address : %s \n", hwAddr.String())
}
