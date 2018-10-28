FROM golang:1.11-alpine3.8 as builder

# Set up GOPATH
ENV GOPATH /go

# Install Git
RUN apk update && \
    apk upgrade && \
    apk add git && \
    rm -rf /var/cache/apk/*

# Workdir
WORKDIR /app

# Add current working directory
COPY . /app

# Build
RUN go build main.go

#========================== Runtime Image ==========================
FROM golang:1.11-alpine3.8 as runtime

RUN apk update && \
    apk upgrade && \
    apk add tzdata && \
    rm -rf /var/cache/apk/*

WORKDIR /app

COPY --from=builder /app .

RUN chmod +x /app

ENTRYPOINT ["/usr/sbin/crond", "-f", "-l", "9", "-L", ""]