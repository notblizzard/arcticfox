FROM golang:1.14.4-buster

WORKDIR /usr/src/app

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

RUN go build

CMD "./arcticfox"