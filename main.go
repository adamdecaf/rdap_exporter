package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/openrdap/rdap"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const version = "0.1.2-dev"

var (
	defaultInterval, _ = time.ParseDuration("12h")

	// CLI flags
	flagAddress    = flag.String("address", "0.0.0.0:9099", "HTTP listen address")
	flagDomainFile = flag.String("domain-file", "", "Path to file with domains (separated by newlines)")
	flagInterval   = flag.Duration("interval", defaultInterval, "Interval to check domains at")
	flagVersion    = flag.Bool("version", false, "Print the rdap_exporter version")

	// Prometheus metrics
	domainExpiration = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "domain_expiration",
			Help: "Days until the RDAP expiration event states this domain will expire",
		},
		[]string{"domain"},
	)

	defaultDateFormats = []string{
		"2006-01-02T15:04:05Z",
		time.RFC3339,
	}
)

func init() {
	prometheus.MustRegister(domainExpiration)
}

func main() {
	flag.Parse()

	if *flagVersion {
		fmt.Println(version)
		os.Exit(1)
	}

	log.Printf("starting rdap_exporter (%s)", version)

	// read and verify config file
	if *flagDomainFile == "" {
		log.Fatalf("no -domain-file specified")
	}
	domains, err := readDomainFile(*flagDomainFile)
	if err != nil {
		log.Fatalf("error getting domains %q: %v", *flagDomainFile, err)
	}

	// Setup internal checker
	check := &checker{
		domains:  domains,
		handler:  domainExpiration,
		client:   &rdap.Client{},
		interval: *flagInterval,
	}
	go check.checkAll()

	// Add domain_expiration to /metrics
	h := promhttp.HandlerFor(prometheus.DefaultGatherer, promhttp.HandlerOpts{})
	http.Handle("/metrics", h)

	log.Printf("listenting on %s", *flagAddress)
	if err := http.ListenAndServe(*flagAddress, nil); err != nil {
		log.Fatalf("ERROR binding to %s: %v", *flagAddress, err)
	}
}

type checker struct {
	domains []string
	handler *prometheus.GaugeVec

	client *rdap.Client

	t        *time.Ticker
	interval time.Duration
}

func (c *checker) checkAll() {
	if c.t == nil {
		c.t = time.NewTicker(c.interval)
		c.checkNow() // check domains right away after ticker setup
	}
	for _ = range c.t.C {
		c.checkNow()
	}
}

func (c *checker) checkNow() {
	for i := range c.domains {
		expr, err := c.getExpiration(c.domains[i])
		if err != nil {
			log.Printf("error getting RDAP expiration for %s: %v", c.domains[i], err)
			// Present an invalid value
			c.handler.WithLabelValues(c.domains[i]).Set(0)
			continue
		}
		days := math.Floor(time.Until(*expr).Hours() / 24)
		c.handler.WithLabelValues(c.domains[i]).Set(days)
		log.Printf("%s expires in %.2f days", c.domains[i], days)
	}
}

func (c *checker) getExpiration(d string) (*time.Time, error) {
	req := &rdap.Request{
		Type:  rdap.DomainRequest,
		Query: d,
	}
	resp, err := c.client.Do(req)

	domain, ok := resp.Object.(*rdap.Domain)
	if !ok {
		return nil, fmt.Errorf("unable to read domain response: %v", err)
	}
	for i := range domain.Events {
		event := domain.Events[i]
		if event.Action == "expiration" {
			for j := range defaultDateFormats {
				when, err := time.Parse(defaultDateFormats[j], event.Date)
				if err != nil {
					continue
				}
				return &when, nil
			}
			return nil, fmt.Errorf("unable to find parsable format for %q", event.Date)
		}
	}
	return nil, err
}

func readDomainFile(where string) ([]string, error) {
	fullPath, err := filepath.Abs(where)
	if err != nil {
		return nil, fmt.Errorf("when expanding %s: %v", *flagDomainFile, err)
	}

	fd, err := os.Open(fullPath)
	if err != nil {
		return nil, fmt.Errorf("when opening %s: %v", fullPath, err)
	}
	defer fd.Close()
	r := bufio.NewScanner(fd)

	var domains []string
	for r.Scan() {
		domains = append(domains, strings.TrimSpace(r.Text()))
	}
	if len(domains) == 0 {
		return nil, fmt.Errorf("no domains found in %s", fullPath)
	}
	return domains, nil
}
