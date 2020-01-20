# Start from golang:1.13-alpine base image
FROM alpine

# Set the Current Working Directory inside the container
WORKDIR /go/src/github.com/RoomFlow/backend/usermanagement/search

# Copy the source from the current directory to the Working Directory inside the container
COPY services/usermanagement/main-alpine ./main-alpine

# Copy common files to container
COPY pkg /go/src/github.com/RoomFlow/backend/pkg

# Copy usermanagement certs to container
COPY internal/certs/usermanagement ./certs

# Set location of SSL cert and key
ENV SSL_CERT_PATH ./certs/app.crt
ENV SSL_KEY_PATH ./certs/app.key

# Expose port 10002 to the outside world
EXPOSE 10002

# Run the executable
CMD ["./main-alpine"]