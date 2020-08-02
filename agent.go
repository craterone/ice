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

func (a *Agent) GatherCandidates() error {
	a.hostCandidates()
	return nil
}
