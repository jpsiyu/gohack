package ch5

import (
	"fmt"

	"github.com/miekg/dns"
)

func SimpleLookup() {
	var msg dns.Msg
	fqdn := dns.Fqdn("baidu.com")
	msg.SetQuestion(fqdn, dns.TypeA)
	backMsg, err := dns.Exchange(&msg, "8.8.8.8:53")
	if err != nil {
		panic(err)
	}
	if len(backMsg.Answer) == 0 {
		fmt.Println("No records")
		return
	}
	for _, answer := range backMsg.Answer {
		if a, ok := answer.(*dns.A); ok {
			fmt.Println(a.A)
		}
	}
}
