FROM golang:1.22.1-alpine AS builder

WORKDIR /usr/local/src

RUN apk --no-cache add bash git make gettext

#dependencies
COPY ["go.mod", "go.sum", "./"] 
RUN go mod download

#build
COPY . ./
RUN go build -o ./bin/app cmd/adamenblog-api/main.go

FROM alpine AS runner

COPY --from=builder /usr/local/src/bin/app /
COPY ./config/local.yaml /config/local.yaml
COPY .env .env

EXPOSE 8080

CMD ["/app"]