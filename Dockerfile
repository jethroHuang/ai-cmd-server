FROM golang:1.20

WORKDIR /app
COPY ./ /app

RUN go build -o server

EXPOSE 1323

ENTRYPOINT ["./server"]