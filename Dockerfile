# syntax=docker/dockerfile:1

FROM golang:1.19.2-bullseye


WORKDIR /app

COPY go.* ./ 

RUN go mod download

COPY *.go ./


RUN go build -o .

RUN ls

EXPOSE 8080

RUN chmod +x ./music-memory

CMD ["./music-memory"]
