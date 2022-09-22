FROM golang:1.19.1-bullseye

WORKDIR /grpc_server

COPY . .

RUN go build -o /server

EXPOSE 9000

RUN chmod +x /server

CMD [ "/server" ]