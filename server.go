package main

import (
	"github.com/miekg/dns"
	// "net"
	"fmt"
)

func handleRequest(w dns.ResponseWriter, r *dns.Msg) {
	m := new(dns.Msg)
	m.SetReply(r)
	w.WriteMsg(m)
}

func main() {
	// let's create a simple A resource record
	r := new(dns.A)
	r.Hdr = dns.RR_Header{Name: "8level.ru.", Rrtype: dns.TypeA,
		Class: dns.ClassINET, Ttl: 300}
	r.A = net.IPv4(185, 37, 61, 185)

	// create a simple message
	m := new(dns.Msg)
	m.SetQuestion("8level.ru.", dns.TypeA)

	// creat and start server

	server := &dns.Server{Addr: ":8080", Net: "udp"}

	dns.HandleFunc(".", handleRequest)
	// panic(server.ListenAndServe())
	err := server.ListenAndServe()

	if err != nil {
		fmt.Println(err)
	}
}
