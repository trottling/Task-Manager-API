FROM golang:1.24 AS build
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .


RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server ./cmd/app

FROM scratch
WORKDIR /
COPY --from=build /app/server /server

EXPOSE 8080
ENTRYPOINT ["/server"]
