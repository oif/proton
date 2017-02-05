package core

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/coocood/freecache"
	"github.com/miekg/dns"
	"github.com/oif/proton/gdns"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

// Setup main setup func
func Setup(c *ProtonConfig) {

	fmt.Print(ProtonLOGO)

	setupLog(c)
	setupStat(c)
	setupProxy(c)
	setupCache(c)
	setupService(c)
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

// setupStat setup statistics service
func setupStat(c *ProtonConfig) {
	statistics = NewProtonStat()
}

// setupProxy set proxy for proton
func setupProxy(c *ProtonConfig) {
	gdns.SetProxyAddr(c.Proxy.Protocol, c.Proxy.Addr, c.Proxy.Port)
}

// setupCache initialize DNS cache service
func setupCache(c *ProtonConfig) {
	cache = freecache.NewCache(c.Cache.Size * 1024) // kb
}

// setupService DNS service
func setupService(c *ProtonConfig) {

	setPublicIP()
	refreshHost()

	dns.HandleFunc(".", protonHandle)
	tcpAddr := fmt.Sprintf("%s:%d", c.TCP.Addr, c.TCP.Port)
	udpAddr := fmt.Sprintf("%s:%d", c.UDP.Addr, c.UDP.Port)
	go serve(tcpAddr, "tcp")
	go serve(udpAddr, "udp")
	log.Infoln("proton ready")
	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	s := <-sig
	log.Infof("Signal %s received, stopping", s)
}

func setPublicIP() {
	resp, err := http.Get("http://myip.ipip.net")
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		log.Error("error " + err.Error())
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error("error " + err.Error())
	}

	data := strings.Split(string(body), "：")
	if len(data) > 2 {
		ip := strings.Split(data[1], " ")
		servicePublicIP = ip[0]
		log.Infof("Public IP: %s => %s", ip[0], data[2])
		return
	}
	// Error
	log.Debug(string(body))
	log.Error("fail to get public IP, please check the network")
	os.Exit(1)
}

// serve dns service listener
func serve(addr, prot string) {
	server := &dns.Server{Addr: addr, Net: prot, TsigSecret: nil}
	if err := server.ListenAndServe(); err != nil {
		log.Errorf("Failed to run %s service, %s", prot, err.Error())
		os.Exit(1)
	}
}
