package ch5

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/miekg/dns"
)

func lookupA(fqds, serverAddr string) ([]string, error) {
	var m dns.Msg
	var ips []string
	m.SetQuestion(dns.Fqdn(fqds), dns.TypeA)
	in, err := dns.Exchange(&m, serverAddr)
	if err != nil {
		return ips, err
	}
	if len(in.Answer) == 0 {
		return ips, errors.New("No answer")
	}
	for _, answer := range in.Answer {
		if a, ok := answer.(*dns.A); ok {
			ips = append(ips, a.A.String())
		}
	}
	return ips, nil
}

func lookupCNAME(fqds, serverAddr string) ([]string, error) {
	var m dns.Msg
	var ips []string
	m.SetQuestion(dns.Fqdn(fqds), dns.TypeCNAME)
	in, err := dns.Exchange(&m, serverAddr)
	if err != nil {
		return ips, err
	}
	if len(in.Answer) == 0 {
		return ips, errors.New("No answer")
	}
	for _, answer := range in.Answer {
		if a, ok := answer.(*dns.CNAME); ok {
			ips = append(ips, a.Target)
		}
	}
	return ips, nil
}

func FlagParser() {
	var (
		flDomain      = flag.String("domain", "", "The domain to perform guessing againest.")
		flWordlist    = flag.String("wordlist", "", "the wordlist to use for guessing.")
		flWorkerCount = flag.Int("c", 100, "The amount of worker to use.")
		flServerAddr  = flag.String("server", "8.8.8.8:53", "The DNS server to use.")
	)
	flag.Parse()

	if *flDomain == "" || *flWordlist == "" {
		fmt.Println("-domain and -woldlist are required.")
		os.Exit(1)
	}
	fmt.Println(*flWorkerCount, *flServerAddr)
}
