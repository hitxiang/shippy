FROM golang:1.9.0

RUN mkdir /app
WORKDIR /app
#ADD consignment.json /app/consignment.json
#ADD consignment-cli /app/consignment-cli

ENTRYPOINT ["./consignment-cli", "consignment.json"]
CMD ["./consignment-cli"]
