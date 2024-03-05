package dnsresolver

import (
	"fmt"
	"log"
	"time"

	"github.com/miekg/dns"
)

func resolver(domain string, qtype uint16) []dns.RR {
	msg := new(dns.Msg)
	msg.SetQuestion(domain, qtype)

	c := &dns.Client{Timeout: 5 * time.Second}
	response, _, err := c.Exchange(msg, "8.8.8.8:53")
	if err != nil {
		log.Printf("ERROR : %v\n", err)
		return nil
	}

	if response == nil {
		log.Printf("ERROR : no response from server\n")
		return nil
	}

	return response.Answer
}

type dnsHandler struct{}

func (h *dnsHandler) ServeDNS(w dns.ResponseWriter, r *dns.Msg) {
	msg := new(dns.Msg)
	msg.SetReply(r)
	msg.Authoritative = true

	for _, question := range r.Question {
		answers := resolver(question.Name, question.Qtype)
		msg.Answer = append(msg.Answer, answers...)
	}
	w.WriteMsg(msg)
}

func StartDNSServer() {
	handler := new(dnsHandler)
	server := &dns.Server{
		Addr:      ":53",
		Net:       "udp",
		Handler:   handler,
		ReusePort: true,
	}
	err := server.ListenAndServe()
	if err != nil {
		fmt.Printf("Failed to start server: %s\n", err.Error())
	}
}
