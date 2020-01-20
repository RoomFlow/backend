# Start from golang:1.13-alpine base image
FROM alpine

# Set the Current Working Directory inside the container
WORKDIR /go/src/github.com/RoomFlow/backend/services/search

# Copy the source from the current directory to the Working Directory inside the container
COPY services/search/main-alpine ./main-alpine

# Copy common files to container
COPY pkg /go/src/github.com/RoomFlow/backend/pkg

# Copy search certs to container
COPY internal/certs/search ./certs

# Copy firebase credentials to container
COPY internal/secrets/firebase-credentials.json ./secrets/firebase-credentials.json

# Set location of SSL cert and key
ENV SSL_CERT_PATH ./certs/app.crt
ENV SSL_KEY_PATH ./certs/app.key
ENV GOOGLE_APPLICATION_CREDENTIALS ./secrets/firebase-credentials.json

# Expose port 10001 to the outside world
EXPOSE 10001

# Run the executable
CMD ["./main-alpine"]