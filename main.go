package main

import (
	"fmt"
	dnslookup "lookup/dnsLookup"
	iplookup "lookup/ipLookup"
)

func main() {

	ip := []string{"8.8.8.8", "104.21.33.123", "19.108.117.229", "173.38.92.29", "94.127.8.32", "56.141.250.184", "163.191.203.165", "91.71.50.27", "89.11.197.14", "61.148.126.78"}
	for _, v := range ip {
		ipResp, err := iplookup.IpLookup(v)
		if err != nil {
			fmt.Println(err)
			continue
		}
		dnsResp, err := dnslookup.DnsLookup(v)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("ipLookup\n%v\ndnslookup\n%v\n", ipResp, dnsResp)
	}
}
