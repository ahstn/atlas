FROM golang:latest as builder
WORKDIR /go/src/atlas
COPY . .
RUN go get -u -v golang.org/x/vgo
RUN vgo install


FROM alpine:latest as tz-certs
RUN apk --no-cache add tzdata zip ca-certificates
WORKDIR /usr/share/zoneinfo
# -0 means no compression.  Needed because go's
# tz loader doesn't handle compressed data.
RUN zip -r -0 /zoneinfo.zip .


FROM scratch
COPY --from=builder /go/bin/atlas /atlas
ENV ZONEINFO /zoneinfo.zip
COPY --from=tz-certs /zoneinfo.zip /
COPY --from=tz-certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
ENTRYPOINT ["/atlas"]
