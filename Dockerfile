FROM golang:1.22-alpine AS build

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /webserver cmd/api/main.go

FROM gcr.io/distroless/static-debian11

ENV GIN_MODE=release
ENV PORT=8080

WORKDIR /

COPY --from=build /webserver /webserver

EXPOSE $PORT

USER nonroot:nonroot

CMD [ "/webserver" ]
