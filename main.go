package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"time"
)

func main() {

	lookup("google.com", false, "ip")
	lookup("google.com", false, "ip4")
	lookup("google.com", false, "ip6")
	lookup("google.com", true, "ip")
	lookup("google.com", true, "ip4")
	lookup("google.com", true, "ip6")
}

func lookup(name string, usego bool, network string) {
	ctx := context.Background()

	r := net.Resolver{
		PreferGo:     usego,
		StrictErrors: false,
	}

	t0 := time.Now()
	ips, err := r.LookupIP(ctx, network, "google.com")
	tt := time.Since(t0)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not get IPs: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("duration: %v name %v num IPs found %v    usego:%v network:%v\n", tt, name, len(ips), r.PreferGo, network)
}
