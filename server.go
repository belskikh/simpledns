package main

import (
	"fmt"
	"github.com/miekg/dns"
	"net"
)

func main() {
	// let's create a simple A resource record
	r := new(dns.A)
	r.Hdr = dns.RR_Header{Name: "8level.ru.", Rrtype: dns.TypeA,
		Class: dns.ClassINET, Ttl: 300}
	r.A = net.IPv4(185, 37, 61, 185)
	// testing value
	fmt.Println(r)
}
