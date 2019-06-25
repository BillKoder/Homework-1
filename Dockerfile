# Select image to use as a "builder" image and give it a name
FROM golang:latest as GOLANG-BUILDER

# Copy or source code into the builder image
COPY main.go /src/

WORKDIR /src

RUN go get "github.com/gorilla/mux"

# Compile the code
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/restfulAPI ./main.go

# Select a base image for out final image
FROM debian

# Copy the compiled file from the base image to the final image
COPY --from=GOLANG-BUILDER /app/restfulAPI /app/restfulAPI

# Set the default command to be executed when container is started from this image
CMD [ "/app/restfulAPI" ]