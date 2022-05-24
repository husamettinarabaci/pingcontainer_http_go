# Start from the latest golang base image
FROM golang:latest as builder

LABEL maintainer="Hsmtek Developers <developer@hsmteknoloji.com>"

WORKDIR /app

COPY devhsmtekpingcontainer/go.mod devhsmtekpingcontainer/go.sum ./

RUN go mod download

COPY devhsmtekpingcontainer/. .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o devhsmtekpingcontainer .


######## Start a new stage from scratch #######
FROM alpine:latest  

ENV PORT=10001
RUN apk --no-cache add ca-certificates
RUN apk add --no-cache tzdata
ENV TZ=Europe/Istanbul
WORKDIR /root/

ENV CONTAINER_TYPE=devhsmtekpingcontainer

COPY --from=builder /app/devhsmtekpingcontainer .
CMD ["./devhsmtekpingcontainer"]