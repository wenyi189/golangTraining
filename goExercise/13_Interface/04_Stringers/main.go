package main

import "fmt"

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.

func (ip IPAddr) String() string {
	////return fmt.Sprintf("%v.%v.%v.%v",ip[0],ip[1],ip[2],ip[3])
	//var ipaddr string
	//ipaddrLen := len(ip)
	//for i := 0;i < ipaddrLen; i++ {
	//	ipaddr += fmt.Sprintf(".%v",ip[i])
	//}
	ipaddr = ipaddr[1:]
	return ipaddr
}

func main() {
	addrs := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for n, a := range addrs {
		fmt.Printf("%v: %v\n", n, a)
	}
}
