FROM golang:1.15.0-alpine as builder

# Add Maintainer Info
LABEL maintainer="Rajeev Singh <rajeevhub@gmail.com>"

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

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
EXPOSE 3000

# Command to run the executable
CMD ["./main"]



#
#RUN go run -o main .
#
#CMD ["/app/main"]

#
#RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
#
#FROM alpine:latest
#
#RUN apk --no-cache add ca-certificates
#RUN apk add --no-cache git make musl-dev go
#COPY --from=builder /app/main .
#
#ENV GOROOT /user/lib/go
#ENV GOPATH /go
#ENV PATH /go/bin:$PATH
#
#
#RUN mkdir -p ${GOPATH}/src ${GoPATH}/bin