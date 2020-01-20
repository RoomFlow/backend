# Start from golang:1.13-alpine base image
FROM golang:1.13-alpine

# The latest alpine images don't have some tools like (`git` and `bash`).
# Adding git, bash and openssh to the image
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

# Set the Current Working Directory inside the container
WORKDIR /go/src/github.com/RoomFlow/backend/usermanagement/search

# Copy the source from the current directory to the Working Directory inside the container
COPY services/usermanagement .

# Copy usermanagement certs to container
COPY internal/certs/usermanagement ./certs

# Set location of SSL cert and key
ENV SSL_CERT_PATH ./certs/app.crt
ENV SSL_KEY_PATH ./certs/app.key

# Copy common files to container
COPY pkg /go/src/github.com/RoomFlow/backend/pkg

# Expose port 10002 to the outside world
EXPOSE 10002

# Run the executable
CMD ["./main-alpine"]