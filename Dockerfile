FROM alpine:latest AS ca
RUN apk add -U ca-certificates

FROM scratch
COPY --from=ca /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY bin/rdap-exporter-linux /bin/rdap-exporter
EXPOSE 9099
ENTRYPOINT ["/bin/rdap-exporter"]
CMD [""]
