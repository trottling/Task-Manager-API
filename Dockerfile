FROM golang:1.24 AS build
WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /server ./cmd

FROM scratch
WORKDIR /
COPY --from=build /server /server

EXPOSE 8080
ENTRYPOINT ["/server"]
