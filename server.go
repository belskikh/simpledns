package main

import (
	"fmt"
	"github.com/miekg/dns"
	// "net"
	// "reflect"
)

func serve() {
	server := &dns.Server{Addr: ":53", Net: "udp"}

	err := server.ListenAndServe()
	defer server.Shutdown()

	if err != nil {
		fmt.Printf("Error - %s", err)
	}
}

// func handleQuery(m *dns.Msg) {
// 	var rr dns.RR

// }

func handleRequest(w dns.ResponseWriter, r *dns.Msg) {

	records := make(map[string]dns.RR)

	rr, _ := dns.NewRR("8level.ru. 300 IN A 185.37.61.185")

	records["8level.ru."] = rr

	rr2, _ := dns.NewRR("panel.8level.ru. 300 IN A 185.37.61.185")
	records["panel.8level.ru."] = rr2

	m := new(dns.Msg)
	m.SetReply(r)
	m.Answer = append(m.Answer, records["8level.ru."])
	fmt.Println(records["8level.ru."])
	w.WriteMsg(m)
}

func main() {

	// create and start server
	dns.HandleFunc(".", handleRequest)

	go serve()

	var input string
	fmt.Scanln(&input)
}
