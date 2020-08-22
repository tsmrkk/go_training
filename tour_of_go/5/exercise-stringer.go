package main

import (
	"fmt"
	"strconv"
	"strings"
)

type IPAddr [4]byte

type Stringer interface {
	String() string
}

func (ip IPAddr) String() string {
	s := []string{}
	for i := 0; i < 4; i++ {
		s = append(s, strconv.Itoa(int(ip[i])))
	}
	res := strings.Join(s, ".")
	return res
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
