package main

import (
	"fmt"
	"flag"
	"net/http"
	"log"
	"os"
	"time"

	"github.com/openrdap/rdap"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const Version = "0.1.0-dev"

var (
	// CLI flags
	flagAddress = flag.String("address", ":9099", "HTTP listen address")
	flagDomainFile = flag.String("domain-file", "", "Path to file with domains (separated by newlines)")
	flagVersion = flag.Bool("version", false, "Print the rdap_exporter version")

	// Prometheus metrics
	domainExpiration = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "domain_expiration",
			Help: "Days until the RDAP expiration event states this domain will expire",
		},
		[]string{"domain"},
	)

	defaultDateFormat = "2006-01-02T15:04:05Z"
)

func main() {
	flag.Parse()

	if *flagVersion {
		fmt.Println(Version)
		os.Exit(1)
	}

	// read flagDomainFile, grab configs, iff error exit

	prometheus.Register(domainExpiration)
	http.Handle("/metrics", promhttp.Handler())

	//
	// days := math.Floor(date.Sub(time.Now()).Hours() / 24)
	//
	// handler.WithLabelValues(domain).Set(days)

	client := &rdap.Client{} // TODO(Adam): ???
	lookup(client, "banno.com")

	if err := http.ListenAndServe(*flagAddress, nil); err != nil {
		log.Fatalf("ERROR binding to %s: %v", flagAddress, err)
	}
}

func getExpiration(client *rdap.Client, d string) (*time.Time, error) {
	req := &rdap.Request{
		Type: rdap.DomainRequest,
		Query: d,
	}
	resp, err := client.Do(req)

	domain, ok := resp.Object.(*rdap.Domain)
	if !ok {
		return nil, fmt.Errorf("unable to read domain response: %v", err)
	}
	for i := range domain.Events {
		event := domain.Events[i]
		if event.Action == "expiration" {
			when, err := time.Parse(defaultDateFormat, event.Date)
			if err != nil {
				return nil, fmt.Errorf("when parsing %q got error: %v", event.Date, err)
			}
			return &when, nil
		}
	}
	return nil, err
}
