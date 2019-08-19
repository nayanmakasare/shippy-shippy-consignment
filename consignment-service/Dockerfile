#FROM golang:alpine as builder
#
#RUN apk --no-cache add git
#
#WORKDIR /app/shippy-service-consignment
#
#COPY . .
#
#
#RUN go get -u "github.com/micro/go-micro"
#RUN go get -u "github.com/nayanmakasare/shippy-shippy-consignment/consignment-service/proto/consignment"
#RUN go get -u "go.mongodb.org/mongo-driver/mongo"
#RUN go get -u "go.mongodb.org/mongo-driver/mongo/options"
#
#RUN go mod download
#
## building binary from go code
#RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o shippy-service-consignment
#
#FROM alpine:latest
#
#RUN apk --no-cache add ca-certificates
#
#RUN mkdir -p /app
#WORKDIR /app
#
#COPY --from=builder /app/shippy-service-consignment/shippy-service-consignment .
#
#CMD ["./shippy-service-consignment"]


FROM debian:latest

RUN mkdir -p /app
WORKDIR /app

COPY consignment-service .

CMD ["./consignment-service"]