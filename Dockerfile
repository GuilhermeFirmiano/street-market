FROM golang:1.14.2-alpine as builder

RUN grep nobody /etc/passwd > /etc/passwd.nobody \
    && grep nobody /etc/group > /etc/group.nobody \
    && apk --no-cache update \
    && apk add --no-cache ca-certificates git

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o street-market

FROM scratch
WORKDIR /

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/group.nobody /etc/group
COPY --from=builder /etc/passwd.nobody /etc/passwd
USER nobody

COPY --from=builder /app/street-market .
COPY --from=builder /app/config.yaml .

EXPOSE 9001
ENTRYPOINT ["/street-market"]
