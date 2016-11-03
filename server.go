package main

import (
	"github.com/kirillDanshin/myutils"
	"github.com/miekg/dns"
	"log"
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

func getExternalRecord(r *dns.Msg) []dns.RR {

	// creating new message for sending to external server
	if len(r.Question) == 0 {
		log.Printf("DNS Message %s does not contain question", r)
		// return empty answer
		answer := make([]dns.RR, 0)
		return answer
	}
	question := r.Question[0]

	m := new(dns.Msg)
	m.SetQuestion(question.Name, dns.TypeA)

	// hardcoded external server adress
	servAdress := "8.8.8.8:53"

	// uncomment these lines to get advanced features
	// c := new(dns.Client)
	// in, rtt, err := c.Exchange(m, servAdress)

	// simple UDP request. comment if using dns.Client
	in, err := dns.Exchange(m, servAdress)
	if err != nil {
		log.Printf("Error with getting answer from server %s: %s", servAdress, err)
		// return empty answer
		answer := make([]dns.RR, 0)
		return answer
	}

	return in.Answer

}

func handleRequest(w dns.ResponseWriter, r *dns.Msg) {

	m := new(dns.Msg)
	m.SetReply(r)
	var q dns.Question

	if len(r.Question) == 0 {
		log.Printf("DNS Message %s does not contain question", r)
	} else {
		q = r.Question[0]
	}

	if record, ok := records[q.Name]; ok {
		m.Answer = append(m.Answer, record)
	} else {
		externalRecord := getExternalRecord(r) // type []dns.RR

		for _, record := range externalRecord {
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
