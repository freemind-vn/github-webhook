# FROM scratch
# FROM docker.io/alpine
# FROM docker.io/nginx:alpine-slim
FROM docker.io/debian:trixie-slim

ARG arch=amd64

WORKDIR /home/webhook

COPY --chmod=001 build/webhook-linux-${arch} /usr/local/bin/webhook
COPY build/data/ data/

# RUN apk --no-cache add ca-certificates

EXPOSE 8080

CMD	[ \
	"webhook", \
	"serve" \
	]

