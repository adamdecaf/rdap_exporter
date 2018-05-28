<img src="https://www.openrdap.org/public/img/logo.png">

OpenRDAP is an command line [RDAP](https://datatracker.ietf.org/wg/weirds/documents/) client implementation in Go.
[![Build Status](https://travis-ci.org/openrdap/rdap.svg?branch=master)](https://travis-ci.org/openrdap/rdap)

https://www.openrdap.org - homepage

https://www.openrdap.org/demo - live demo

## Features
* Command line RDAP client
* Query types supported:
    * ip
    * domain
    * autnum
    * nameserver
    * entity
    * help
    * url
    * domain-search
    * domain-search-by-nameserver
    * domain-search-by-nameserver-ip
    * nameserver-search
    * nameserver-search-by-ip
    * entity-search
    * entity-search-by-handle
* Query bootstrapping (automatic RDAP server URL detection for ip/domain/autnum/(experimental) entity queries)
* Bootstrap cache (optional, uses ~/.openrdap by default)
* X.509 client authentication
* Output formats: text, JSON, WHOIS style
* Experimental [object tagging](https://datatracker.ietf.org/doc/draft-ietf-regext-rdap-object-tag/) support

## Installation

TBD, it's a standard Go progrem (go install ...)

This is under construction...

## Usage

| Query type  | Usage   |
|---|---|
| Domain (.com)   | rdap -v example.com |
| Domain (.みんな) | rdap -v -e nic.みんな  |
| Network | rdap -v 2001:db8:: |
| Autnum | rdap -v AS15169 |
| Entity (test bootstrap) | rdap -v -e 1-VRSN |
| Nameserver | rdap -v -t nameserver -s https://rdap-pilot.verisignlabs.com/rdap/v1 ns1.google.com |
| Help | rdap -v -t help -s https://rdap-pilot.verisignlabs.com/rdap/v1 |
| Domain Search	| rdap -v -t domain-search -s https://rdap-pilot.verisignlabs.com/rdap/v1 exampl*.com |
| Domain Search (by NS)	| rdap -v -t domain-search-by-nameserver -s https://rdap-pilot.verisignlabs.com/rdap/v1 ns1.google.com |
| Domain Search (by NS IP) | rdap -v -t domain-search-by-nameserver-ip -s https://rdap-pilot.verisignlabs.com/rdap/v1 194.72.238.11 |
| Nameserver Search	| rdap -v -t nameserver-search -s https://rdap-pilot.verisignlabs.com/rdap/v1 ns*.yahoo.com |
| Nameserver Search (by IP)	| rdap -v -t nameserver-search-by-ip -s https://rdap-pilot.verisignlabs.com/rdap/v1 194.72.238.11 |
| Entity Search	| rdap -v -t entity-search -s https://rdap-pilot.verisignlabs.com/rdap/v1 Register*-VRSN |
| Entity Search (by handle)	| rdap -v -t entity-search-by-handle -s https://rdap-pilot.verisignlabs.com/rdap/v1 1*-VRSN |

See https://www.openrdap.org/docs.

## Go docs
[![godoc](https://godoc.org/github.com/openrdap/rdap?status.png)](https://godoc.org/github.com/openrdap/rdap)

## Requires
Go 1.7+

## Links
- Wikipedia - [Registration Data Access Protocol](https://en.wikipedia.org/wiki/Registration_Data_Access_Protocol)
- [ICANN RDAP pilot](https://www.icann.org/rdap)

- [OpenRDAP](https://www.openrdap.org)

- https://data.iana.org/rdap/ - Official IANA bootstrap information
- https://test.rdap.net/rdap/ - Test alternate bootstrap service with more experimental RDAP servers

- [RFC 7480 HTTP Usage in the Registration Data Access Protocol (RDAP)](https://tools.ietf.org/html/rfc7480)
- [RFC 7481 Security Services for the Registration Data Access Protocol (RDAP)](https://tools.ietf.org/html/rfc7481)
- [RFC 7482 Registration Data Access Protocol (RDAP) Query Format](https://tools.ietf.org/html/rfc7482)
- [RFC 7483 JSON Responses for the Registration Data Access Protocol (RDAP)](https://tools.ietf.org/html/rfc7483)
- [RFC 7484 Finding the Authoritative Registration Data (RDAP) Service](https://tools.ietf.org/html/rfc7484)

