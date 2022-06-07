# build stage
FROM golang:1.18 as build

ENV CGO_ENABLED 0
ENV GO111MODULE on

WORKDIR /go/src/go-grpc-microservice-template
COPY . .

RUN go get -v
RUN go vet -v
RUN go install -v

# run stage
FROM busybox as run

COPY --from=build /go/bin/go-grpc-microservice-template /go-grpc-microservice-template

EXPOSE 15200
EXPOSE 8080

CMD ["/go-grpc-microservice-template"]
