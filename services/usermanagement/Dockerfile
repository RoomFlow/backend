# Alpine base image
FROM alpine

# Set the Current Working Directory inside the container
WORKDIR /go/src/github.com/RoomFlow/backend/services/usermanagement

# Copy the source from service directory to the Working Directory inside the container
COPY services/usermanagement/main-alpine ./main-alpine

# Copy service certs to container
COPY internal/certs/usermanagement ./certs

# Set location of SSL cert and key
ENV SSL_CERT_PATH ./certs/app.crt
ENV SSL_KEY_PATH ./certs/app.key

# Expose port to the outside world
EXPOSE 10002

# Run the executable
CMD ["./main-alpine"]
