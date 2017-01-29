package core

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/coocood/freecache"
	"github.com/miekg/dns"
	"os"
	"os/signal"
	"syscall"
)

// Setup main setup func
func Setup(c *ProtonConfig) {

	fmt.Print(ProtonLOGO)

	setupLog(c)
	setupStat(c)
	setupCache(c)
	setupService(c)
}

// setupService DNS service
func setupService(c *ProtonConfig) {
	dns.HandleFunc(".", protonHandle)
	go serve("tcp")
	go serve("udp")
	log.Infoln("proton ready")
	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	s := <-sig
	log.Infof("Signal %s received, stopping", s)
}

// serve dns service listener
func serve(prot string) {
	server := &dns.Server{Addr: ":8053", Net: prot, TsigSecret: nil}
	if err := server.ListenAndServe(); err != nil {
		log.Errorf("Failed to run %s service, %s", prot, err.Error())
		os.Exit(1)
	}
}

// setupLog to initialize logrus service
func setupLog(c *ProtonConfig) {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.TextFormatter{})
	// Output to stdout instead of the default stderr, could also be a file.
	//f, _ := os.OpenFile("dns.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	//log.SetOutput(f)
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

// setupCache initialize DNS cache service
func setupCache(c *ProtonConfig) {
	cacheSize := 10 * 1024 * 1024
	cache = freecache.NewCache(cacheSize)
}

// 启动统计
func setupStat(c *ProtonConfig) {
	statistics = NewProtonStat()
}
