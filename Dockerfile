# Start from the latest golang base image
FROM golang:latest as builder

# Add Maintainer Info
LABEL maintainer="Gehad Hamdy"

# prepare the workspace
WORKDIR /go/src/jumia
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

######## Start a new stage from scratch #######
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /go/src/jumia .

ADD migrate /usr/local/bin/migrate
ADD scripts/up_migrate.sh /usr/local/bin/up_migrate.sh
ADD .docker/entrypoint.sh /usr/local/bin/

# Expose port 3050 to the outside world
EXPOSE 3400

# Command to run the executable
ENTRYPOINT ["/usr/local/bin/entrypoint.sh"]
