package ch5

import (
	"github.com/miekg/dns"
)

func SimpleLookup() {
	var msg dns.Msg
	fqdn := dns.Fqdn("baidu.com")
	msg.SetQuestion(fqdn, dns.TypeA)
	dns.Exchange(&msg, "8.8.8.8:53")
}
