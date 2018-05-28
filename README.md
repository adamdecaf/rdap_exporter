# rdap_exporter
Registration Data Access Protocol (RDAP) Prometheus Exporter



### Install / Usage

Docker

```
$ docker run -it -p 9099:9099 -v $(pwd)/testdata:/conf adamdecaf/rdap-exporter:0.1.0-dev -domain-file=/conf/good.domains
2018/05/28 20:32:09 starting rdap_exporter (0.1.0-dev)
2018/05/28 20:32:11 example.cz expires in 458.00 days
```
