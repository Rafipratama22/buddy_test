FROM golang:1.16-alpine

WORKDIR /home/sprout/Documents/Go/mnc_test

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ./ ./

RUN go build -o /docker/mnc_test

EXPOSE 8080

CMD ["/docker/mnc_test"]