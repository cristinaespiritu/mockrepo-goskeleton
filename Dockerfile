# Copyright 2020 Dell Inc, or its subsidiaries.
#
# SPDX-License-Identifier: Apache-2.0

# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from the latest golang base image
FROM golang:latest as builder

# Labels
LABEL Author="Dell Technologies"
LABEL Version="0.0.1"

# Variables
ENV PORT 8080
ENV WDIR /go/src/app

# Create app directory and set the Current Working Directory inside the container
RUN mkdir -p ${WDIR}
WORKDIR ${WDIR}

# Bundle app source
COPY . ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod tidy

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .


######## Start a new stage from scratch #######
FROM alpine:latest  

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .

# Expose port 8080 to the outside world
# Port for Web
EXPOSE ${PORT}

# Healthcheck
HEALTHCHECK --interval=5m --timeout=3s \
  CMD curl -f http://localhost:${PORT} || exit 1

# Command to run the executable
CMD ["./main"] 