# Telling to use Docker's golang ready image
FROM golang
# Name and Email of the author 
MAINTAINER OM 
# Create app folder 
RUN mkdir /app
# Copy our file in the host contianer to our contianer
ADD . /app
# Set /app to the go folder as workdir
WORKDIR /app
# Generate binary file from our /app
RUN go build
# Expose the port 8080
EXPOSE 8080
# Run the app binarry file 
CMD ["./app"]