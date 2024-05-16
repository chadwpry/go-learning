FROM golang:1.22-alpine3.19 AS build

WORKDIR /app

COPY web-service-gin ./web-service-gin
COPY rest ./rest
COPY go.work .
COPY go.work.sum .

RUN go mod download
RUN go build -o web-service-gin/web-service-gin ./web-service-gin
RUN chmod +x web-service-gin/web-service-gin

EXPOSE 8080

CMD ["/app/web-service-gin/web-service-gin"]
