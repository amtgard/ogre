# build stage
FROM golang:1.11 AS build-env
WORKDIR /go/src/app
COPY . .
RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/ogre

# final stage
FROM scratch
COPY --from=build-env /go/bin/ogre /go/bin/ogre
EXPOSE 3736
CMD ["/go/bin/ogre"]