package core

import (
	log "github.com/Sirupsen/logrus"
	"github.com/coocood/freecache"
	"github.com/miekg/dns"
	"os"
	"os/signal"
	"syscall"
)

var statistics *ProtonStat

func Setup(c *ProtonConfig) {
	setupLog(c)
	log.Infoln("logger ready")

	setupStat(c)
	log.Infoln("statistics ready")

	setupCache(c)
	log.Infoln("cache ready")

	setupService(c)
}

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

func serve(prot string) {
	server := &dns.Server{Addr: ":8053", Net: prot, TsigSecret: nil}
	if err := server.ListenAndServe(); err != nil {
		log.Errorf("Failed to run %s service, %s", prot, err.Error())
		os.Exit(1)
	}
}

func setupLog(c *ProtonConfig) {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.TextFormatter{})
	// Output to stdout instead of the default stderr, could also be a file.
	//f, _ := os.OpenFile("dns.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	//log.SetOutput(f)
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

func setupCache(c *ProtonConfig) {
	cacheSize := 10 * 1024 * 1024
	cache = freecache.NewCache(cacheSize)
}

func setupStat(c *ProtonConfig) {
	statistics = NewProtonStat()
}
