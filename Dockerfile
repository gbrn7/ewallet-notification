FROM golang:1.23

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o ewallet-notification

EXPOSE 7003
EXPOSE 8083

CMD [ "./ewallet-notification" ]
