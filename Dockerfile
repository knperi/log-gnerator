FROM icr.io/codeengine/golang:alpine
RUN apk -U upgrade
COPY helloworld.go /
RUN go build -ldflags '-s -w -extldflags "-static"' -o /tinyapp /icl-tinyapp.go

# Copy the exe into a smaller base image
FROM icr.io/codeengine/alpine
RUN apk -U upgrade
COPY --from=0 /icl-tinyapp /icl-tinyapp
CMD /icl-tinyapp
