package main

import (
	"context"
	"strings"

	//"fmt"
	"net"

	//	"strings"
	"time"

	"github.com/miekg/dns"
)

func main() {

	lookup("google.com", false, "ip")

}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func lookup(name string, usego bool, network string) {
	ctx := context.Background()
	d := net.Dialer{
		Timeout: time.Millisecond * time.Duration(10000),
	}

	r := &net.Resolver{
		PreferGo: usego,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {

			println("Dialfunc network:", network)
			// if strings.HasPrefix(network,"udp") {
			// 	return nil, fmt.Errorf("no Udp support")
			// }
			return d.DialContext(ctx, "tcp", "8.8.8.8:53")
		},
	}

	x, err := r.LookupNS(ctx, "google.com")
	check(err)
	for _, v := range x {
		println(v.Host)

	}

	//fmt.Printf("duration: %v name %v num IPs found %v    usego:%v network:%v\n", tt, name, len(ips), r.PreferGo, network)
}

func checkPanic(err error) {
	if err != nil {
		panic(err)
	}
}

var resolverDialer = net.Dialer{
	Timeout: time.Second * time.Duration(4),
}

func DialXXXX(ctx context.Context, network string, address string) (net.Conn, error) {

	return resolverDialer.DialContext(ctx, network, "8.8.8.8:53")
}

func makeResolver(overrideAddress string, overrideNetwork string, preferGo bool) net.Resolver {
	if overrideAddress != "" {
		h, _, err := net.SplitHostPort(overrideAddress)
		checkPanic(err)

		if net.ParseIP(h) == nil {
			panic("no hostnames for resolver addresses")
		}
	}

	r := net.Resolver{
		PreferGo:     preferGo,
		StrictErrors: false,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			if overrideAddress != "" {
				address = overrideAddress
			}
			if overrideNetwork != "" {
				network = overrideNetwork
			}
			return resolverDialer.DialContext(ctx, network, overrideAddress)
		},
	}

	return r
}

func lookupIP(ctx context.Context, r net.Resolver, name chan string) {

}

// The original Lego 'checkDNSPropagation()' would lookup the zone nameservers to avoid intermediate caches
// If we made ddns5 records immutable, we wouldnt need to do the same,
// but what the heck, lets also lookup nameservers
func checkDNSPropagationNotLego(ctx context.Context, name string, val string, resolvers []string, dnstype uint16) bool {

	// We ignore the resolvers at this time.
	// Can later be implemented using net.Resolver.Dial func(...) https://golang.org/pkg/net/#Resolver

	r1 := makeResolver("", "", false)
	r2 := makeResolver("", "", true)

	switch dnstype {
	case dns.TypeA:
		ips, err := r.LookupIP(ctx, "ip4", name)
		return ip.String()
	case dns.TypeAAAA:
		ip := net.ParseIP(ipstr)
		return strings.ToLower(ip.String())
		//return strings.ToLower(fmt.Sprintf("%02x%02x:%02x%02x:%02x%02x:%02x%02x:%02x%02x:%02x%02x:%02x%02x:%02x%02x",
		//	ip[0], ip[1], ip[2], ip[3], ip[4], ip[5], ip[6], ip[7], ip[8], ip[9], ip[10], ip[11], ip[12], ip[13], ip[14], ip[15]))
	case dns.TypeTXT:
		return ipstr
	default:
		panic("A or AAAA only")
	}

	return "", nil
}
