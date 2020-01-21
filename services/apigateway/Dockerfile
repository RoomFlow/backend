############################
# STEP 1 build executable binary
############################
FROM golang:alpine AS builder

WORKDIR /go/src/github.com/RoomFlow/backend/services/apigateway

RUN set -ex \
    && apk update \
    && apk add --no-cache git

COPY services/apigateway .

# Copy common files to container
COPY pkg /go/src/github.com/RoomFlow/backend/pkg

RUN go get -d -v

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/bin/main-alpine

############################
# STEP 2 build a small image
############################
FROM alpine

# Set the Current Working Directory inside the container
WORKDIR /go/src/github.com/RoomFlow/backend/services/apigateway

COPY --from=builder /go/bin/main-alpine ./main-alpine

# Copy apigateway certs to container
COPY internal/certs/apigateway ./certs

ENV SSL_CERT_PATH ./certs/app.crt
ENV SSL_KEY_PATH ./certs/app.key

# Expose port 8443 to the outside world
EXPOSE 8443

# Run the executable
CMD ["./main-alpine"]