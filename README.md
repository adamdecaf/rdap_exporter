# rdap_exporter

Registration Data Access Protocol (RDAP) prometheus exporter. Originally based on [shift/domain_exporter](https://github.com/shift/domain_exporter).

### Metrics

Currently only one metric (`domain_expiration`) is exported ontop of the default `prometheus/client_golang` metrics.

```
# HELP domain_expiration Days until the RDAP expiration event states this domain will expire
# TYPE domain_expiration gauge
domain_expiration{domain="example.cz"} 458
```

### Install / Usage

You can download and run the latest docker image [`adamdecaf/rdap_exporter`](https://hub.docker.com/r/adamdecaf/rdap_exporter/) from the Docker Hub.

Running the image looks like the following:

```
# Using testdata/good.domains from repository
$ docker run -it -p 9099:9099 -v $(pwd)/testdata:/conf adamdecaf/rdap_exporter:0.1.0-dev -domain-file=/conf/good.domains
2018/05/28 21:15:31 starting rdap_exporter (0.1.0-dev)
2018/05/28 21:15:34 example.cz expires in 458.00 days
```

### Example Prometheus Alert

The following alert will be triggered when domains expire within 45 days

```yaml
groups:
 - name: ./domain.rules
   rules:
    - alert: DomainExpiring
      expr: domain_expiration{} < 45
      for: 24h
      labels:
        severity: warning
      annotations:
        description: "{{ $labels.domain }} expires in {{ $value }} days"
```

### Developing / Contributing

If you find a bug, have a question or want more metrics exposed feel free to open either an issue or a Pull Request. I'll try and review it quickly and have it merged.

You can build the sources with `make build`. Run tests with `make test`. Currently we required Go 1.10.
