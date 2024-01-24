From golang:1.21-alpine as builder
WORKDIR /app

COPY . .
RUN go mod download
# Build the Go app
RUN go build -o main .
# Use Ubuntu as the base image for the final image
FROM ubuntu:20.04

# Set the working directory to /app
WORKDIR /app

# Copy only the built executable from the builder image
COPY --from=builder /app/main .

EXPOSE 8080

CMD [ "./main" ]