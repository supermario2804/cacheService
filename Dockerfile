FROM golang:latest AS builder

RUN mkdir /go/src/app
RUN go get -u github.com/golang/dep/cmd/dep

WORKDIR /go/src/cacheDataService
COPY . ./
COPY Gopkg.toml Gopkg.lock ./
RUN dep ensure --vendor-only
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o ./dist/cacheDataService

FROM alpine
COPY --from=builder /go/src/cacheDataService/dist/cacheDataService ./app
EXPOSE 8080 8090 5672 64919
ENTRYPOINT ["./app"]