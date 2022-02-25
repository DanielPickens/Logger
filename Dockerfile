FROM golang:1.17

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/github.com/DanielPickens/Logger

# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY . .

# Download all the dependencies
RUN go mod tidy 
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

# This container exposes port 8080 to the outside world
EXPOSE 8080

# Run the executable
CMD ["logger"]
