FROM golang:1.16-alpine

WORKDIR /home/sprout/Documents/Go/mnc_test

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /docker-gs-ping

EXPOSE 8080

CMD ["/docker-gs-ping"]