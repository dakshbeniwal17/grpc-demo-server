FROM golang:1.19.1-bullseye

WORKDIR /grpc_server

COPY . .
ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.9.0/wait /wait

RUN go build -o /server

EXPOSE 9000

RUN chmod +x /server
RUN chmod +x /wait
# RUN chmod +x ./wait-for.sh

CMD /wait && /server