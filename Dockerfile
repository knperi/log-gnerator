FROM icr.io/codeengine/golang:alpine
RUN apk -U upgrade
COPY icl-demo-app.go /
RUN go build -ldflags '-s -w -extldflags "-static"' -o /tinyapp /icl-demo-app.go

# Copy the exe into a smaller base image
FROM icr.io/codeengine/alpine
RUN apk -U upgrade
COPY --from=0 /icl-demo-app /icl-demo-app
CMD /icl-demo-app
