FROM golang:1.20

WORKDIR /usr/src/app

COPY . .

RUN go build -o bin/ChessNet

CMD ["./bin/ChessNet"]