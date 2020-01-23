# Alpine base image
FROM alpine

# Set the Current Working Directory inside the container
WORKDIR /go/src/github.com/RoomFlow/backend/services/search

# Copy the source from service directory to the Working Directory inside the container
COPY services/search/main-alpine ./main-alpine

# Copy service certs to container
COPY internal/certs/search ./certs

# Copy firebase credentials to container
COPY internal/secrets/firebase-credentials.json ./secrets/firebase-credentials.json

# Set location of SSL cert and key
ENV SSL_CERT_PATH ./certs/app.crt
ENV SSL_KEY_PATH ./certs/app.key
ENV GOOGLE_APPLICATION_CREDENTIALS ./secrets/firebase-credentials.json

# Expose port to the outside world
EXPOSE 10001

# Run the executable
CMD ["./main-alpine"]
