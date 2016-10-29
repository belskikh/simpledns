package main

import (
	"fmt"
	"github.com/miekg/dns"
	"net"
)

func handlerRequest(w dns.ResponseWriter, r *dns.Msg) {
	m := new(dns.Msg)
	m.SetReply(r)
	// if r.IsTsig() != nil {
	// 	if w.TsigStatus() == nil {
	// 		// *Msg r has an TSIG record and it was validated
	// 		m.SetTsig("axfr.", dns.HmacMD5, 300, time.Now().Unix())
	// 	} else {
	// 		// *Msg r has an TSIG records and it was not valided
	// 	}
	// }
	w.WriteMsg(m)
	fmt.Println("hello!")
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

	server := &dns.Server{Addr: ":53", Net: "upd"}
	go server.ListenAndServe()
	dns.HandleFunc(".", handlerRequest)
}
