FROM golang:1.20 AS build-stage

WORKDIR /build

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go install github.com/swaggo/swag/cmd/swag@latest \
        && swag init -g main.go

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /heimdall

FROM gcr.io/distroless/static-debian12:nonroot AS run-stage

WORKDIR /run

COPY --from=build-stage /heimdall .

EXPOSE 8080

ENTRYPOINT ["/run/heimdall"]
