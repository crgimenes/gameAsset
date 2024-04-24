package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/grandcat/zeroconf"
)

func findServers() {
	service := "_gameAssets._tcp"
	domain := "local"
	waitTime := 3
	resolver, err := zeroconf.NewResolver(nil)
	if err != nil {
		log.Fatalln("Failed to initialize resolver:", err.Error())
	}

	entries := make(chan *zeroconf.ServiceEntry)
	go func(results <-chan *zeroconf.ServiceEntry) {
		for entry := range results {
			fmt.Println("Found service:", entry.ServiceInstanceName())

			fmt.Println("Text:")
			for _, v := range entry.Text {
				fmt.Printf("\t%s\n", v)
			}

			fmt.Println("IPv4:")
			for _, v := range entry.AddrIPv4 {
				fmt.Printf("\t%s\n", v)
			}

			fmt.Println("IPv6:")
			for _, v := range entry.AddrIPv6 {
				fmt.Printf("\t%s\n", v)
			}

			fmt.Println("Port:", entry.Port)
			fmt.Println("Host:", entry.HostName)
			fmt.Println("Domain:", entry.Domain)
		}
	}(entries)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(waitTime))
	defer cancel()
	err = resolver.Browse(ctx, service, domain, entries)
	if err != nil {
		log.Fatalln("Failed to browse:", err.Error())
	}

	<-ctx.Done()
}

func main() {

	go findServers()

	// Server
	server, err := zeroconf.Register(
		"gameAssets",
		"_gameAssets._tcp",
		"local.",
		42424,
		[]string{
			"a=0",
			"b=1",
			"c=2",
		},
		nil,
	)
	if err != nil {
		panic(err)
	}
	defer server.Shutdown()

	<-time.After(50 * time.Second)
}
