FROM debian:latest
RUN mkdir -p /app
WORKDIR /app

ADD consignment-cli /app/consignment-cli
ADD consignment.json /app/consignment.json

CMD ["./consignment-cli"]