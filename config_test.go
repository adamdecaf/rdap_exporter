package main

import (
	"strings"
	"testing"
)

func TestConfig__good(t *testing.T) {
	domains, err := readDomainFile("testdata/good.domains")
	if err != nil {
		t.Fatal(err)
	}
	if len(domains) == 0 {
		t.Fatal("no domains found")
	}
}

func TestConfig__empty(t *testing.T) {
	domains, err := readDomainFile("testdata/empty.domains")
	if err == nil {
		t.Fatal("expected error")
	}
	if !strings.Contains(err.Error(), "no domains found") {
		t.Fatalf("got %q", err.Error())
	}
	if len(domains) != 0 {
		t.Fatalf("got %v domains", domains)
	}
}
