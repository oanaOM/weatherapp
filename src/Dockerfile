# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

# Copy the local package files to the container's workspace.
COPY . /go/src/github.com/oanaOM/weatherapp/src
COPY templates/index.html /go/src/github.com/oanaOM/weatherapp/src/templates

# Build the weatherapp command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN go install github.com/oanaOM/weatherapp/src

# Run the weatherapp command by default when the container starts.
ENTRYPOINT /go/bin/src

# Document that the service listens on port 8080.
EXPOSE 8080

