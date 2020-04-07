FROM golang:latest

# Move to working directory
WORKDIR /go/src/app

# Copy and download dependency using go mod
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the code into the container
COPY . .

# Build the application
RUN go build -o build github.com/wayming/superdata/command/dbloader

# Move to /dist directory as the place for resulting binary folder
WORKDIR /dist

# Copy binary from build to main folder
RUN cp /go/src/app/build/dbloader .

# Copy datafiles to main folder
RUN cp -r /go/src/app/datafiles .
