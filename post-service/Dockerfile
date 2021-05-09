FROM golang:1.15.12-alpine3.13

# Set the Current Working Directory inside the container
WORKDIR /app

RUN apk add git

# Download all the dependencies
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY . .

## Our project will now successfully build with the necessary go libraries included.
RUN go build -o main .

# This container exposes port 3000 to the outside world
EXPOSE 3000

# Run
CMD ["/app/main"]