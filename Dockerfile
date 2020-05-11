# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/oanaOM/weatherapp/src

# Build the weatherapp command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN go install github.com/oanaOM/weatherapp/src

# Run the weatherapp command by default when the container starts.
ENTRYPOINT /go/bin/weatherapp

# Document that the service listens on port 8080.
EXPOSE 8080

