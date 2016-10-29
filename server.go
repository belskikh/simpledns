package main

import (
	"log"
	"github.com/kirillDanshin/myutils"
	"github.com/miekg/dns"
)

// init sample storage to test server
var records map[string]dns.RR

func init() {
	records = make(map[string]dns.RR)

	rr, err := dns.NewRR("8level.ru. 300 IN A 185.37.61.185")
	myutils.LogFatalError(err)
	records["8level.ru."] = rr

	rr, err = dns.NewRR("panel.8level.ru. 300 IN A 185.37.61.185")
	myutils.LogFatalError(err)

	records["panel.8level.ru."] = rr
}

func serve(net string) {
	server := &dns.Server{Addr: ":53", Net: net}
	defer server.Shutdown()
	
	err := server.ListenAndServe()

	if err != nil {
		log.Printf("Error - %s", err)
		return
	}
}

func handleRequest(w dns.ResponseWriter, r *dns.Msg) {

	m := new(dns.Msg)
	m.SetReply(r)

	for _, q := range r.Question {
		if record, ok := records[q.Name]; ok {
			m.Answer = append(m.Answer, record)
		}
	}

	w.WriteMsg(m)
}

func main() {
	dns.HandleFunc(".", handleRequest)

	go serve("tcp")
	// block
	serve("udp")
}
