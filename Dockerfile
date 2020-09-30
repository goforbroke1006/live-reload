FROM golang:1.14 as build-env

WORKDIR /code/

COPY go.mod .
COPY go.sum .

#RUN GOPROXY=direct go mod download
RUN go mod download

COPY ./ ./
RUN ls -lah

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o live-reload ./

FROM debian:stretch

COPY --from=build-env /code/live-reload /usr/local/bin/live-reload

WORKDIR /app/

ENTRYPOINT [ "live-reload" ]
