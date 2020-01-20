# Start from golang:1.13-alpine base image
FROM golang:1.13-alpine

# The latest alpine images don't have some tools like (`git` and `bash`).
# Adding git, bash and openssh to the image
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

# Set the Current Working Directory inside the container
WORKDIR /go/src/github.com/RoomFlow/backend/services/apigateway

# Copy the source from the current directory to the Working Directory inside the container
COPY services/apigateway .

# Copy apigateway certs to container
COPY internal/certs/apigateway ./certs

ENV SSL_CERT_PATH ./certs/app.crt
ENV SSL_KEY_PATH ./certs/app.key

COPY pkg /go/src/github.com/RoomFlow/backend/pkg

# Expose port 8443 to the outside world
EXPOSE 8443

# Run the executable
CMD ["./main-alpine"]