# Use a builder image to compile the go lang code
FROM icr.io/codeengine/golang:alpine AS builder
COPY tinyapp.go /
RUN go build -o /tinyapp /tinyapp.go

# Copy the exe into a smaller base image for runtime
FROM icr.io/codeengine/alpine
COPY --from=builder /tinyapp /tinyapp
CMD /tinyapp
