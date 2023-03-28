FROM golang:1.19-alpine AS build

WORKDIR /app

COPY . .

RUN go mod init todo/server

RUN go get .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /webserver


FROM gcr.io/distroless/static-debian11

ENV GIN_MODE=release
ENV PORT=8080

WORKDIR /

COPY --from=build /webserver /webserver

EXPOSE $PORT

USER nonroot:nonroot

CMD [ "/webserver" ]
