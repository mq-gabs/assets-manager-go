FROM golang:1.21-alpine

WORKDIR /usr/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN go build -o app

CMD [ "./app" ]