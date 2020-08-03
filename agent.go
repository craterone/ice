package ice

import (
	"fmt"
	"net"
)

type Agent struct {
}

func NewAgent() (*Agent, error) {
	agent := &Agent{}
	return agent, nil
}

//https://github.com/pion/transport/blob/a27a8ecf59bcc276ad5d783d8dc99d3406b81591/vnet/net.go
func  (a *Agent) hostCandidates() {
	ifaces, err := net.Interfaces()
	if err != nil {
		fmt.Print(fmt.Errorf("localAddresses: %+v\n", err.Error()))
		return
	}
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			fmt.Print(fmt.Errorf("localAddresses: %+v\n", err.Error()))
			continue
		}
		for _, a := range addrs {
			fmt.Println(a.String())
		}
	}
}

func (a *Agent) getAllIPAddress(ipv6 bool, filterLocal bool) []net.IP {
	var ips []net.IP

	ifaces, _ := net.Interfaces()

	for _, ifc := range ifaces {
		addrs, err := ifc.Addrs()
		if err != nil {
			continue
		}

		for _, addr := range addrs {
			var ip net.IP
			if ipNet, ok := addr.(*net.IPNet); ok {
				ip = ipNet.IP
			} else if ipAddr, ok := addr.(*net.IPAddr); ok {
				ip = ipAddr.IP
			} else {
				continue
			}

			if !ipv6 {
				if ip.To4() != nil && filterLocal && ip.String() != "127.0.0.1" {
					ips = append(ips, ip)
				}
			}
		}
	}

	return ips
}

func (a *Agent) GatherCandidates() error {
	//a.hostCandidates()
	ips := a.getAllIPAddress(false)
	for _,ip := range ips  {
		fmt.Println(ip)
	}
	return nil
}
