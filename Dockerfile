FROM golang:1.17-alpine AS build

RUN apk add --no-cache git

WORKDIR /tmp/go-api

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o ./out/go-api .

FROM alpine:3.9

RUN apk add ca-certificates

COPY --from=build /tmp/go-api/out/go-api /app/go-api

EXPOSE 8080

CMD ["/app/go-api"]