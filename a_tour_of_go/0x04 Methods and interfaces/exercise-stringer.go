package main

import (
	"fmt"
)

type IPAddr [4]byte

// 通过让 IPAddr 类型实现 fmt.Stringer 来打印点号分隔的地址
func (ip IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])
}

//func (ip IPAddr) String() string {
//	s := make([]string, len(ip))
//	for i, val := range ip {
//		s[i] = strconv.Itoa(int(val))
//	}
//	return strings.Join(s, ".")
//}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
