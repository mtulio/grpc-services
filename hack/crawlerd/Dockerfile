# Start from golang v1.11 base image
FROM golang:1.11 as builder

# Add Maintainer Info
LABEL maintainer="MarcoB <root @ mtulio dot eng dot br>"

WORKDIR /go/src/github.com/mtulio/grpc-services

ENV GO111MODULE=on

# Copy everything from the current directory to the PWD(Present Working Directory) inside the container
COPY . .

# Download all the dependencies
# https://stackoverflow.com/questions/28031603/what-do-three-dots-mean-in-go-command-line-invocations
# RUN go get -d -v ./...

# Build the Go app
#RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/crawlerd ./cmd/crawlerd/*.go
#RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo ./cmd/crawlerd/*.go


######## Start a new stage from scratch #######
FROM alpine:latest  

RUN apk --no-cache add ca-certificates

WORKDIR /grpc-services/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /go/bin/crawlerd .

EXPOSE 8080

CMD ["/grpc-services/crawlerd"]