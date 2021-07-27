FROM golang:1.16

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/github.com/ernesgonzalez33/httpsserver

# Copy everything from the current directory to the PWD(Present Working Directory) inside the container
COPY . .

# Download all the dependencies
#RUN go get -d -v ./...

# This container exposes ports 8000 and 9000 to the outside world
EXPOSE 8000
EXPOSE 9000

# Run the executable
CMD ["go", "run", "server.go"]