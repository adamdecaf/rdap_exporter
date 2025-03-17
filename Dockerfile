FROM golang:1.24-alpine as builder
RUN adduser -D -g '' --shell /bin/false rdap
WORKDIR /go/src/github.com/adamdecaf/rdap_exporter
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/rdap-exporter .

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /go/src/github.com/adamdecaf/rdap_exporter/bin/rdap-exporter /bin/rdap-exporter
COPY --from=builder /etc/passwd /etc/passwd
USER rdap
EXPOSE 9099
ENTRYPOINT ["/bin/rdap-exporter"]
