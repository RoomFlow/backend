# Start from golang:1.13-alpine base image
FROM alpine

# Set the Current Working Directory inside the container
WORKDIR /go/src/github.com/RoomFlow/backend/services/apigateway

# Copy the source from the current directory to the Working Directory inside the container
COPY services/apigateway/main-alpine ./main-alpine

# Copy common files to container
COPY pkg /go/src/github.com/RoomFlow/backend/pkg

# Copy apigateway certs to container
COPY internal/certs/apigateway ./certs

ENV SSL_CERT_PATH ./certs/app.crt
ENV SSL_KEY_PATH ./certs/app.key

# Expose port 8443 to the outside world
EXPOSE 8443

# Run the executable
CMD ["./main-alpine"]