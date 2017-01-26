package core

import (
	"fmt"
	"github.com/miekg/dns"
	"os"
	"os/signal"
	"syscall"
)

func Setup() {
	dns.HandleFunc(".", protonHandle)
	go serve("tcp")
	go serve("udp")
	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	s := <-sig
	fmt.Printf("Signal (%s) received, stopping\n", s)
}

func serve(prot string) {
	server := &dns.Server{Addr: ":8053", Net: prot, TsigSecret: nil}
	if err := server.ListenAndServe(); err != nil {
		fmt.Printf("Failed to setup the "+prot+" server: %s\n", err.Error())
	}
}
