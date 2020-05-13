# Start from the latest golang base image
FROM golang:latest as builder

LABEL maintainer="Afatek Developers <developer@afatek.com.tr>"

WORKDIR /app

COPY devafatekpingcontainer/go.mod devafatekpingcontainer/go.sum ./

RUN go mod download

COPY devafatekpingcontainer/. .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o devafatekpingcontainer .


######## Start a new stage from scratch #######
FROM alpine:latest  

ENV PORT=10001
RUN apk --no-cache add ca-certificates
RUN apk add --no-cache tzdata
ENV TZ=Europe/Istanbul
WORKDIR /root/

ENV CONTAINER_TYPE=devafatekpingcontainer

COPY --from=builder /app/devafatekpingcontainer .
CMD ["./devafatekpingcontainer"]